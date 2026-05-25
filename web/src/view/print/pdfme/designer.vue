<template>
  <div class="page-container">
    <header class="page-header">
      <div class="left-action-area">
        <el-button type="primary" @click="toggleView">
          {{ currentView === 'designer' ? '切换至 JSON 模式' : '切换至 画布模式' }}
        </el-button>
      </div>

      <div class="right-action-area">
        <div v-show="currentView === 'designer'" class="button-group">
          <el-button type="primary" size="small" plain @click="handleSave">完成</el-button>
          <el-button type="warning" size="small" plain @click="handleClear">清空</el-button>
          <el-button type="info" size="small" plain @click="handleCancel">取消</el-button>
        </div>
        <div v-show="currentView === 'json'" class="button-group">
          <el-button type="primary" size="small" plain @click="handleSave">完成</el-button>
          <el-button type="info" size="small" plain @click="handleCancel">取消</el-button>
        </div>
      </div>
    </header>

    <main class="page-body">
      <div class="content-box">
        <div v-show="currentView === 'designer'" ref="designerRootRef" class="designer-canvas-container"></div>

        <div v-show="currentView === 'json'" class="json-editor-wrapper">
          <div class="editor-tip">💡 提示：您可以直接在下方修改 JSON 字符串，切回画布模式后将自动应用更改。</div>
          <el-input v-model="templateJsonString" type="textarea" autosize placeholder="正在实时同步画布 JSON 数据..."
            class="json-textarea" />
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, shallowRef, onMounted, onUnmounted, toRaw } from 'vue'
import { Designer } from '@pdfme/ui'
import { Template } from '@pdfme/common'
import { text, image, table, barcodes, line, rectangle, ellipse, dateTime, select, radioGroup, checkbox } from '@pdfme/schemas'
import { ElMessage, ElMessageBox } from 'element-plus'

// --- 1. 常量与插件配置 ---
const PDFME_PLUGINS = {
  text,
  image,
  table,
  qrcode: barcodes.qrcode,
  line,
  rectangle,
  ellipse,
  dateTime,
  select,
  radioGroup,
  checkbox
};

// --- 2. 响应式状态管理 ---
const currentView = ref<'designer' | 'json'>('designer') // 替代原 activeDesign，更具可扩展性
const designerRootRef = ref<HTMLDivElement | null>(null)
const designerInstance = shallowRef<Designer | null>(null) // 第三方复杂实例使用 shallowRef 防止 Vue 深度监听导致克隆错误

const templateJsonString = ref('')

// 内部核心维护的 Template 对象（单一事实源）
const templateData = ref<Template>({
  basePdf: { width: 210, height: 297, padding: [10, 10, 10, 10] },
  schemas: [
    [
      { name: 'title', type: 'text', position: { x: 10, y: 10 }, width: 100, height: 20 },
      { name: 'subtitle', type: 'text', position: { x: 10, y: 35 }, width: 100, height: 15 }
    ]
  ]
})

// --- 3. 核心数据同步逻辑 ---

/**
 * 将当前画布的设计数据同步至 JSON 文本变量中
 */
const syncDesignerToText = () => {
  if (!designerInstance.value) return
  const latestTemplate = designerInstance.value.getTemplate()
  templateData.value = latestTemplate
  templateJsonString.value = JSON.stringify(latestTemplate, null, 2)
}

/**
 * 将文本框中的 JSON 数据解析并同步至设计器画布中
 */
const syncTextToDesigner = (): boolean => {
  try {
    const parsedTemplate = JSON.parse(templateJsonString.value) as Template

    // 校验 PDFme 格式合法性基准
    if (!parsedTemplate.basePdf || !parsedTemplate.schemas) {
      throw new Error("JSON 格式不符合 pdfme 模板规范，必须包含 basePdf 和 schemas 属性。")
    }

    templateData.value = parsedTemplate

    if (designerInstance.value) {
      // 深度断开响应式引用（清洗数据），防止响应式代理污染 pdfme 内部
      const rawTemplate = JSON.parse(JSON.stringify(toRaw(templateData.value)))
      designerInstance.value.updateTemplate(rawTemplate)
    }
    return true
  } catch (error: any) {
    ElMessage.error(error.message || 'JSON 解析失败，请检查语法格式')
    return false
  }
}

/**
 * 视图模式切换开关
 */
