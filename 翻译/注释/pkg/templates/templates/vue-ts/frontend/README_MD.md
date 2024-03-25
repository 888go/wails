
<原文开始>
Vue 3 + TypeScript + Vite

This template should help get you started developing with Vue 3 and TypeScript in Vite. The template uses Vue
3 `<script setup>` SFCs, check out
the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.


<原文结束>

# <翻译开始>
# Vue 3 + TypeScript + Vite

这个模板旨在帮助您在Vite中使用Vue 3和TypeScript进行开发。该模板采用了Vue
3 `<script setup>` 单文件组件（SFCs），想要了解更多，请查阅[script setup文档](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup)。

# <翻译结束>


<原文开始>
Recommended IDE Setup

- [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar)


<原文结束>

# <翻译开始>
# 推荐的 IDE 设置

- [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar)

# <翻译结束>


<原文开始>
Type Support For `.vue` Imports in TS

Since TypeScript cannot handle type information for `.vue` imports, they are shimmed to be a generic Vue component type
by default. In most cases this is fine if you don't really care about component prop types outside of templates.
However, if you wish to get actual prop types in `.vue` imports (for example to get props validation when using
manual `h(...)` calls), you can enable Volar's Take Over mode by following these steps:

1. Run `Extensions: Show Built-in Extensions` from VS Code's command palette, look
   for `TypeScript and JavaScript Language Features`, then right click and select `Disable (Workspace)`. By default,
   Take Over mode will enable itself if the default TypeScript extension is disabled.
2. Reload the VS Code window by running `Developer: Reload Window` from the command palette.

You can learn more about Take Over mode [here](https://github.com/johnsoncodehk/volar/discussions/471).

<原文结束>

# <翻译开始>
# Vue 文件导入的类型支持

由于 TypeScript 无法处理 `.vue` 导入的类型信息，默认情况下，它们会被模拟为一个通用的 Vue 组件类型。在大多数情况下，如果你不在意模板之外的组件属性类型，这样做是可行的。然而，如果你想在 `.vue` 导入中获得实际的属性类型（例如，在手动使用 `h(...)` 调用时获取属性验证），可以通过以下步骤启用 Volar 的接管模式：

1. 在 VS Code 命令面板中运行 `Extensions: Show Built-in Extensions`，找到 `TypeScript and JavaScript Language Features`，然后右键点击并选择 `Disable (Workspace)`。默认情况下，如果禁用了默认的 TypeScript 扩展，接管模式将自动启用。
2. 通过命令面板运行 `Developer: Reload Window` 以重新加载 VS Code 窗口。

你可以在这里了解更多关于接管模式的信息：[https://github.com/johnsoncodehk/volar/discussions/471](https://github.com/johnsoncodehk/volar/discussions/471)。

# <翻译结束>

