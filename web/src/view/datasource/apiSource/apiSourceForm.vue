
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模板名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入模板名称" />
</el-form-item>
        <el-form-item label="api路径:" prop="path">
    <el-input v-model="formData.path" :clearable="true" placeholder="请输入api路径" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="参数1默认值:" prop="argv1">
    <el-input v-model="formData.argv1" :clearable="true" placeholder="请输入参数1默认值" />
</el-form-item>
        <el-form-item label="参数2默认值:" prop="argv2">
    <el-input v-model="formData.argv2" :clearable="true" placeholder="请输入参数2默认值" />
</el-form-item>
        <el-form-item label="参数3默认值:" prop="argv3">
    <el-input v-model="formData.argv3" :clearable="true" placeholder="请输入参数3默认值" />
</el-form-item>
        <el-form-item label="参数4默认值:" prop="argv4">
    <el-input v-model="formData.argv4" :clearable="true" placeholder="请输入参数4默认值" />
</el-form-item>
        <el-form-item label="参数5默认值:" prop="argv5">
    <el-input v-model="formData.argv5" :clearable="true" placeholder="请输入参数5默认值" />
</el-form-item>
        <el-form-item label="参数6默认值:" prop="argv6">
    <el-input v-model="formData.argv6" :clearable="true" placeholder="请输入参数6默认值" />
</el-form-item>
        <el-form-item label="参数7默认值:" prop="argv7">
    <el-input v-model="formData.argv7" :clearable="true" placeholder="请输入参数7默认值" />
</el-form-item>
        <el-form-item label="参数8默认值:" prop="argv8">
    <el-input v-model="formData.argv8" :clearable="true" placeholder="请输入参数8默认值" />
</el-form-item>
        <el-form-item label="参数9默认值:" prop="argv9">
    <el-input v-model="formData.argv9" :clearable="true" placeholder="请输入参数9默认值" />
</el-form-item>
        <el-form-item label="参数10默认值:" prop="argv10">
    <el-input v-model="formData.argv10" :clearable="true" placeholder="请输入参数10默认值" />
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
  createAPISource,
  updateAPISource,
  findAPISource
} from '@/api/datasource/apiSource'

defineOptions({
    name: 'APISourceForm'
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
            name: '',
            path: '',
            remark: '',
            argv1: '',
            argv2: '',
            argv3: '',
            argv4: '',
            argv5: '',
            argv6: '',
            argv7: '',
            argv8: '',
            argv9: '',
            argv10: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               path : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAPISource({ ID: route.query.id })
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
               res = await createAPISource(formData.value)
               break
             case 'update':
               res = await updateAPISource(formData.value)
               break
             default:
               res = await createAPISource(formData.value)
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
