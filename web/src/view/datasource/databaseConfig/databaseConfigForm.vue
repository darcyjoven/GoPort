
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="数据库名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入数据库名称" />
</el-form-item>
        <el-form-item label="数据库类型:" prop="dbType">
    <el-tree-select v-model="formData.dbType" placeholder="请选择数据库类型" :data="databaseOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="主机地址/IP:" prop="host">
    <el-input v-model="formData.host" :clearable="true" placeholder="请输入主机地址/IP" />
</el-form-item>
        <el-form-item label="端口号:" prop="port">
    <el-input v-model.number="formData.port" :clearable="true" placeholder="请输入端口号" />
</el-form-item>
        <el-form-item label="数据库名:" prop="server">
    <el-input v-model="formData.server" :clearable="true" placeholder="请输入数据库名" />
</el-form-item>
        <el-form-item label="用户名:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户名" />
</el-form-item>
        <el-form-item label="密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="其它参数:" prop="extraParams">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.extraParams 后端会按照json的类型进行存取
    {{ formData.extraParams }}
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="是否启用:" prop="enable">
    <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createDatabaseConfig,
  updateDatabaseConfig,
  findDatabaseConfig
} from '@/api/datasource/databaseConfig'

defineOptions({
    name: 'DatabaseConfigForm'
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
const databaseOptions = ref([])
const formData = ref({
            name: '',
            dbType: '',
            host: '',
            port: undefined,
            server: '',
            username: '',
            password: '',
            extraParams: {},
            remark: '',
            enable: false,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dbType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               host : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               enable : [{
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
      const res = await findDatabaseConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    databaseOptions.value = await getDictFunc('database')
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
               res = await createDatabaseConfig(formData.value)
               break
             case 'update':
               res = await updateDatabaseConfig(formData.value)
               break
             default:
               res = await createDatabaseConfig(formData.value)
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
