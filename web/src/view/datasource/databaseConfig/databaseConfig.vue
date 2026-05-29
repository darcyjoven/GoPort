
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
      
            <el-form-item label="数据库名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="数据库类型" prop="dbType">
    <el-tree-select v-model="searchInfo.dbType" placeholder="请选择数据库类型" :data="databaseOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
            
            <el-form-item label="是否启用" prop="enable">
  <el-select v-model="searchInfo.enable" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="主机地址/IP" prop="host">
  <el-input v-model="searchInfo.host" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="端口号" prop="port">
  <el-input v-model.number="searchInfo.port" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="数据库名" prop="server">
  <el-input v-model="searchInfo.server" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="用户名" prop="username">
  <el-input v-model="searchInfo.username" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="密码" prop="password">
  <el-input v-model="searchInfo.password" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="备注" prop="remark">
  <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="上次测试时间" prop="lastTestTime">
<el-date-picker v-model="searchInfo.lastTestTime" type="datetime" placeholder="搜索条件"></el-date-picker></el-form-item>
          
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
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="datasource_DatabaseConfig" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="datasource_DatabaseConfig" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="datasource_DatabaseConfig" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="数据库名称" prop="name" width="120" />

            <el-table-column sortable align="left" label="数据库类型" prop="dbType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.dbType,databaseOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="主机地址/IP" prop="host" width="120" />

            <el-table-column align="left" label="端口号" prop="port" width="120" />

            <el-table-column align="left" label="数据库名" prop="server" width="120" />

            <el-table-column align="left" label="用户名" prop="username" width="120" />

            <el-table-column align="left" label="密码" prop="password" width="120" />

            <el-table-column label="其它参数" prop="extraParams" width="200">
    <template #default="scope">
        [JSON]
    </template>
</el-table-column>
            <el-table-column align="left" label="备注" prop="remark" width="120" />

            <el-table-column align="left" label="是否启用" prop="enable" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.enable) }}</template>
</el-table-column>
            <el-table-column align="left" label="上次测试时间" prop="lastTestTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastTestTime) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateDatabaseConfigFunc(scope.row)">编辑</el-button>
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="数据库名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="数据库类型">
    {{ detailForm.dbType }}
</el-descriptions-item>
                    <el-descriptions-item label="主机地址/IP">
    {{ detailForm.host }}
</el-descriptions-item>
                    <el-descriptions-item label="端口号">
    {{ detailForm.port }}
</el-descriptions-item>
                    <el-descriptions-item label="数据库名">
    {{ detailForm.server }}
</el-descriptions-item>
                    <el-descriptions-item label="用户名">
    {{ detailForm.username }}
</el-descriptions-item>
                    <el-descriptions-item label="密码">
    {{ detailForm.password }}
</el-descriptions-item>
                    <el-descriptions-item label="其它参数">
    {{ detailForm.extraParams }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
                    <el-descriptions-item label="是否启用">
    {{ detailForm.enable }}
</el-descriptions-item>
                    <el-descriptions-item label="上次测试时间">
    {{ detailForm.lastTestTime }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createDatabaseConfig,
  deleteDatabaseConfig,
  deleteDatabaseConfigByIds,
  updateDatabaseConfig,
  findDatabaseConfig,
  getDatabaseConfigList
} from '@/api/datasource/databaseConfig'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'DatabaseConfig'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               dbType : [{
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
               host : [{
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
               enable : [{
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
            dbType: 'db_type',
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
    if (searchInfo.value.enable === ""){
        searchInfo.value.enable=null
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
  const table = await getDatabaseConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    databaseOptions.value = await getDictFunc('database')
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
            deleteDatabaseConfigFunc(row)
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
      const res = await deleteDatabaseConfigByIds({ IDs })
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
const updateDatabaseConfigFunc = async(row) => {
    const res = await findDatabaseConfig({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDatabaseConfigFunc = async (row) => {
    const res = await deleteDatabaseConfig({ ID: row.ID })
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
        dbType: '',
        host: '',
        port: undefined,
        server: '',
        username: '',
        password: '',
        extraParams: {},
        remark: '',
        enable: false,
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
  const res = await findDatabaseConfig({ ID: row.ID })
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
