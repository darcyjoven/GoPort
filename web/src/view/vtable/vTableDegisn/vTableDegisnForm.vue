
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="说明:" prop="desciption">
    <el-input v-model="formData.desciption" :clearable="true" placeholder="请输入说明" />
</el-form-item>
        <el-form-item label="模块:" prop="module">
    <el-tree-select v-model="formData.module" placeholder="请选择模块" :data="moudlesOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="当前版本:" prop="currentVersion">
    <el-input v-model="formData.currentVersion" :clearable="true" placeholder="请输入当前版本" />
</el-form-item>
        <el-form-item label="签出否:" prop="checkout">
    <el-switch v-model="formData.checkout" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="签出版本:" prop="checkoutVersion">
    <el-input v-model="formData.checkoutVersion" :clearable="true" placeholder="请输入签出版本" />
</el-form-item>
        <el-form-item label="上次签出时间:" prop="lastCheckOutTime">
    <el-date-picker v-model="formData.lastCheckOutTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="生效否:" prop="active">
    <el-switch v-model="formData.active" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createVTableDesign,
  updateVTableDesign,
  findVTableDesign
} from '@/api/vtable/vTableDegisn'

defineOptions({
    name: 'VTableDesignForm'
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
const moudlesOptions = ref([])
const formData = ref({
            name: '',
            desciption: '',
            module: '',
            currentVersion: '',
            checkout: false,
            checkoutVersion: '',
            lastCheckOutTime: new Date(),
            remark: '',
            active: false,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               module : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               currentVersion : [{
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
      const res = await findVTableDesign({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    moudlesOptions.value = await getDictFunc('moudles')
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
               res = await createVTableDesign(formData.value)
               break
             case 'update':
               res = await updateVTableDesign(formData.value)
               break
             default:
               res = await createVTableDesign(formData.value)
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
