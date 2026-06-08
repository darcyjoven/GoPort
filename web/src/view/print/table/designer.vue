<template>
  <div class="vtable-designer">
    <!-- 顶部工具栏 -->
    <div class="toolbar">
      <div class="left-controls">
        <el-switch v-model="showJsonEditor" inactive-text="JSON 编辑器" @change="handleEditorVisibilityChange" />
      </div>

      <div class="right-actions">
        <slot name="toolbar">
          <el-button type="info" size="small" plain @click="$emit('reset')">重置</el-button>
          <el-button type="primary" size="small" plain @click="handleComplete">完成</el-button>
          <el-button type="warning" size="small" plain @click="$emit('cancel')">取消</el-button>
        </slot>
      </div>

      <div class="editor-tip">
        💡 提示：{{ showJsonEditor ? '修改 JSON 后失焦同步到表格，聚焦时从表格拉取' : '当前已关闭 JSON 视图，专注于表格设计' }}
      </div>
    </div>

    <!-- 主体区域 -->
    <div class="main-content">
      <!-- 左侧：配置编辑器 -->
      <div v-if="showJsonEditor" class="left-panel" :style="{ width: leftWidth + '%' }">
        <el-tabs v-model="activeTab" type="border-card">
          <el-tab-pane label="配置项" name="config">
            <div class="editor-container">
              <div ref="configEditorRef" class="monaco-editor"></div>
            </div>
          </el-tab-pane>
          <el-tab-pane label="数据" name="data">
            <div class="editor-container">
              <div ref="dataEditorRef" class="monaco-editor"></div>
            </div>
          </el-tab-pane>
        </el-tabs>

        <div v-if="jsonError" class="error-message">{{ jsonError }}</div>
      </div>

      <!-- 分隔线 -->
      <div v-if="showJsonEditor" class="resizer" :class="{ dragging: isDragging }" @mousedown="startDrag">
        <span class="resize-icon">⋮⋮</span>
      </div>

      <!-- 右侧：预览 -->
      <div class="right-panel" :style="{ width: showJsonEditor ? (100 - leftWidth - 0.5) + '%' : '100%' }">
        <div ref="tableContainer" class="table-wrapper"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { ListTable } from '@visactor/vtable'
import { ElMessage } from 'element-plus'
import * as monaco from 'monaco-editor'

// --- Props ---
const props = defineProps<{
  modelValue?: {
    config: string
    data: string
  }
}>()

// --- Emits ---
const emit = defineEmits<{
  (e: 'update:modelValue', value: { config: string; data: string }): void
  (e: 'reset'): void
  (e: 'cancel'): void
  (e: 'complete', data: { config: string; data: string }): void
}>()

// --- State ---
const showJsonEditor = ref(false)
const activeTab = ref('config')
const leftWidth = ref(40)
const isDragging = ref(false)

const defaultConfigJson = JSON.stringify({
  columns: [
    { field: 'name', title: '名称', sort: true },
    { field: 'age', title: '年龄', sort: true },
    { field: 'gender', title: '性别', sort: true },
    { field: 'hobby', title: '爱好', sort: true }
  ]
}, null, 2)

const defaultDataJson = JSON.stringify([
  { name: '张三', age: 25, gender: '男', hobby: '篮球' },
  { name: '李四', age: 30, gender: '女', hobby: '游泳' },
  { name: '王五', age: 28, gender: '男', hobby: '足球' }
], null, 2)

const jsonError = ref('')
const tableContainer = ref<HTMLElement | null>(null)
const configEditorRef = ref<HTMLElement | null>(null)
const dataEditorRef = ref<HTMLElement | null>(null)

let tableInstance: ListTable | null = null
let configEditor: monaco.editor.IStandaloneCodeEditor | null = null
let dataEditor: monaco.editor.IStandaloneCodeEditor | null = null

interface VTableConfig {
  columns: Array<{ field: string; title: string }>
}

// --- Methods ---
const validateJson = (jsonString: string) => {
  if (!jsonString.trim()) return false
  try {
    JSON.parse(jsonString)
    return true
  } catch {
    return false
  }
}

const updateTable = (showMessage = false) => {
  if (!configEditor || !dataEditor) return

  const configValue = configEditor.getValue()
  const dataValue = dataEditor.getValue()

  if (!validateJson(configValue)) {
    jsonError.value = '配置JSON格式错误'
    return
  }
  if (!validateJson(dataValue)) {
    jsonError.value = '数据JSON格式错误'
    return
  }

  jsonError.value = ''

  try {
    const config: VTableConfig = JSON.parse(configValue)
    const data = JSON.parse(dataValue)

    if (tableInstance) {
      tableInstance.release()
      tableInstance = null
    }

    if (tableContainer.value) {
      tableInstance = new ListTable({
        container: tableContainer.value,
        columns: config.columns,
        records: data,
        multipleSort: true
      })
    }

    emit('update:modelValue', { config: configValue, data: dataValue })

    if (showMessage) {
      ElMessage({ message: 'JSON 修改成功，已同步到表格', type: 'success', duration: 1500 })
    }
  } catch (e) {
    jsonError.value = (e as Error).message
  }
}

