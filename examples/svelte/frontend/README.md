# # Svelte + Vite

这个模板旨在帮助您开始在Vite中使用Svelte进行开发。
## # 推荐的IDE设置

[VS Code](https://code.visualstudio.com/)

+ [Svelte](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode)
## # 需要一个官方的 Svelte 框架吗？

请查看 [SvelteKit](https://github.com/sveltejs/kit#readme)。它同样由 Vite 提供支持，采用无服务器优先的方式，可部署到任何地方，并能适应各种平台。它开箱即用支持 TypeScript、SCSS 和 Less，并且可以轻松添加对 mdsvex、GraphQL、PostCSS、Tailwind CSS 等的支持。
## # 技术考量

**为何选择这个而非SvelteKit？**

- 它自带路由解决方案，但可能并非所有用户都偏好这种方式。
- 首先，它是一个框架，只是恰好在底层使用了Vite，而不是一个Vite应用。例如，在SvelteKit环境中，`vite dev` 和 `vite build` 不会有效。

此模板包含了尽可能少的起步内容，以便开始使用Vite + Svelte，并考虑到了开发者体验，包括HMR（热模块替换）和智能感知功能。该模板展示了与其他`create-vite`模板相当的能力，是初学者尝试Vite + Svelte项目的良好起点。

若日后需要SvelteKit提供的扩展能力和可扩展性，该模板结构与SvelteKit相似，便于迁移。

**为何使用`global.d.ts`而非在`jsconfig.json`或`tsconfig.json`中设置`compilerOptions.types`？**

设置`compilerOptions.types`会导致未在配置中明确列出的所有其他类型被排除。使用三斜线引用可以保留TypeScript的默认设置，即接受整个工作区中的类型信息，同时添加`svelte`和`vite/client`的类型信息。

**为何包含`.vscode/extensions.json`？**

其他模板通常通过README间接推荐扩展，而这个文件可以让VS Code在用户打开项目时提示安装推荐的扩展。

**为何在JS模板中启用`checkJs`？**

在运行时改变变量类型的大多数情况很可能是意外而非刻意为之。启用`checkJs`提供了高级的类型检查功能。如果你想利用JavaScript的动态类型特性，只需简单地更改配置即可。

**为何HMR无法保留我的组件本地状态？**

HMR状态保持存在许多需要注意的地方！由于其经常表现出令人惊讶的行为，因此在`svelte-hmr`和`@sveltejs/vite-plugin-svelte`中，默认已禁用此功能。你可以在[这里](https://github.com/rixo/svelte-hmr#svelte-hmr)阅读详细信息。

如果你有希望在组件内保留的重要状态，请考虑创建一个外部存储，这样就不会被HMR替换。

```js
// store.js
// 一个极其简单的外部存储
import { writable } from 'svelte/store'
export default writable(0)
```
