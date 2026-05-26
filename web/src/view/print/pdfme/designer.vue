<template>
  <div class="pdfme-rigid-sandbox">
    <div class="toolbar">
      <el-button type="primary" size="small" plain @click="handleAction('submit')">完成</el-button>
      <el-button type="warning" size="small" plain @click="handleClear">清空</el-button>
      <el-button type="info" size="small" plain @click="handleAction('cancel')">取消</el-button>
      <div class="editor-tip">💡 提示：修改左侧JSON后自动同步到右侧设计器</div>
    </div>

    <div class="main-content">
      <div class="left-panel" :style="{ width: leftWidth + '%' }">
        <div class="editor-container">
          <div ref="jsonEditorRef" class="monaco-editor"></div>
        </div>
      </div>

      <div class="resizer" :class="{ dragging: isDragging }" @mousedown="startDrag">
        <span class="resize-icon">⋮⋮</span>
      </div>

      <div class="right-panel" :style="{ width: (100 - leftWidth - 0.5) + '%' }">
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

const leftWidth = ref(40)
const isDragging = ref(false)
const designerRootRef = ref<HTMLDivElement | null>(null)
const designerInstance = shallowRef<Designer | null>(null)
const jsonEditorRef = ref<HTMLDivElement | null>(null)
const templateJsonString = ref('')
const localTemplateData = ref<Template>({ basePdf: { width: 210, height: 297, padding: [10, 10, 10, 10] }, schemas: [[]] })

let jsonEditor: monaco.editor.IStandaloneCodeEditor | null = null
let debounceTimer: number | null = null
// 用来阻止内部循环更新的标志位
let isInternalUpdating = false

// 【优化】右侧设计器变动 -> 同步到左侧文本框
const syncDesignerToText = () => {
  if (!designerInstance.value || isInternalUpdating) return
  isInternalUpdating = true
  try {
    const latestTemplate = designerInstance.value.getTemplate()
    localTemplateData.value = latestTemplate
    const jsonString = JSON.stringify(latestTemplate, null, 2)
    templateJsonString.value = jsonString
    if (jsonEditor) {
      const currentValue = jsonEditor.getValue()
      if (currentValue !== jsonString) {
        // preserveFocus 保证用户在操作右侧时，左侧安静更新而不抢焦点
        jsonEditor.setValue(jsonString)
      }
    }
  } catch (error) {
    console.error('Failed to sync designer to text', error)
  } finally {
    isInternalUpdating = false
  }
}

// 【优化】左侧文本框变动 -> 同步到右侧设计器
const syncTextToDesigner = (): boolean => {
  if (!jsonEditor || isInternalUpdating) return false
  isInternalUpdating = true
  try {
    const jsonString = jsonEditor.getValue()
    const parsedTemplate = JSON.parse(jsonString) as Template
    if (!parsedTemplate.basePdf || !parsedTemplate.schemas) {
      throw new Error("JSON 格式不合规范，必须包含 basePdf 和 schemas 属性。")
    }
    localTemplateData.value = parsedTemplate
    templateJsonString.value = jsonString

    if (designerInstance.value) {
      // 深度拷贝切断引用，防止 pdfme 篡改 Vue 响应式数据
      const rawData = JSON.parse(JSON.stringify(toRaw(localTemplateData.value)))
      designerInstance.value.updateTemplate(rawData)
    }
    return true
  } catch (error: any) {
    // 语法错误时仅在控制台提示或轻量提示，避免打字时高频弹窗打断思路
    console.warn('JSON Parsing...', error.message)
    return false
  } finally {
    isInternalUpdating = false
  }
}

const debouncedSyncTextToDesigner = () => {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = window.setTimeout(() => {
    syncTextToDesigner()
  }, 500) // 500ms 延迟合理，打字停顿时同步
}

const initDesigner = async (initialTemplate: Template) => {
  if (!designerRootRef.value || designerInstance.value) return

  const fontOptions: any = await autoLoadAllFonts()
  designerInstance.value = new Designer({
    domContainer: designerRootRef.value,
    template: JSON.parse(JSON.stringify(toRaw(initialTemplate))),
    plugins: PDFME_PLUGINS,
    options: {
      font: fontOptions
    }
  })

  designerInstance.value.onChangeTemplate(() => {
    syncDesignerToText()
  })
}