const createEditor = (container: HTMLElement, initialValue: string) => {
  return monaco.editor.create(container, {
    value: initialValue,
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
}

const initEditors = () => {
  const configValue = props.modelValue?.config || defaultConfigJson
  const dataValue = props.modelValue?.data || defaultDataJson

  if (configEditorRef.value && !configEditor) {
    configEditor = createEditor(configEditorRef.value, configValue)
    configEditor.onDidBlurEditorText(() => {
      updateTable(true)
    })
  }

  if (dataEditorRef.value && !dataEditor) {
    dataEditor = createEditor(dataEditorRef.value, dataValue)
    dataEditor.onDidBlurEditorText(() => {
      updateTable(true)
    })
  }
}

const disposeEditors = () => {
  if (configEditor) {
    configEditor.dispose()
    configEditor = null
  }
  if (dataEditor) {
    dataEditor.dispose()
    dataEditor = null
  }
}

const handleEditorVisibilityChange = (visible: boolean | string | number) => {
  if (visible) {
    nextTick(() => {
      initEditors()
    })
  } else {
    disposeEditors()
  }
}

const handleComplete = () => {
  if (showJsonEditor.value && configEditor && dataEditor) {
    updateTable(false)
  }

  let configValue = defaultConfigJson
  let dataValue = defaultDataJson

  if (showJsonEditor.value && configEditor && dataEditor) {
    configValue = configEditor.getValue()
    dataValue = dataEditor.getValue()
  } else if (tableInstance) {
    // 从表格实例获取数据
  }

  emit('complete', { config: configValue, data: dataValue })
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
}

const stopDrag = () => {
  isDragging.value = false
  document.removeEventListener('mousemove', onDrag)
  document.removeEventListener('mouseup', stopDrag)
}

// --- Watch ---
watch(() => props.modelValue, (newVal) => {
  if (newVal && configEditor && dataEditor) {
    const currentConfig = configEditor.getValue()
    const currentData = dataEditor.getValue()

    if (currentConfig !== newVal.config) {
      configEditor.setValue(newVal.config)
    }
    if (currentData !== newVal.data) {
      dataEditor.setValue(newVal.data)
    }

    // 直接更新表格，不调用 updateTable 避免循环
    try {
      const config: VTableConfig = JSON.parse(newVal.config)
      const data = JSON.parse(newVal.data)

      if (tableInstance) {
        tableInstance.release()
        tableInstance = null
      }

      if (tableContainer.value) {
        tableInstance = new ListTable({
          container: tableContainer.value,
          columns: config.columns,
          records: data,
          multipleSort: true
        })
      }
    } catch (e) {
      console.error('Failed to update table from modelValue:', e)
    }
  }
}, { deep: true })

// --- Lifecycle ---
onMounted(() => {
  nextTick(() => {
    const configValue = props.modelValue?.config || defaultConfigJson
    const dataValue = props.modelValue?.data || defaultDataJson

    if (tableContainer.value) {
      const config: VTableConfig = JSON.parse(configValue)
      const data = JSON.parse(dataValue)

      tableInstance = new ListTable({
        container: tableContainer.value,
        columns: config.columns,
        records: data,
        multipleSort: true
      })
    }
  })
})

onUnmounted(() => {
  if (tableInstance) {
    tableInstance.release()
    tableInstance = null
  }
  disposeEditors()
})
</script>

<style scoped>
.vtable-designer {
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
  gap: 8px;
  overflow: hidden;
  height: 100%;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  transition: width 0.1s ease;
}

.left-panel :deep(.el-tabs) {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.left-panel :deep(.el-tabs__content) {
  flex: 1;
  overflow: hidden;
}

.left-panel :deep(.el-tab-pane) {
  height: 100%;
  display: flex;
  flex-direction: column;
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

.resizer:hover .resize-icon,
.resizer.dragging .resize-icon {
  color: #fff;
}

.editor-container {
  flex: 1;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  overflow: hidden;
  min-height: 0;
}

.monaco-editor {
  width: 100%;
  height: 100%;
}

.error-message {
  color: #f56c6c;
  font-size: 12px;
  padding: 8px;
}

.right-panel {
  min-height: 0;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  overflow: hidden;
  transition: width 0.1s ease;
}

.table-wrapper {
  width: 100%;
  height: 100%;
}
</style>
