<template>
  <div class="vtable-container">
    <!-- 顶部工具栏 -->
    <div class="toolbar">
      <el-button type="primary" size="small" plain @click="handleComplete">完成</el-button>
      <el-button type="success" size="small" plain @click="handleExport">导出Excel</el-button>
      <el-button type="warning" size="small" plain @click="handleCancel">取消</el-button>
    </div>

    <!-- 表格区域 -->
    <div ref="tableContainer" class="table-wrapper"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { ListTable } from '@visactor/vtable';
import { exportVTableToExcel, downloadExcel } from '@visactor/vtable-export';

// 表格容器引用
const tableContainer = ref<HTMLElement | null>(null);

// 表格实例
let tableInstance: ListTable | null = null;

// 定义表格配置
const tableOptions = ref({
  columns: [
    { field: '0', caption: '名称' },
    { field: '1', caption: '年龄' },
    { field: '2', caption: '性别' },
    { field: '3', caption: '爱好' },
    { field: '4', caption: '项次' },
  ],
  headers: [
    { field: '0', caption: 'A' },
    { field: '1', caption: 'B' },
    { field: '2', caption: 'C' },
    { field: '3', caption: 'D' },
    { field: '4', caption: 'E' },
  ],
  pagination: { perPageCount: 20, currentPage: 2 },
  records: Array.from({ length: 1000 }, (v: unknown, k: number) => [
    '张三',
    18 + Math.floor(Math.random() * 20),
    Math.random() > 0.5 ? '男' : '女',
    '🏀',
    k,
  ],)
});

// 初始化表格
const initTable = () => {
  if (!tableContainer.value || tableInstance) return;

  tableInstance = new ListTable({
    container: tableContainer.value,
    columns: tableOptions.value.columns,
    header: tableOptions.value.headers,
    records: tableOptions.value.records,
    // pagination: { perPageCount: 20, currentPage: 2 },
    // autoFillWidth: true,
    // autoFillHeight: true,
    defaultRowHeight: 36
  });
};

// 监听配置变化
watch(tableOptions, (newOptions) => {
  if (tableInstance) {
    tableInstance.updateColumns(newOptions.columns);
    tableInstance.setRecords(newOptions.records);
  }
}, { deep: true });

const handleComplete = () => {
  console.log('完成');
};

const handleCancel = () => {
  console.log('取消');
};

// 导出Excel
const handleExport = async () => {
  try {
    if (!tableInstance) {
      console.error('表格实例未找到');
      return;
    }

    // 导出数据
    const exportData = await exportVTableToExcel(tableInstance);

    // 下载文件
    await downloadExcel(exportData, '表格数据');

    console.log('导出成功');
  } catch (error) {
    console.error('导出失败:', error);
    alert('导出失败，请重试');
  }
};

// 组件挂载时初始化表格
onMounted(() => {
  initTable();
});

// 组件卸载时清理
onUnmounted(() => {
  if (tableInstance) {
    tableInstance.release();
    tableInstance = null;
  }
});
</script>

<style scoped>
.vtable-container {
  width: 100%;
  height: 400px;
  display: flex;
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