<template>
  <TablePreview
    v-model="tableOptions"
    @complete="handleComplete"
    @cancel="handleCancel"
    @export-success="handleExportSuccess"
    @export-error="handleExportError"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import TablePreview from '../preview.vue'

const tableOptions = ref({
  columns: [
    { field: '0', title: '名称' },
    { field: '1', title: '年龄' },
    { field: '2', title: '性别' },
    { field: '3', title: '爱好' },
    { field: '4', title: '项次' },
  ],
  headers: [
    { field: '0', caption: 'A' },
    { field: '1', caption: 'B' },
    { field: '2', caption: 'C' },
    { field: '3', caption: 'D' },
    { field: '4', caption: 'E' },
  ],
  records: Array.from({ length: 1000 }, (v: unknown, k: number) => ({
    0: '张三',
    1: 18 + Math.floor(Math.random() * 20),
    2: Math.random() > 0.5 ? '男' : '女',
    3: '🏀',
    4: k,
  }))
})

const handleComplete = () => {
  console.log('完成')
}

const handleCancel = () => {
  console.log('取消')
}

const handleExportSuccess = () => {
  console.log('导出成功')
}

const handleExportError = (error: Error) => {
  console.error('导出失败:', error)
}
</script>
