// src/utils/fontLoader.ts

// 1. Vite 自动扫描 src/fonts/ 目录下的所有 ttf 文件 (这里采用 eager: false 的动态按需引入，且作为原始 ArrayBuffer 载入)
// 注意：?arraybuffer 后缀需要保证你的 Vite 支持或配置了相应 loader，更通用的做法是引入 url 后再 fetch 
const fontModules = import.meta.glob('../fonts/*.ttf', { query: '?url', eager: true });

/**
 * 自动识别 src/fonts 目录并动态加载全部字体
 * @param fallbackFontName 指定哪个字体名作为全局中文字体兜底
 */
export async function autoLoadAllFonts(fallbackFontName = '黑体') {
    try {
        const fontConfig = {};

        // 收集所有的动态 fetch 任务
        const loadTasks = Object.entries(fontModules).map(async ([path, module]) => {
            // 从路径中自动提取字体文件名作为 key。例如: "../fonts/SimHei.ttf" -> "SimHei"
            const fontName = path.split('/').pop()?.replace('.ttf', '') || 'UnknownFont';

            // module 此时是 Vite 处理后的静态资源 URL 路径
            const response = await fetch(module.default || module);
            const buffer = await response.arrayBuffer();

            return {
                fontName,
                data: buffer
            };
        });

        // 并行执行所有的加载任务
        const loadedFonts = await Promise.all(loadTasks);

        // 组装格式
        loadedFonts.forEach(({ fontName, data }) => {
            fontConfig[fontName] = {
                data,
                ...(fontName === fallbackFontName ? { fallback: true } : {})
            };
        });
        return fontConfig;
    } catch (error) {
        console.error('Vite 自动目录字体加载失败', error);
        return undefined;
    }
}