
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="文件源ID:" prop="sourceID">
    <el-select v-model="formData.sourceID" placeholder="请选择文件源ID" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.sourceID" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="序列:" prop="index">
    <el-input v-model.number="formData.index" :clearable="true" placeholder="请输入序列" />
</el-form-item>
        <el-form-item label="原始字段值:" prop="key">
    <el-input v-model="formData.key" :clearable="true" placeholder="请输入原始字段值" />
</el-form-item>
        <el-form-item label="字段名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入字段名称" />
</el-form-item>
        <el-form-item label="说明:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入说明" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getFileFieldDataSource,
  createFileField,
  updateFileField,
  findFileField
} from '@/api/datasource/fileField'

defineOptions({
    name: 'FileFieldForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            sourceID: undefined,
            index: undefined,
            key: '',
            name: '',
            description: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getFileFieldDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFileField({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createFileField(formData.value)
               break
             case 'update':
               res = await updateFileField(formData.value)
               break
             default:
               res = await createFileField(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
