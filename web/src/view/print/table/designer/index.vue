<template>
  <div class="vtable-container">
    <!-- 顶部工具栏 -->
    <div class="toolbar">
      <el-button type="info" size="small" plain @click="handleReset">重置</el-button>
      <el-button type="primary" size="small" plain @click="handleComplete">完成</el-button>
      <el-button type="warning" size="small" plain @click="handleCancel">取消</el-button>
    </div>

    <!-- 主体区域 -->
    <div class="main-content">
      <!-- 左侧：配置编辑器 -->
      <div class="left-panel" :style="{ width: leftWidth + '%' }">
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
      <div 
        class="resizer" 
        :class="{ dragging: isDragging }"
        @mousedown="startDrag"
      >
        <span class="resize-icon">⋮⋮</span>
      </div>

      <!-- 右侧：预览 -->
      <div class="right-panel" :style="{ width: (100 - leftWidth - 0.5) + '%' }">
        <div ref="tableContainer" class="table-wrapper"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { ListTable } from '@visactor/vtable'
import * as monaco from 'monaco-editor'

const activeTab = ref('config')
const leftWidth = ref(40)
const isDragging = ref(false)

const defaultConfigJson = JSON.stringify({
  columns: [
    { field: 'name', title: '名称' },
    { field: 'age', title: '年龄' },
    { field: 'gender', title: '性别' },
    { field: 'hobby', title: '爱好' }
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

const validateJson = (jsonString: string) => {
  if (!jsonString.trim()) return false
  try {
    JSON.parse(jsonString)
    return true
  } catch {
    return false
  }
}

const updateTable = () => {
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
        defaultRowHeight: 36
      })
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

const handleReset = () => {
  if (configEditor) {
    configEditor.setValue(defaultConfigJson)
  }
  if (dataEditor) {
    dataEditor.setValue(defaultDataJson)
  }
}

const handleComplete = () => {
  console.log('完成')
}

const handleCancel = () => {
  console.log('取消')
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

onMounted(() => {
  nextTick(() => {
    if (configEditorRef.value) {
      configEditor = createEditor(configEditorRef.value, defaultConfigJson)
      configEditor.onDidChangeModelContent(() => {
        updateTable()
      })
    }

    if (dataEditorRef.value) {
      dataEditor = createEditor(dataEditorRef.value, defaultDataJson)
      dataEditor.onDidChangeModelContent(() => {
        updateTable()
      })
    }

    updateTable()
  })
})

onUnmounted(() => {
  if (tableInstance) {
    tableInstance.release()
    tableInstance = null
  }
  if (configEditor) {
    configEditor.dispose()
    configEditor = null
  }
  if (dataEditor) {
    dataEditor.dispose()
    dataEditor = null
  }
})
</script>

<style scoped>
.vtable-container {
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
}

.main-content {
  flex: 1;
  display: flex;
  gap: 12px;
  min-height: 0;
}

.left-panel {
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow: hidden;
  height: 100%;
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
  width: 4px;
  cursor: col-resize;
  background-color: #dcdfe6;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
  flex-shrink: 0;
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
  flex: 1;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
  min-height: 0;
}

.table-wrapper {
  width: 100%;
  height: 100%;
}
</style>
