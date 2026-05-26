<template>
  <div class="page-container flex-column">
    <div class="gva-table-box compact-main flex-1 flex-column">

      <div class="compact-header flex-row justify-between align-center">
        <div class="left-action-group flex-row align-center">
          <span class="page-title">单据预览 (TS 版)</span>
          <el-divider direction="vertical" />

          <el-button type="primary" icon="Printer" :loading="printLoading" @click="handlePrint">
            调用浏览器打印
          </el-button>

        </div>

        <div class="right-action-group">
        </div>
      </div>

      <div v-loading="pageLoading" element-loading-text="正在加载字体与渲染预览..." class="viewer-wrapper flex-1">
        <div ref="viewerRootRef" class="viewer-root"></div>
      </div>

    </div>

    <iframe id="pdf-print-iframe" allow="unload" style="display: none;"></iframe>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, toRaw } from 'vue'
import { ElMessage, ElLoading } from 'element-plus'
import { Viewer } from '@pdfme/ui'
import { generate } from '@pdfme/generator'
import type { Template, Font } from '@pdfme/common'
import { autoLoadAllFonts } from '@/utils/fontLoader'
import {
  text, image, table, barcodes, line, rectangle, ellipse, dateTime, select, radioGroup,
  checkbox, multiVariableText, list, svg, signature, date, time
} from '@pdfme/schemas'

// --- 声明必要的命名空间与类型接口 ---
interface MockInputData {
  title: string
  orderNo: string
  [key: string]: string
}

const PDFME_PLUGINS = {
  text, dateTime, date, time,
  multiVariableText, list,
  select, radioGroup, checkbox, line, rectangle, ellipse,
  image, table, svg, signature,
  qrcode: barcodes.qrcode, gs1datamatrix: barcodes.gs1datamatrix,
  ean13: barcodes.ean13, ean8: barcodes.ean8,
  code128: barcodes.code128, code39: barcodes.code39, itf14: barcodes.itf14,
  pdf317: barcodes.pdf417,
}

// 使用符合新版规范的标准 A4 空白 PDF
const initialTemplate: Template = {
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

// 模拟绑定的业务数据
const mockInputs= [
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

// --- 响应式 Ref 变量声明 ---
const viewerRootRef = ref<HTMLDivElement | null>(null)
const viewerInstance = ref<Viewer | null>(null)
const pageLoading = ref<boolean>(false)
const printLoading = ref<boolean>(false)
let globalFontOptions: any = null

/**
 * 初始化并渲染 pdfme 纯 Canvas 预览视图
 */
const initViewer = async (): Promise<void> => {
  if (!viewerRootRef.value) return
  pageLoading.value = true

  try {
    const fontOptions: any = await autoLoadAllFonts()
    globalFontOptions = fontOptions
    console.log('Fonts loaded successfully:', fontOptions)

    const cleanTemplate = JSON.parse(JSON.stringify(toRaw(initialTemplate))) as Template

    viewerInstance.value = new Viewer({
      domContainer: viewerRootRef.value,
      template: cleanTemplate,
      inputs: mockInputs,
      plugins: PDFME_PLUGINS,
      options: {
        font: fontOptions
      }
    })
  } catch (error) {
    console.error('渲染预览失败:', error)
    ElMessage.error('渲染预览失败，请检查 Base64 模板或字体文件')
  } finally {
    pageLoading.value = false
  }
}

/**
 * 调用浏览器静默唤起打印面板
 */
const handlePrint = async (): Promise<void> => {
  if (printLoading.value) return
  printLoading.value = true

  const loading = ElLoading.service({
    lock: true,
    text: '正在生成高保真打印流并唤起面板...',
    background: 'rgba(255, 255, 255, 0.7)'
  })

  try {
    const cleanTemplate = JSON.parse(JSON.stringify(toRaw(initialTemplate))) as Template

    const pdfUint8Array = await generate({
      template: cleanTemplate,
      inputs: mockInputs,
      plugins: PDFME_PLUGINS,
      options: {
        font: globalFontOptions || {}
      }
    })

    const blob = new Blob([pdfUint8Array], { type: 'application/pdf' })
    const pdfUrl = URL.createObjectURL(blob)

    const iframe = document.getElementById('pdf-print-iframe') as HTMLIFrameElement | null
    if (iframe) {
      iframe.src = pdfUrl
      iframe.onload = () => {
        if (iframe.contentWindow) {
          iframe.contentWindow.focus()
          iframe.contentWindow.print()
        }
        setTimeout(() => {
          URL.revokeObjectURL(pdfUrl)
        }, 1500)
      }
    } else {
      throw new Error('未找到专属打印 Iframe 节点')
    }
  } catch (error) {
    console.error('唤起系统打印失败:', error)
    ElMessage.error('打印调度失败，请重新尝试')
  } finally {
    printLoading.value = false
    loading.close()
  }
}

onMounted(() => {
  initViewer()
})

onBeforeUnmount(() => {
  if (viewerInstance.value) {
    try {
      viewerInstance.value.destroy()
    } catch (e) {
      console.warn('销毁实例时捕获警告:', e)
    }
    viewerInstance.value = null
  }
})
</script>

<style scoped lang="scss">
/* 兼容 GVA 主体界面的标准弹性布局 */
.page-container {
  padding: 16px;
  background-color: #f0f2f5;
  height: calc(100vh - 110px);
  /* 扣除 GVA 顶栏及多标签页栏高度 */
  box-sizing: border-box;
}

/* 基础原子布局类 */
.flex-column {
  display: flex;
  flex-direction: column;
}

.flex-row {
  display: flex;
  flex-direction: row;
}

.flex-1 {
  flex: 1;
  min-height: 0;
}

.justify-between {
  justify-content: space-between;
}

.align-center {
  align-items: center;
}

/* 🧠 核心优化：高内聚单体大卡表 */
.compact-main {
  background-color: #ffffff;
  border-radius: 4px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  overflow: hidden;
  /* 防止内部边界溢出 */
  padding: 0 !important;
  /* 彻底移除卡片四周的大边距，改为由内部子模块微调 */

  /* 一体化紧凑控制栏（直接贴顶） */
  .compact-header {
    padding: 12px 16px;
    background-color: #ffffff;
    border-bottom: 1px solid #e4e7ed;
    /* 用轻量线条平滑过渡，消除卡片断层感 */

    .left-action-group {
      gap: 8px;
      /* 按钮横向间距，保证从左往右紧凑排列 */

      .page-title {
        font-size: 14px;
        font-weight: 600;
        color: #303133;
      }
    }
  }

  /* 下预览区域（与控制栏紧密连接） */
  .viewer-wrapper {
    width: 100%;
    height: 100%;
    background-color: #f5f7fa;
    overflow: auto;
    /* 仅让画布区域局部滚动 */
    position: relative;
    padding: 12px;
    /* 压缩画布外部的呼吸留白，最大化释放给 Canvas */
    box-sizing: border-box;

    .viewer-root {
      width: 100%;
      height: 100%;
      min-height: 500px;
    }
  }
}
</style>