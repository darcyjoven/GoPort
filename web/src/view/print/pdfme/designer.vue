<template>
  <div class="pdfme-rigid-sandbox">
    <div class="toolbar">
      <div class="left-controls">
        <el-switch v-model="showJsonEditor" inactive-text="JSON 编辑器" @change="handleEditorVisibilityChange" />
      </div>

      <div class="right-actions">
        <el-button type="primary" size="small" plain @click="handleToolbarAction('submit')">完成</el-button>
        <el-button type="warning" size="small" plain @click="handleClearCanvas">清空</el-button>
        <el-button type="info" size="small" plain @click="handleToolbarAction('cancel')">取消</el-button>
      </div>
      <div class="editor-tip">
        💡 提示：{{ showJsonEditor ? '修改 JSON 后失焦同步到设计器，聚焦时从设计器拉取' : '当前已关闭 JSON 视图，专注于画布拖拽设计' }}
      </div>
    </div>

    <div class="main-content">
      <div v-if="showJsonEditor" class="left-panel" :style="{ width: panelSplitRatio + '%' }">
        <div class="editor-container">
          <div ref="jsonEditorRef" class="monaco-editor"></div>
        </div>
      </div>

      <div v-if="showJsonEditor" class="resizer" :class="{ dragging: isResizing }" @mousedown="handleResizeStart">
        <span class="resize-icon">⋮⋮</span>
      </div>

      <div class="right-panel" :style="{ width: showJsonEditor ? (100 - panelSplitRatio - 0.5) + '%' : '100%' }">
        <div ref="designerRootRef" class="pdfme-native-mount-point"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, shallowRef, onMounted, onUnmounted, toRaw, nextTick } from 'vue'
import { Designer } from '@pdfme/ui'
import { Template } from '@pdfme/common'
import {
  text, image, table, barcodes, line, rectangle, ellipse, dateTime, select, radioGroup,
  checkbox, multiVariableText, list, svg, signature, date, time
} from '@pdfme/schemas'
import { ElMessage, ElMessageBox } from 'element-plus'
import { autoLoadAllFonts } from '@/utils/fontLoader'
import * as monaco from 'monaco-editor'

// --- 类型声明与组件通信 ---
interface Props {
  modelValue?: Template | null
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: null
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: Template): void
  (e: 'action-completed', data: { actionStatus: 'submit' | 'cancel'; template: Template | null }): void
}>()

// --- PDFme 常量配置 ---
const PDFME_PLUGINS = {
  text, dateTime, date, time,
  multiVariableText, list,
  select, radioGroup, checkbox, line, rectangle, ellipse,
  image, table, svg, signature,
  qrcode: barcodes.qrcode, gs1datamatrix: barcodes.gs1datamatrix,
  ean13: barcodes.ean13, ean8: barcodes.ean8,
  code128: barcodes.code128, code39: barcodes.code39, itf14: barcodes.itf14,
  pdf317: barcodes.pdf417,
}

const DEFAULT_A4_TEMPLATE: Template = {
  basePdf: { width: 210, height: 297, padding: [10, 10, 10, 10] },
  schemas: [[
    { name: 'title', type: 'text', position: { x: 10, y: 10 }, width: 100, height: 20, value: '默认标题' }
  ]]
}

// --- 响应式核心状态 ---
const showJsonEditor = ref(false)   // 【核心控制】默认不显示 JSON 编辑器
const panelSplitRatio = ref(40)     // 左右面板黄金分割比 (百分比)
const isResizing = ref(false)       // 侧边栏拖拽状态

// DOM 节点引用
const jsonEditorRef = ref<HTMLDivElement | null>(null)
const designerRootRef = ref<HTMLDivElement | null>(null)

// 实例引用
const designerInstance = shallowRef<Designer | null>(null)
let jsonEditorInstance: monaco.editor.IStandaloneCodeEditor | null = null

// 核心同步与事务锁
let isSyncingLock = false
let resizeAnimationFrameId: number | null = null

// --- 核心业务逻辑：双向双关卡按需同步调度 ---

