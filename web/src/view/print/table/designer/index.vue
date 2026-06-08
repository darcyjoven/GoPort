<template>
  <TableDesigner v-model="tableData" @reset="handleReset" @cancel="handleCancel" @complete="handleComplete" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import TableDesigner from '../designer.vue'

const tableData = ref({
  config: JSON.stringify({
    columns: [
      { field: 'name', title: '名称', sort: true },
      { field: 'age', title: '年龄', sort: true },
      { field: 'gender', title: '性别' },
      { field: 'hobby', title: '爱好' }
    ]
  }, null, 2),
  data: JSON.stringify([
    { name: '张三', age: 25, gender: '男', hobby: '篮球' },
    { name: '李四', age: 30, gender: '女', hobby: '游泳' },
    { name: '王五', age: 28, gender: '男', hobby: '足球' }
  ], null, 2)
})

const handleReset = () => {
  tableData.value = {
    config: JSON.stringify({
      columns: [
        { field: 'name', title: '名称' },
        { field: 'age', title: '年龄' },
        { field: 'gender', title: '性别' },
        { field: 'hobby', title: '爱好' }
      ]
    }, null, 2),
    data: JSON.stringify([
      { name: '张三', age: 25, gender: '男', hobby: '篮球' },
      { name: '李四', age: 30, gender: '女', hobby: '游泳' },
      { name: '王五', age: 28, gender: '男', hobby: '足球' }
    ], null, 2)
  }
}

const handleCancel = () => {
  console.log('取消')
}

const handleComplete = (data: { config: string; data: string }) => {
  console.log('完成:', data)
}
</script>
