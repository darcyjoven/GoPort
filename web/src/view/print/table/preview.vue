<template>
  <div class="vtable-preview">
    <!-- 顶部工具栏 -->
    <div class="toolbar">
      <slot name="toolbar">
        <el-button size="small" @click="goBack"> <el-icon>
            <ArrowLeft />
          </el-icon> 返回 </el-button>
        <el-button type="primary" size="small" plain @click="$emit('complete')">完成</el-button>
        <el-button type="success" size="small" plain @click="handleExport">导出Excel</el-button>
        <el-button type="warning" size="small" plain @click="$emit('cancel')">取消</el-button>
      </slot>
    </div>

    <!-- 表格区域 -->
    <div ref="tableContainer" class="table-wrapper"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { ListTable } from '@visactor/vtable'
import { exportVTableToExcel, downloadExcel } from '@visactor/vtable-export'

// --- Props ---
const props = defineProps<{
  modelValue?: {
    columns: Array<{ field: string; title: string }>
    records: Array<Record<string, unknown>>
    headers?: Array<{ field: string; caption: string }>
  }
}>()

const goBack = () => {
  history.back()
}

// --- Emits ---
const emit = defineEmits<{
  (e: 'complete'): void
  (e: 'cancel'): void
  (e: 'export-success'): void
  (e: 'export-error', error: Error): void
}>()

// --- State ---
const tableContainer = ref<HTMLElement | null>(null)
let tableInstance: ListTable | null = null

// --- Methods ---
const initTable = () => {
  if (!tableContainer.value || tableInstance) return

  const { columns = [], records = [], headers } = props.modelValue || {}

  tableInstance = new ListTable({
    container: tableContainer.value,
    columns,
    header: headers,
    records,
    defaultRowHeight: 36
  })
}

const handleExport = async () => {
  try {
    if (!tableInstance) {
      throw new Error('表格实例未找到')
    }

    const exportData = await exportVTableToExcel(tableInstance)
    await downloadExcel(exportData, '表格数据')

    emit('export-success')
  } catch (error) {
    emit('export-error', error as Error)
  }
}

// --- Watch ---
watch(() => props.modelValue, (newVal) => {
  if (newVal && tableInstance) {
    tableInstance.release()
    tableInstance = null

    const { columns = [], records = [], headers } = newVal

    if (tableContainer.value) {
      tableInstance = new ListTable({
        container: tableContainer.value,
        columns,
        header: headers,
        records,
        defaultRowHeight: 36
      })
    }
  }
}, { deep: true })

// --- Lifecycle ---
onMounted(() => {
  initTable()
})

onUnmounted(() => {
  if (tableInstance) {
    tableInstance.release()
    tableInstance = null
  }
})
</script>

<style scoped>
.vtable-preview {
  width: 100%;
  height: 100%;
  display: flex;
  padding: 0 20px;
  flex-direction: column;
  gap: 12px;
}

.toolbar {
  display: flex;
  gap: 12px;
  padding: 8px 0 0 0;
}

.table-wrapper {
  flex: 1;
  min-height: 0;
  border: 1px solid #ebf0f5;
  border-radius: 4px;
}
</style>