/**
 * 核心单向流：设计器 -> 编辑器 
 */
const syncDesignerToEditor = () => {
  // 安全阀门：如果编辑器根本没有显示，或者处于事务锁中，坚决不进行数据处理
  if (!showJsonEditor.value || !designerInstance.value || !jsonEditorInstance || isSyncingLock) return

  isSyncingLock = true
  try {
    const latestTemplate = designerInstance.value.getTemplate()
    const formattedJsonString = JSON.stringify(latestTemplate, null, 2)

    if (jsonEditorInstance.getValue() !== formattedJsonString) {
      jsonEditorInstance.setValue(formattedJsonString)
    }
  } catch (error) {
    console.error('[GVA-Pdfme] Failed to sync data from designer to editor:', error)
  } finally {
    isSyncingLock = false
  }
}

/**
 * 核心单向流：编辑器 -> 设计器
 * @returns {boolean} 同步成功返回 true，语法错误返回 false
 */
const syncEditorToDesigner = (): boolean => {
  // 安全阀门：如果编辑器没有显示，不具备同步源
  if (!showJsonEditor.value || !jsonEditorInstance || !designerInstance.value || isSyncingLock) return false

  isSyncingLock = true
  try {
    const editorContent = jsonEditorInstance.getValue()
    const parsedTemplate = JSON.parse(editorContent) as Template

    if (!parsedTemplate.basePdf || !parsedTemplate.schemas) {
      throw new Error("缺少规范定义的必需属性 (basePdf / schemas)")
    }

    const cleanTemplateData = JSON.parse(JSON.stringify(toRaw(parsedTemplate)))
    designerInstance.value.updateTemplate(cleanTemplateData)
    return true
  } catch (error: any) {
    console.warn('[GVA-Pdfme] JSON Validation Incomplete:', error.message)
    return false
  } finally {
    isSyncingLock = false
  }
}

// --- 组件初始化与销毁驱动 ---

/**
 * 初始化 Monaco JSON 编辑器
 */
const initMonacoEditor = (initialContent: string) => {
  if (!jsonEditorRef.value || jsonEditorInstance) return

  jsonEditorInstance = monaco.editor.create(jsonEditorRef.value, {
    value: initialContent,
    language: 'json',
    theme: 'vs',
    automaticLayout: true,
    minimap: { enabled: false },
    fontSize: 14,
    fontFamily: 'Consolas, "Courier New", monospace',
    tabSize: 2,
    insertSpaces: true,
    lineNumbers: 'on',
    scrollBeyondLastLine: false,
    folding: true,
    foldingHighlight: true,
    bracketPairColorization: { enabled: true },
    renderLineHighlight: 'line',
    padding: { top: 12, bottom: 12 }
  })

  // 严格按需触发：仅在失去焦点（Blur）时同步到设计器
  jsonEditorInstance.onDidBlurEditorText(() => {
    if (syncEditorToDesigner()) {
      ElMessage({ message: 'JSON 修改成功，已同步到设计器', type: 'success', duration: 1500 })
    } else {
      ElMessage({ message: 'JSON 格式不规范，未同步到设计器', type: 'warning', duration: 2000 })
    }
  })

  // 严格按需触发：仅在获取焦点（Focus）时，从画布拉取最新状态
  jsonEditorInstance.onDidFocusEditorText(() => {
    syncDesignerToEditor()
    ElMessage({ message: '已从设计器同步最新 JSON', type: 'info', duration: 1500 })
  })
}

/**
 * 销毁 Monaco 编辑器，释放内存空间并解除事件绑定
 */
const disposeMonacoEditor = () => {
  if (jsonEditorInstance) {
    jsonEditorInstance.dispose()
    jsonEditorInstance = null
  }
}

/**
 * 初始化 PDFme 画布设计器
 */
