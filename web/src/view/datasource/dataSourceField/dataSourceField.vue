
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="数据源类型" prop="sourceType">
    <el-tree-select v-model="searchInfo.sourceType" placeholder="请选择数据源类型" :data="datasource_typeOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
            
            <el-form-item label="字段顺序" prop="fieldIndex">
  <el-input v-model.number="searchInfo.fieldIndex" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="字段类型" prop="filedType">
    <el-tree-select v-model="searchInfo.filedType" placeholder="请选择字段类型" :data="data_typeOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
            
            <el-form-item label="字段描述" prop="description">
  <el-input v-model="searchInfo.description" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="字段别名" prop="fieldName">
  <el-input v-model="searchInfo.fieldName" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="排序" prop="sortable">
    <el-tree-select v-model="searchInfo.sortable" placeholder="请选择排序" :data="sortOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
          
          <el-form-item label="宽度" prop="width">
  <el-input v-model.number="searchInfo.width" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="格式化" prop="format">
  <el-input v-model="searchInfo.format" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="换行" prop="warp">
  <el-select v-model="searchInfo.warp" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
          
          <el-form-item label="对齐方式" prop="align">
    <el-tree-select v-model="searchInfo.align" placeholder="请选择对齐方式" :data="alignOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
          
          <el-form-item label="其它配置" prop="extra">
  <el-input v-model="searchInfo.extra" placeholder="搜索条件" />
</el-form-item>
          
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column sortable align="left" label="名称" prop="name" width="120" />

            <el-table-column sortable align="left" label="数据源类型" prop="sourceType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.sourceType,datasource_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="数据源ID" prop="sourceID" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.sourceID,scope.row.sourceID) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="字段顺序" prop="fieldIndex" width="120" />

            <el-table-column align="left" label="原始字段内容" prop="fieldKey" width="120" />

            <el-table-column align="left" label="字段类型" prop="filedType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.filedType,data_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="字段别名" prop="fieldName" width="120" />

            <el-table-column align="left" label="字段描述" prop="description" width="120" />

            <el-table-column align="left" label="排序" prop="sortable" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.sortable,sortOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="宽度" prop="width" width="120" />

            <el-table-column align="left" label="格式化" prop="format" width="120" />

            <el-table-column align="left" label="换行" prop="warp" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.warp) }}</template>
</el-table-column>
            <el-table-column align="left" label="对齐方式" prop="align" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.align,alignOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="其它配置" prop="extra" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateDataSourceFieldFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="数据源类型">
    {{ detailForm.sourceType }}
</el-descriptions-item>
                    <el-descriptions-item label="数据源ID">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.sourceID,detailForm.sourceID) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="字段顺序">
    {{ detailForm.fieldIndex }}
</el-descriptions-item>
                    <el-descriptions-item label="原始字段内容">
    {{ detailForm.fieldKey }}
</el-descriptions-item>
                    <el-descriptions-item label="字段类型">
    {{ detailForm.filedType }}
</el-descriptions-item>
                    <el-descriptions-item label="字段别名">
    {{ detailForm.fieldName }}
</el-descriptions-item>
                    <el-descriptions-item label="字段描述">
    {{ detailForm.description }}
</el-descriptions-item>
                    <el-descriptions-item label="排序">
    {{ detailForm.sortable }}
</el-descriptions-item>
                    <el-descriptions-item label="宽度">
    {{ detailForm.width }}
</el-descriptions-item>
                    <el-descriptions-item label="格式化">
    {{ detailForm.format }}
</el-descriptions-item>
                    <el-descriptions-item label="换行">
    {{ detailForm.warp }}
</el-descriptions-item>
                    <el-descriptions-item label="对齐方式">
    {{ detailForm.align }}
</el-descriptions-item>
                    <el-descriptions-item label="其它配置">
    {{ detailForm.extra }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getDataSourceFieldDataSource,
  createDataSourceField,
  deleteDataSourceField,
  deleteDataSourceFieldByIds,
  updateDataSourceField,
  findDataSourceField,
  getDataSourceFieldList
} from '@/api/datasource/dataSourceField'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'DataSourceField'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getDataSourceFieldDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               sourceType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               sourceID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fieldIndex : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               filedType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               fieldName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               sortable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               warp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            name: 'name',
            sourceType: 'source_type',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.warp === ""){
        searchInfo.value.warp=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getDataSourceFieldList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    sortOptions.value = await getDictFunc('sort')
    datasource_typeOptions.value = await getDictFunc('datasource_type')
    alignOptions.value = await getDictFunc('align')
    data_typeOptions.value = await getDictFunc('data_type')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteDataSourceFieldFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteDataSourceFieldByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDataSourceFieldFunc = async(row) => {
    const res = await findDataSourceField({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDataSourceFieldFunc = async (row) => {
    const res = await deleteDataSourceField({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findDataSourceField({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
