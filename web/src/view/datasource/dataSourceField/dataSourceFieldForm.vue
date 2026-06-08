
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="数据源类型:" prop="sourceType">
    <el-tree-select v-model="formData.sourceType" placeholder="请选择数据源类型" :data="datasource_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="数据源ID:" prop="sourceID">
    <el-select v-model="formData.sourceID" placeholder="请选择数据源ID" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.sourceID" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="字段顺序:" prop="fieldIndex">
    <el-input v-model.number="formData.fieldIndex" :clearable="true" placeholder="请输入字段顺序" />
</el-form-item>
        <el-form-item label="原始字段内容:" prop="fieldKey">
    <el-input v-model="formData.fieldKey" :clearable="true" placeholder="请输入原始字段内容" />
</el-form-item>
        <el-form-item label="字段类型:" prop="filedType">
    <el-tree-select v-model="formData.filedType" placeholder="请选择字段类型" :data="data_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="字段别名:" prop="fieldName">
    <el-input v-model="formData.fieldName" :clearable="true" placeholder="请输入字段别名" />
</el-form-item>
        <el-form-item label="字段描述:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入字段描述" />
</el-form-item>
        <el-form-item label="排序:" prop="sortable">
    <el-tree-select v-model="formData.sortable" placeholder="请选择排序" :data="sortOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="宽度:" prop="width">
    <el-input-number v-model="formData.width" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="格式化:" prop="format">
    <el-input v-model="formData.format" :clearable="true" placeholder="请输入格式化" />
</el-form-item>
        <el-form-item label="换行:" prop="warp">
    <el-switch v-model="formData.warp" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="对齐方式:" prop="align">
    <el-tree-select v-model="formData.align" placeholder="请选择对齐方式" :data="alignOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="其它配置:" prop="extra">
    <el-input v-model="formData.extra" :clearable="true" placeholder="请输入其它配置" />
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
    getDataSourceFieldDataSource,
  createDataSourceField,
  updateDataSourceField,
  findDataSourceField
} from '@/api/datasource/dataSourceField'

defineOptions({
    name: 'DataSourceFieldForm'
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
const sortOptions = ref([])
const datasource_typeOptions = ref([])
const alignOptions = ref([])
const data_typeOptions = ref([])
const formData = ref({
            name: '',
            sourceType: '',
            sourceID: undefined,
            fieldIndex: undefined,
            fieldKey: '',
            filedType: '',
            fieldName: '',
            description: '',
            sortable: '',
            width: 0,
            format: '',
            warp: false,
            align: '',
            extra: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fieldIndex : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               filedType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fieldName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sortable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               warp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getDataSourceFieldDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDataSourceField({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    sortOptions.value = await getDictFunc('sort')
    datasource_typeOptions.value = await getDictFunc('datasource_type')
    alignOptions.value = await getDictFunc('align')
    data_typeOptions.value = await getDictFunc('data_type')
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
               res = await createDataSourceField(formData.value)
               break
             case 'update':
               res = await updateDataSourceField(formData.value)
               break
             default:
               res = await createDataSourceField(formData.value)
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