const initPdfmeDesigner = async (initialTemplate: Template) => {
  if (!designerRootRef.value || designerInstance.value) return

  const loadedFonts: any = await autoLoadAllFonts()
  designerInstance.value = new Designer({
    domContainer: designerRootRef.value,
    template: JSON.parse(JSON.stringify(toRaw(initialTemplate))),
    plugins: PDFME_PLUGINS,
    options: { font: loadedFonts }
  })
}

// --- 显隐切换交互行为 ---

/**
 * 响应 Switch 开关切换
 */
const handleEditorVisibilityChange = (visible: boolean | string | number) => {
  if (visible) {
    // 开启编辑器：1. 开启 DOM 挂载进程 2. 从当前设计器抓取最新版式代码 3. 初始化渲染 Monaco
    nextTick(() => {
      if (!designerInstance.value) return
      const currentDesignerTemplate = designerInstance.value.getTemplate()
      const jsonString = JSON.stringify(currentDesignerTemplate, null, 2)

      initMonacoEditor(jsonString)

      // 触发设计器的重绘布局，防止因为突然多出左侧分栏导致画布容器尺寸变形
      if (designerInstance.value) {
        designerInstance.value.onChangeTemplate(() => { }) // 刷新其内部监听
      }
    })
  } else {
    // 关闭编辑器：物理层面彻底干掉 Monaco 实例，断开全部 Blur/Focus 调度，完全停止数据同步解析逻辑
    disposeMonacoEditor()
  }
}

// --- 顶部工具栏行为响应 ---

/**
 * 完成提交与取消逻辑调度
 */
const handleToolbarAction = (actionStatus: 'submit' | 'cancel') => {
  if (actionStatus === 'cancel') {
    emit('action-completed', { actionStatus: 'cancel', template: null })
    return
  }

  // 提交时，如果编辑器处于打开状态且光标在内部，先做一次最后的冲刷同步
  if (showJsonEditor.value && jsonEditorInstance) {
    syncEditorToDesigner()
  }

  // 依据当前谁是主导者来提取最终的 Template 数据
  try {
    let finalTemplateData: Template
    if (showJsonEditor.value && jsonEditorInstance) {
      finalTemplateData = JSON.parse(jsonEditorInstance.getValue()) as Template
    } else if (designerInstance.value) {
      finalTemplateData = designerInstance.value.getTemplate()
    } else {
      finalTemplateData = props.modelValue || DEFAULT_A4_TEMPLATE
    }

    emit('update:modelValue', finalTemplateData)
    emit('action-completed', { actionStatus: 'submit', template: finalTemplateData })
  } catch {
    ElMessage.error('当前 JSON 格式有误，拦截提交请求')
  }
}

/**
 * 彻底清空当前画布上的图层控件结构
 */
const handleClearCanvas = () => {
  ElMessageBox.confirm('确定要清空当前画布上的所有控件吗？', '操作警告', {
    confirmButtonText: '确定清空',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    if (!designerInstance.value) return

    try {
      isSyncingLock = true

      const currentCanvasLayout = designerInstance.value.getTemplate()
      const emptiedTemplate: Template = {
        basePdf: JSON.parse(JSON.stringify(toRaw(currentCanvasLayout.basePdf))),
        schemas: [[]]
      }

      designerInstance.value.updateTemplate(emptiedTemplate)

      // 如果编辑器正开着，也同步清空
      if (showJsonEditor.value && jsonEditorInstance) {
        jsonEditorInstance.setValue(JSON.stringify(emptiedTemplate, null, 2))
      }

      ElMessage.success('画布已成功清空')
    } catch (err) {
      console.error(err)
    } finally {
      isSyncingLock = false
    }
  }).catch(() => { })
}

// --- 侧边栏拖拽缩放 ---

const handleResizeStart = (e: MouseEvent) => {
  isResizing.value = true
  document.addEventListener('mousemove', handleResizing)
  document.addEventListener('mouseup', handleResizeStop)
  e.preventDefault()
}