const initJsonEditor = () => {
  if (!jsonEditorRef.value || jsonEditor) return

  jsonEditor = monaco.editor.create(jsonEditorRef.value, {
    value: templateJsonString.value,
    language: 'json',
    theme: 'vs',
    automaticLayout: true, // 保持开启，作为防守基础
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

  jsonEditor.onDidChangeModelContent(() => {
    debouncedSyncTextToDesigner()
  })
}

const handleAction = (status: 'submit' | 'cancel') => {
  if (status === 'cancel') {
    emit('action-completed', { actionStatus: 'cancel', template: null })
    return
  }

  const jsonString = jsonEditor ? jsonEditor.getValue() : templateJsonString.value
  try {
    const finalData = JSON.parse(jsonString) as Template
    emit('update:modelValue', finalData)
    emit('action-completed', { actionStatus: 'submit', template: finalData })
  } catch {
    ElMessage.error('当前 JSON 格式有误，无法提交')
  }
}

const handleClear = () => {
  ElMessageBox.confirm('确定要清空当前画布上的所有控件吗？', '提示', {
    confirmButtonText: '确定清空', cancelButtonText: '取消', type: 'warning'
  }).then(() => {
    if (designerInstance.value) {
      const emptyTemplate: Template = { basePdf: toRaw(localTemplateData.value.basePdf), schemas: [[]] }

      isInternalUpdating = true // 锁定，防止 setValue 触发多余更新
      designerInstance.value.updateTemplate(emptyTemplate)
      localTemplateData.value = emptyTemplate
      const jsonString = JSON.stringify(emptyTemplate, null, 2)
      templateJsonString.value = jsonString
      if (jsonEditor) {
        jsonEditor.setValue(jsonString)
      }
      isInternalUpdating = false

      ElMessage.success('画布已成功清空')
    }
  }).catch(() => { })
}

const startDrag = (e: MouseEvent) => {
  isDragging.value = true
  document.addEventListener('mousemove', onDrag)
  document.addEventListener('mouseup', stopDrag)
  e.preventDefault()
}

const onDrag = (e: MouseEvent) => {
  if (!isDragging.value) return

  const container = document.querySelector('.main-content') as HTMLElement
  if (!container) return

  const rect = container.getBoundingClientRect()
  const x = e.clientX - rect.left
  const percentage = (x / rect.width) * 100

  leftWidth.value = Math.max(20, Math.min(70, percentage))

  // 【优化】拖拽时实时让 Monaco 重新计算宽高，防止出现白边和闪烁
  if (jsonEditor) {
    jsonEditor.layout()
  }
}

const stopDrag = () => {
  isDragging.value = false
  document.removeEventListener('mousemove', onDrag)
  document.removeEventListener('mouseup', stopDrag)
  // 【优化】停止拖拽时再强制 layout 一次确保精确度
  if (jsonEditor) {
    jsonEditor.layout()
  }
}

onMounted(async () => {
  if (props.modelValue && props.modelValue.basePdf && props.modelValue.schemas) {
    localTemplateData.value = JSON.parse(JSON.stringify(toRaw(props.modelValue)))
  } else {
    localTemplateData.value = JSON.parse(JSON.stringify(DEFAULT_A4_TEMPLATE))
  }
  templateJsonString.value = JSON.stringify(localTemplateData.value, null, 2)

  nextTick(() => {
    initJsonEditor()
  })

  await initDesigner(localTemplateData.value)
})

onUnmounted(() => {
  // 【优化】严格按照先注销再清空的顺序销毁实例，防止内存泄漏
  if (designerInstance.value) {
    try {
      designerInstance.value.destroy()
    } catch (e) {
      console.error(e)
    }
    designerInstance.value = null
  }
  if (jsonEditor) {
    jsonEditor.dispose()
    jsonEditor = null
  }
  if (debounceTimer) {
    clearTimeout(debounceTimer)
  }
})
</script>

<style scoped lang="scss">
/* 保持你的优秀样式不变 */
.pdfme-rigid-sandbox {
  width: 100%;
  height: calc(100vh - 180px);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.toolbar {
  display: flex;
  gap: 12px;
  padding: 8px 0 0 0;
  align-items: center;

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
}

.left-panel {
  display: flex;
  flex-direction: column;
  overflow: hidden;
  height: 100%;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
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
  /* 稍微调宽2px增加鼠标易用性 */
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

.resizer:hover .resize-icon,
.resizer.dragging .resize-icon {
  color: #fff;
}

.right-panel {
  flex: 1;
  min-height: 0;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  overflow: hidden;
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