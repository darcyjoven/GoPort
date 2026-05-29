<template>
    <div class="vtable-container">
        <div class="toolbar">
            <el-button size="small" @click="goBack"> <el-icon>
                    <ArrowLeft />
                </el-icon> 返回 </el-button>
            <el-button type="primary" size="small" icon="Printer" :loading="printLoading" @click="handlePrint">
                调用浏览器打印
            </el-button>
            <slot name="extra-actions"></slot>
            <div class="right-action-group">
                <slot name="right-actions"></slot>
            </div>
        </div>

        <div v-loading="pageLoading" element-loading-text="正在加载字体与渲染预览..." class="table-wrapper">
            <div ref="viewerRootRef" class="viewer-root"></div>
        </div>

        <iframe id="pdf-print-iframe" allow="unload" style="display: none;"></iframe>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, toRaw, watch } from 'vue'
import { ElMessage, ElLoading } from 'element-plus'
import { Viewer } from '@pdfme/ui'
import { generate } from '@pdfme/generator'
import type { Template } from '@pdfme/common'
import { autoLoadAllFonts } from '@/utils/fontLoader'
import {
    text, image, table, barcodes, line, rectangle, ellipse, dateTime, select, radioGroup,
    checkbox, multiVariableText, list, svg, signature, date, time
} from '@pdfme/schemas'

// --- 规范化定义组件的 Props ---
interface Props {
    templateData?: Template         // 允许外部传入样式模板，不传则使用内置默认
    inputsData?: Record<string, any>[] // 允许外部传入业务数据集，不传则为空
    title?: string                  // 允许自定义标题
}

const goBack = () => {
    history.back()
}

const props = withDefaults(defineProps<Props>(), {
    templateData: () => ({
        basePdf: { width: 210, height: 297, padding: [2, 2, 2, 2] },
        schemas: [[]]
    }),
    inputsData: () => [],
    title: '单据预览消费端'
})

// --- pdfme 插件注册表 ---
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

// --- 响应式状态管理 ---
const viewerRootRef = ref<HTMLDivElement | null>(null)
const viewerInstance = ref<Viewer | null>(null)
const pageLoading = ref<boolean>(false)
const printLoading = ref<boolean>(false)
let globalFontOptions: any = null

/**
 * 核心初始化：载入字体并实例化 Viewer
 */
const initViewer = async (): Promise<void> => {
    if (!viewerRootRef.value) return
    pageLoading.value = true

    try {
        // 异步加载全局多字体
        const fontOptions: any = await autoLoadAllFonts()
        globalFontOptions = fontOptions

        // 遵循单向数据流：断开外部 template 的深层响应式追踪，防止画布被 Vue 频繁扫描引发性能问题
        const cleanTemplate = JSON.parse(JSON.stringify(toRaw(props.templateData))) as Template
        const cleanInputs = JSON.parse(JSON.stringify(toRaw(props.inputsData)))

        // 如果实例已经存在，优先执行清理
        if (viewerInstance.value) {
            viewerInstance.value.destroy()
        }

        viewerInstance.value = new Viewer({
            domContainer: viewerRootRef.value,
            template: cleanTemplate,
            inputs: cleanInputs,
            plugins: PDFME_PLUGINS,
            options: {
                font: fontOptions
            }
        })
    } catch (error) {
        console.error('[PdfmeViewer] 初始化视图失败:', error)
        ElMessage.error('渲染预览失败，请检查模板格式或字体文件')
    } finally {
        pageLoading.value = false
    }
}

/**
 * 静默唤起浏览器系统打印面板
 */
const handlePrint = async (): Promise<void> => {
    if (printLoading.value) return
    printLoading.value = true

    const loading = ElLoading.service({
        lock: true,
        text: '正在编译高保真打印二进制流...',
        background: 'rgba(255, 255, 255, 0.7)'
    })

    try {
        const cleanTemplate = JSON.parse(JSON.stringify(toRaw(props.templateData))) as Template
        const cleanInputs = JSON.parse(JSON.stringify(toRaw(props.inputsData)))

        // 编译纯前端 PDF 二进制字节流
        const pdfUint8Array = await generate({
            template: cleanTemplate,
            inputs: cleanInputs,
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
                    URL.revokeObjectURL(pdfUrl) // 延迟释放本地临时 URL 内存
                }, 2000)
            }
        } else {
            throw new Error('未检测到专用的打印容器节点')
        }
    } catch (error) {
        console.error('[PdfmeViewer] 唤起系统打印错误:', error)
        ElMessage.error('打印调度失败，请重新尝试')
    } finally {
        printLoading.value = false
        loading.close()
    }
}

// --- 💡 核心设计：深度监听数据源变化，实现无缝增量刷新 ---
watch(
    () => props.inputsData,
    (newInputs) => {
        if (viewerInstance.value) {
            try {
                const cleanInputs = JSON.parse(JSON.stringify(toRaw(newInputs)))
                // 增量热更新数据，不销毁 DOM，无闪烁且效率极高
                // viewerInstance.value.updateInputs(cleanInputs)
                viewerInstance.value.setInputs(JSON.parse(JSON.stringify(toRaw(newInputs))))
            } catch (e) {
                console.warn('[PdfmeViewer] 热更新数据集失败，尝试降级重新初始化:', e)
                initViewer()
            }
        }
    },
    { deep: true }
)

// 监听模板结构的变化
watch(
    () => props.templateData,
    () => {
        // 结构样式发生改变时，需要重新生成底层 Canvas 物理拓扑
        initViewer()
    },
    { deep: true }
)

// 生命周期
onMounted(() => {
    initViewer()
})

onBeforeUnmount(() => {
    if (viewerInstance.value) {
        try {
            viewerInstance.value.destroy() // 显式释放 Canvas 内存，斩断闭包，杜绝 GVA 页面频繁切换导致的内存泄漏
        } catch (e) {
            console.warn('清理 Canvas 实例警告:', e)
        }
        viewerInstance.value = null
    }
})
</script>

<style scoped>
.vtable-container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.toolbar {
    display: flex;
    gap: 12px;
    padding: 12px 12px 0 12px;
}

.right-action-group {
    margin-left: auto;
}

.table-wrapper {
    flex: 1;
    min-height: 0;
    border: 1px solid #ebf0f5;
    border-radius: 4px;
    background-color: #f5f7fa;
    overflow: auto;
    position: relative;
    padding: 0 12px;
    box-sizing: border-box;

    .viewer-root {
        width: 100%;
        height: 100%;
        min-height: 500px;
    }
}
</style>