const toggleView = () => {
  if (currentView.value === 'designer') {
    // 离开画布进入 JSON 视图：先拉取画布数据
    syncDesignerToText()
    currentView.value = 'json'
  } else {
    // 离开 JSON 进入画布视图：先尝试解析文本
    const isSuccess = syncTextToDesigner()
    if (isSuccess) {
      currentView.value = 'designer'
      ElMessage.success('数据同步成功，设计画布已重新加载！')
    }
  }
}

// --- 4. PDFme 实例生命周期控制 ---

/**
 * 初始化 PDFme 设计器
 */
const initDesigner = (initialTemplate: Template) => {
  if (!designerRootRef.value || designerInstance.value) return

  // 深度断开 Vue 响应式追踪，避免 pdfme 操作 DOM 触发 DataCloneError
  const pureTemplate = JSON.parse(JSON.stringify(toRaw(initialTemplate)))

  designerInstance.value = new Designer({
    domContainer: designerRootRef.value,
    template: pureTemplate,
    plugins: PDFME_PLUGINS,
  })
}

// --- 5. 页面顶部操作栏按钮业务回调 ---

const handleSave = () => {
  if (currentView.value === 'designer') {
    syncDesignerToText()
  } else {
    const isSuccess = syncTextToDesigner()
    if (!isSuccess) return
  }

  console.log('最终保存的模板数据对象: ', toRaw(templateData.value))
  ElMessage.success('模板数据保存成功')
}

const handleClear = () => {
  ElMessageBox.confirm('确定要清空当前画布上的所有控件吗？此操作不可撤销。', '提示', {
    confirmButtonText: '确定清空',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    if (designerInstance.value) {
      const emptyTemplate: Template = {
        basePdf: toRaw(templateData.value.basePdf),
        schemas: [[]]
      }
      designerInstance.value.updateTemplate(emptyTemplate)
      templateData.value = emptyTemplate
      templateJsonString.value = JSON.stringify(emptyTemplate, null, 2)
      ElMessage.success('画布已成功清空')
    }
  }).catch(() => { })
}

const handleCancel = () => {
  ElMessage.info('已取消操作')
}

// --- 6. 页面生命周期挂载 ---
onMounted(() => {
  // 页面加载完成时同步初始化字符串
  templateJsonString.value = JSON.stringify(templateData.value, null, 2)
  // 全局仅执行一次单例初始化
  initDesigner(templateData.value)
})

onUnmounted(() => {
  if (designerInstance.value) {
    designerInstance.value.destroy()
    designerInstance.value = null
  }
})
</script>

<style scoped lang="scss">
.page-container {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background-color: #ffffff;
  border-bottom: 1px solid var(--el-border-color-light);
}

.left-action-area {
  display: flex;
  gap: 12px;
}

.right-action-area {
  display: flex;
  flex-direction: row-reverse; // 使得先写的按钮位于最右侧

  .button-group {
    display: flex;
    flex-direction: row-reverse;
    gap: 12px;
  }
}

.page-body {
  flex: 1;
  padding: 20px;
  background-color: #f5f7fa;
  overflow: hidden;
}

.content-box {
  width: 100%;
  height: 100%;
  background: #ffffff;
  border-radius: 4px;
}

/* JSON 模式下的外部包裹区样式 */
.json-editor-wrapper {
  width: 100%;
  height: calc(100vh - 120px);
  display: flex;
  flex-direction: column;
  padding: 12px;
  box-sizing: border-box;

  .editor-tip {
    padding: 10px 15px;
    font-size: 13px;
    color: #e6a23c;
    background-color: #fdf6ec;
    border-radius: 4px;
    margin-bottom: 12px;
  }

  .json-textarea {
    flex: 1;
    font-family: 'Courier New', Courier, monospace;

    :deep(.el-textarea__inner) {
      height: 100% !important;
      resize: none;
      background-color: #fafafa;
      color: #333333;
    }
  }
}

/* PDFme 画布模式下的宿主节点样式 */
.designer-canvas-container {
  width: 100%;
  height: calc(100vh - 120px);
  background-color: #f5f7fa;

  // 深度选择器，强行让 pdfme 自有渲染层填满页面容器
  :deep(.pdfme-designer) {
    height: 100% !important;
    width: 100% !important;
  }
}
</style>