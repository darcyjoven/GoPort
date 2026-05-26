<template>
  <div>
    <PdfmeViewer :title="pageTitle" :template-data="currentTemplate" :inputs-data="businessInputs">
      <template #extra-actions>
        <el-button type="success" size="small" @click="changeMockData">
          切换另一组单据数据
        </el-button>
      </template>
    </PdfmeViewer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import PdfmeViewer from '../preview.vue' // 引入刚封装的组件
import type { Template } from '@pdfme/common'
import { ElMessage } from 'element-plus'

const pageTitle = ref('出库单高保真数据预览')

// 1. 表格/凭证模板结构配置 (通常从后端 API 拉取)
const currentTemplate = ref<Template>({
  basePdf: { width: 210, height: 297, padding: [2, 2, 2, 2] },
  schemas: [
    [
      {
        name: 'title',
        type: 'text',
        position: { x: 20, y: 20 },
        width: 170,
        height: 15,
        fontSize: 24,
        alignment: 'center',
        content: '出库凭证报表'
      },
      {
        name: 'orderNo',
        type: 'text',
        position: { x: 20, y: 45 },
        width: 100,
        height: 10,
        fontSize: 12,
        content: '订单编号: '
      }
    ]
  ]
})

// 2. 纯业务动态键值对数据集 (解耦存放，可随时变更)
const businessInputs = ref<Record<string, any>[]>(
  [
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
    { title: '销货出库单 (TS 级精准打印)', orderNo: '订单编号: SO-20260525-8888' },
  ])

// 模拟父组件业务动作：改变数据流，子组件会自动无闪烁热刷新
const changeMockData = () => {
  currentTemplate.value = {
    basePdf: { width: 210, height: 297, padding: [2, 2, 2, 2] },
    schemas: [
      [
        {
          name: 'name',
          type: 'text',
          content: 'Type Something...',
          position: { x: 21.04, y: 19.67 },
          width: 45,
          height: 10,
          rotate: 0,
          alignment: 'left',
          verticalAlignment: 'top',
          fontSize: 13,
          textFormat: 'plain',
          overflow: 'visible',
          fontVariantFallback: 'synthetic',
          lineHeight: 1,
          characterSpacing: 0,
          fontColor: '#000000',
          fontName: '仿宋',
          backgroundColor: '',
          borderColor: '#000000',
          borderWidth: { top: 0, right: 0, bottom: 0, left: 0 },
          padding: { top: 0, right: 0, bottom: 0, left: 0 },
          opacity: 1,
          strikethrough: false,
          underline: false,
          required: false,
          readOnly: false
        },
        {
          name: 'qrcode',
          type: 'qrcode',
          content: 'https://pdfme.com/',
          position: { x: 21.97, y: 48.99 },
          backgroundColor: '#ffffff',
          barColor: '#000000',
          width: 30,
          height: 30,
          rotate: 0,
          opacity: 1,
          required: false,
          readOnly: false
        },
        {
          name: 'table',
          type: 'table',
          position: { x: 21.67, y: 113.06 },
          width: 150,
          height: 57.5184,
          // 修复为原生的 JS 二维数组对象，免去 JSON.parse 困扰
          content: "",
          showHead: true,
          repeatHead: false,
          head: ['Name', 'City', 'Description'],
          headWidthPercentages: [30, 30, 40],
          tableStyles: { borderWidth: 0.3, borderColor: '#000000' },
          headStyles: {
            fontName: '黑体',
            fontSize: 13,
            characterSpacing: 0,
            alignment: 'left',
            verticalAlignment: 'middle',
            lineHeight: 1,
            fontColor: '#ffffff',
            borderColor: '',
            backgroundColor: '#2980ba',
            borderWidth: { top: 0, right: 0, bottom: 0, left: 0 },
            padding: { top: 5, right: 5, bottom: 5, left: 5 }
          },
          bodyStyles: {
            fontName: '黑体',
            fontSize: 13,
            characterSpacing: 0,
            alignment: 'left',
            verticalAlignment: 'middle',
            lineHeight: 1,
            fontColor: '#000000',
            borderColor: '#888888',
            backgroundColor: '',
            alternateBackgroundColor: '#f5f5f5',
            borderWidth: { top: 0.1, right: 0.1, bottom: 0.1, left: 0.1 },
            padding: { top: 5, right: 5, bottom: 5, left: 5 }
          },
          columnStyles: {},
          required: false,
          readOnly: false
        }
      ]
    ]
  }
  businessInputs.value = [
    {
      name: '技术部 - 张三',
      qrcode: 'https://github.com/gva-org/gin-vue-admin',
      table: [
        ['李四', '北京', '全栈架构师，负责ERP核心模块重构'],
        ['王五', '上海', '前端专家，精通Vue3与数据可视化画布']
      ]
    },
    {
      name: '销售部 - 财务对账单',
      qrcode: 'https://gin-vue-admin.com',
      table: [
        ['华为终端', '深圳', '采购Mate 60 Pro等办公设备，共计50台'],
        ['小米科技', '北京', '定制平板设备测试机交付，已开票']
      ]
    },
    {
      name: '供应链 - 出库单据',
      qrcode: 'https://pdfme.com',
      table: [
        ['A001号主板', '无锡仓', '核心控制芯片级元器件，出库2000件'],
        ['B007号外壳', '苏州仓', '防静电定制外壳，加急调配800箱']
      ]
    },
    {
      name: '人事部 - 季度绩效公示',
      qrcode: 'https://element-plus.org',
      table: [
        ['赵六', '广州', '跨境电商运营总监，季度KPI达成率120%'],
        ['孙七', '成都', '西南大区销售经理，新客户开拓破纪录']
      ]
    },
    {
      name: '质检部 - 抽样报告',
      qrcode: 'https://code.visualstudio.com',
      table: [
        ['常规金属件', '一号线', '耐腐蚀压力测试通过，指标优于行业标准'],
        ['绝缘塑胶垫', '三号线', '边缘轻微毛刺，已责令车间现场整改']
      ]
    }
  ]
  ElMessage.success('成功切换至最新业务流水数据！')
}
</script>