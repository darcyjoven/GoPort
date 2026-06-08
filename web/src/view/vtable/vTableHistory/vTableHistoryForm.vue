
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="版本:" prop="version">
    <el-input v-model="formData.version" :clearable="true" placeholder="请输入版本" />
</el-form-item>
        <el-form-item label="样式配置:" prop="config">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.config 后端会按照json的类型进行存取
    {{ formData.config }}
</el-form-item>
        <el-form-item label="默认数据:" prop="defaultData">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.defaultData 后端会按照json的类型进行存取
    {{ formData.defaultData }}
</el-form-item>
        <el-form-item label="数据源ID:" prop="sourceID">
    <el-select v-model="formData.sourceID" placeholder="请选择数据源ID" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.sourceID" :key="key" :label="item.label" :value="item.value" />
    </el-select>
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
    getVTableHistoryDataSource,
  createVTableHistory,
  updateVTableHistory,
  findVTableHistory
} from '@/api/vtable/vTableHistory'

defineOptions({
    name: 'VTableHistoryForm'
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
            version: '',
            config: {},
            defaultData: {},
            sourceID: undefined,
            remark: '',
            active: false,
        })
// 验证规则
const rule = reactive({
               version : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getVTableHistoryDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVTableHistory({ ID: route.query.id })
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
               res = await createVTableHistory(formData.value)
               break
             case 'update':
               res = await updateVTableHistory(formData.value)
               break
             default:
               res = await createVTableHistory(formData.value)
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