const handleResizing = (e: MouseEvent) => {
  if (!isResizing.value || !showJsonEditor.value) return

  if (resizeAnimationFrameId) cancelAnimationFrame(resizeAnimationFrameId)

  resizeAnimationFrameId = requestAnimationFrame(() => {
    const mainLayoutContainer = document.querySelector('.main-content') as HTMLElement
    if (!mainLayoutContainer) return

    const bounds = mainLayoutContainer.getBoundingClientRect()
    const relativeX = e.clientX - bounds.left
    const computedPercentage = (relativeX / bounds.width) * 100

    panelSplitRatio.value = Math.max(20, Math.min(70, computedPercentage))

    if (jsonEditorInstance) {
      jsonEditorInstance.layout()
    }
  })
}

const handleResizeStop = () => {
  isResizing.value = false
  document.removeEventListener('mousemove', handleResizing)
  document.removeEventListener('mouseup', handleResizeStop)

  if (resizeAnimationFrameId) {
    cancelAnimationFrame(resizeAnimationFrameId)
    resizeAnimationFrameId = null
  }

  if (jsonEditorInstance) {
    jsonEditorInstance.layout()
  }
}

// --- 生命钩子挂载与终结销毁 ---

onMounted(async () => {
  let initialTemplateSeed: Template
  if (props.modelValue && props.modelValue.basePdf && props.modelValue.schemas) {
    initialTemplateSeed = JSON.parse(JSON.stringify(toRaw(props.modelValue)))
  } else {
    initialTemplateSeed = JSON.parse(JSON.stringify(DEFAULT_A4_TEMPLATE))
  }

  // 1. 初始化右侧全屏设计器
  await initPdfmeDesigner(initialTemplateSeed)

  // 2. 默认 showJsonEditor 为 false，故不在此初始化 Monaco
})

onUnmounted(() => {
  if (designerInstance.value) {
    try {
      designerInstance.value.destroy()
    } catch (e) {
      console.error('[GVA-Pdfme] Error while destroying designer:', e)
    }
    designerInstance.value = null
  }

  disposeMonacoEditor()

  if (resizeAnimationFrameId) {
    cancelAnimationFrame(resizeAnimationFrameId)
  }
})
</script>

<style scoped lang="scss">
.pdfme-rigid-sandbox {
  width: 100%;
  height: calc(100vh - 180px);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.toolbar {
  display: flex;
  padding: 8px 0 0 0;
  align-items: center;
  width: 100%;

  .left-controls {
    display: flex;
    align-items: center;
    background: #f5f7fa;
    padding: 4px 16px;
    border-radius: 20px;
    border: 1px solid #e4e7ed;
  }

  .right-actions {
    display: flex;
    gap: 12px;
    margin-left: 24px;
  }

  .editor-tip {
    margin-left: auto;
    padding: 6px 12px;
    font-size: 12px;
    color: #e6a23c;
    background-color: #fdf6ec;
    border-radius: 4px;
  }
}

.main-content {
  flex: 1;
  display: flex;
  gap: 0;
  min-height: 0;
  width: 100%;
}

.left-panel {
  display: flex;
  flex-direction: column;
  overflow: hidden;
  height: 100%;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  transition: width 0.1s ease; // 平滑伸缩
}

.editor-container {
  flex: 1;
  overflow: hidden;
  min-height: 0;
}

.monaco-editor {
  width: 100%;
  height: 100%;
}

.resizer {
  width: 6px;
  cursor: col-resize;
  background-color: #dcdfe6;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
  flex-shrink: 0;
  z-index: 10;
}

.resizer:hover {
  background-color: #409eff;
}

.resizer.dragging {
  background-color: #409eff;
  cursor: col-resize;
}

.resize-icon {
  font-size: 12px;
  color: #999;
  user-select: none;
}

.right-panel {
  min-height: 0;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  overflow: hidden;
  transition: width 0.1s ease;
}

.pdfme-native-mount-point {
  width: 100% !important;
  height: 100% !important;
  max-width: 100% !important;
  max-height: 100% !important;
  overflow: hidden !important;

  :deep(.pdfme-designer) {
    width: 100% !important;
    height: 100% !important;
    max-width: 100% !important;
    max-height: 100% !important;
  }
}
</style>