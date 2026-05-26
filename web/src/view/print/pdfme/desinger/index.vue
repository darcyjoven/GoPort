<template>
    <div class="my-gva-page-container">
        <div class="embedded-editor-box">
            <PdfmeUniversalEditor v-model="templateData" @action-completed="handleEditorAction" />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import PdfmeUniversalEditor from '../designer.vue'

const templateData = ref(null)

const handleEditorAction = ({ actionStatus, template }) => {
    if (actionStatus === 'submit') {
        console.log('保存模板成功：', template)
    } else {
        console.log('取消了操作')
    }
}
</script>

<style scoped>
/* 父页面样式确保 */
.my-gva-page-container {
    width: 100%;
}

.embedded-editor-box {
    position: relative;
    /* 核心：声明相对定位，限制子组件的查找范围 */
    width: 100%;
    height: calc(100vh - 180px);
    /* 或者是固定高度如 750px，给组件一个撑开的依据 */
    overflow: hidden;
    /* 核心防线：在父级页面再剪裁一次 */
    border: 1px solid #e4e7ed;
    margin-top: 15px;
}
</style>