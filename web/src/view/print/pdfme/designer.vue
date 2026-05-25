<template>
  <div class="pdfme-rigid-sandbox">

    <div class="editor-control-panel">
      <div class="panel-left">
        <el-button type="primary" @click="toggleView">
          {{ currentView === 'designer' ? '切换至 JSON 模式' : '切换至 画布模式' }}
        </el-button>

        <el-divider direction="vertical" />

        <template v-if="currentView === 'designer'">
          <el-button type="primary" plain @click="handleAction('submit')">完成</el-button>
          <el-button type="warning" plain @click="handleClear">清空</el-button>
          <el-button type="info" plain @click="handleAction('cancel')">取消</el-button>
        </template>

        <template v-else-if="currentView === 'json'">
          <el-button type="primary" plain @click="handleAction('submit')">完成</el-button>
          <el-button type="info" plain @click="handleAction('cancel')">取消</el-button>
        </template>
      </div>
    </div>

    <div class="editor-viewport-body">

      <div v-show="currentView === 'designer'" class="render-view-isolation">
        <div ref="designerRootRef" class="pdfme-native-mount-point"></div>
      </div>

      <div v-if="currentView === 'json'" class="render-view-isolation json-mode-layout">
        <div class="editor-tip">💡 提示：您可以直接在下方修改 JSON 字符串，切回画布模式后将自动应用更改。</div>
        <div class="textarea-container">
          <el-input v-model="templateJsonString" type="textarea" placeholder="正在实时同步画布 JSON 数据..."
            class="json-pure-textarea" />
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, shallowRef, onMounted, onUnmounted, toRaw } from 'vue'
import { Designer } from '@pdfme/ui'
import { Template } from '@pdfme/common'
import {
  text, image, table, barcodes, line, rectangle, ellipse, dateTime, select, radioGroup,
  checkbox, multiVariableText, list, svg, signature,date,time
} from '@pdfme/schemas'
import { ElMessage, ElMessageBox } from 'element-plus'
import { autoLoadAllFonts } from '@/utils/fontLoader'

// --- 1. 组件 Props 定义 ---
interface Props {
  modelValue?: Template | null // 外部 v-model 绑定的数据
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: null
})

// --- 2. 组件 Emits 定义 ---
// 通过 actionStatus 完美区分“完成(submit)”与“取消(cancel)”
const emit = defineEmits<{
  (e: 'update:modelValue', value: Template): void
  (e: 'action-completed', data: { actionStatus: 'submit' | 'cancel'; template: Template | null }): void
}>()

// --- 3. 基础常量与状态管理 ---
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

const currentView = ref<'designer' | 'json'>('designer')
const designerRootRef = ref<HTMLDivElement | null>(null)
const designerInstance = shallowRef<Designer | null>(null)
const templateJsonString = ref('')
const localTemplateData = ref<Template>({ basePdf: { width: 210, height: 297, padding: [10, 10, 10, 10] }, schemas: [[]] })

// --- 4. 数据核心同步链 ---
const syncDesignerToText = () => {
  if (!designerInstance.value) return
  const latestTemplate = designerInstance.value.getTemplate()
  localTemplateData.value = latestTemplate
  templateJsonString.value = JSON.stringify(latestTemplate, null, 2)
}

const syncTextToDesigner = (): boolean => {
  try {
    const parsedTemplate = JSON.parse(templateJsonString.value) as Template
    if (!parsedTemplate.basePdf || !parsedTemplate.schemas) {
      throw new Error("JSON 格式不合规范，必须包含 basePdf 和 schemas 属性。")
    }
    localTemplateData.value = parsedTemplate
    if (designerInstance.value) {
      designerInstance.value.updateTemplate(JSON.parse(JSON.stringify(toRaw(localTemplateData.value))))
    }
    return true
  } catch (error: any) {
    ElMessage.error(error.message || 'JSON 解析失败，请检查语法格式')
    return false
  }
}

const toggleView = () => {
  if (currentView.value === 'designer') {
    syncDesignerToText()
    currentView.value = 'json'
  } else {
    if (syncTextToDesigner()) {
      currentView.value = 'designer'
      ElMessage.success('数据同步成功，设计画布已重新加载！')
    }
  }
}

// --- 5. 初始化与销毁 ---
const initDesigner = async (initialTemplate: Template) => {
  if (!designerRootRef.value || designerInstance.value) return

  const fontOptions :any = await autoLoadAllFonts()
  console.log(fontOptions)
  designerInstance.value = new Designer({
    domContainer: designerRootRef.value,
    template: JSON.parse(JSON.stringify(toRaw(initialTemplate))),
    plugins: PDFME_PLUGINS,
    options:{
      font: fontOptions
    }
  })
}

// --- 6. 完成 / 取消 统一控制器 ---
const handleAction = (status: 'submit' | 'cancel') => {
  if (status === 'cancel') {
    emit('action-completed', { actionStatus: 'cancel', template: null })
    return
  }

  if (currentView.value === 'designer') {
    syncDesignerToText()
  } else {
    if (!syncTextToDesigner()) return
  }

  const finalData = JSON.parse(JSON.stringify(toRaw(localTemplateData.value)))
  emit('update:modelValue', finalData)
  emit('action-completed', { actionStatus: 'submit', template: finalData })
}

const handleClear = () => {
  ElMessageBox.confirm('确定要清空当前画布上的所有控件吗？', '提示', {
    confirmButtonText: '确定清空', cancelButtonText: '取消', type: 'warning'
  }).then(() => {
    if (designerInstance.value) {
      const emptyTemplate: Template = { basePdf: toRaw(localTemplateData.value.basePdf), schemas: [[]] }
      designerInstance.value.updateTemplate(emptyTemplate)
      localTemplateData.value = emptyTemplate
      templateJsonString.value = JSON.stringify(emptyTemplate, null, 2)
      ElMessage.success('画布已成功清空')
    }
  }).catch(() => { })
}

onMounted(async () => {  

  if (props.modelValue && props.modelValue.basePdf && props.modelValue.schemas) {
    localTemplateData.value = JSON.parse(JSON.stringify(toRaw(props.modelValue)))
  } else {
    localTemplateData.value = JSON.parse(JSON.stringify(DEFAULT_A4_TEMPLATE))
  }
  templateJsonString.value = JSON.stringify(localTemplateData.value, null, 2)
  initDesigner(localTemplateData.value)
})

onUnmounted(() => {
  if (designerInstance.value) {
    designerInstance.value.destroy()
    designerInstance.value = null
  }
})
</script>

<style scoped lang="scss">
/* 【降维打击重构】：利用 Grid 彻底替代 Flex 和 Absolute
  确保无论外部父级是什么妖魔鬼怪环境，本组件在物理上严格限制在可用视口内。
*/
.pdfme-rigid-sandbox {
  display: grid;
  grid-template-rows: 56px 1fr;
  /* 按钮固定 56px，下方占满剩余空间 */
  width: 100% !important;
  max-width: 100% !important;
  height: 100% !important;
  max-height: 100% !important;
  overflow: hidden !important;
  box-sizing: border-box;

  /* 现代 CSS 特性：强制将其声明为一个独立布局上下文容器，阻止内部溢出影响全局 */
  container-type: inline-size;
}

/* 1. 上层按钮组区域（刚性块，宽度死死咬住 100%） */
.editor-control-panel {
  width: 100%;
  max-width: 100%;
  background-color: #ffffff;
  border-bottom: 1px solid var(--el-border-color-light);
  box-sizing: border-box;
  padding: 0 20px;
  display: flex;
  align-items: center;
  overflow: hidden;

  .panel-left {
    display: flex;
    align-items: center;
    gap: 12px;
  }
}

/* 2. 下层设计器视口区 
  【极关键】：Grid 布局下的子项必须配合 min-width: 0 和 min-height: 0
  否则其尺寸会被内部庞大的 pdfme Canvas 节点强行撑开至 2500px+
*/
.editor-viewport-body {
  width: 100%;
  max-width: 100%;
  min-width: 0;
  min-height: 0;
  padding: 16px;
  box-sizing: border-box;
  background-color: #f5f7fa;
  overflow: hidden;
}

/* 视口内部单项视图隔离区 */
.render-view-isolation {
  width: 100%;
  height: 100%;
  min-width: 0;
  min-height: 0;
  background: #ffffff;
  border-radius: 4px;
  overflow: hidden;
}

/* 强力镇压 pdfme 渲染引擎向外、向下无节制撑大页面的宿主节点 */
.pdfme-native-mount-point {
  width: 100% !important;
  height: 100% !important;
  max-width: 100% !important;
  max-height: 100% !important;
  overflow: hidden !important;

  /* 穿透覆盖 pdfme-designer 自身可能携带的巨幅物理宽度 */
  :deep(.pdfme-designer) {
    width: 100% !important;
    height: 100% !important;
    max-width: 100% !important;
    max-height: 100% !important;
  }
}

/* JSON 模式下的独立纵向排版 */
.json-mode-layout {
  display: flex;
  flex-direction: column;
  padding: 16px;
  box-sizing: border-box;

  .editor-tip {
    padding: 10px 14px;
    font-size: 13px;
    color: #e6a23c;
    background-color: #fdf6ec;
    border-radius: 4px;
    margin-bottom: 12px;
    flex-shrink: 0;
  }

  .textarea-container {
    flex: 1;
    width: 100%;
    min-height: 0;
    /* 锁住 textarea，不让其撑开 */
  }

  .json-pure-textarea {
    height: 100%;
    font-family: 'Courier New', Courier, monospace;

    :deep(.el-textarea__inner) {
      height: 100% !important;
      resize: none;
      background-color: #fafafa;
      color: #333333;
      padding: 12px;
    }
  }
}
</style>