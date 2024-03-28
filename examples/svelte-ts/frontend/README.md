# # Svelte + TS + Vite

这个模板旨在帮助您在Vite中使用Svelte和TypeScript进行开发，从而快速上手。
## # 推荐的IDE设置

[VS Code](https://code.visualstudio.com/)

+ [Svelte](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode)
## # 需要一个官方的 Svelte 框架吗？

请查看 [SvelteKit](https://github.com/sveltejs/kit#readme)。它同样由 Vite 提供支持，采用无服务器优先的方式，可部署到任何地方，并能适应各种平台。它开箱即用支持 TypeScript、SCSS 和 Less，并且可以轻松添加对 mdsvex、GraphQL、PostCSS、Tailwind CSS 等的支持。
## # 技术考量

**为何选择此模板而非SvelteKit？**

- 该模板自带路由解决方案，但可能并非所有用户都偏爱。
- 它首先是一个框架，只是恰巧在底层使用了Vite，而不是一个Vite应用。例如，在SvelteKit环境中，`vite dev` 和 `vite build` 不会生效。

这个模板尽可能精简，仅包含起步所需的Vite + TypeScript + Svelte配置，同时考虑到了HMR（热模块替换）和智能感知的开发者体验。它展示了与其他`create-vite`模板相当的功能，并且是初学者涉足Vite + Svelte项目的良好起点。

若后续需要SvelteKit提供的增强功能与扩展性，本模板已按照类似SvelteKit的结构进行组织，便于迁移。

**为何选用`global.d.ts`而非在`jsconfig.json`或`tsconfig.json`中设置`compilerOptions.types`？**

设置`compilerOptions.types`会排除未在配置中明确列出的所有类型。而使用三斜线引用则保留TypeScript默认设置，即接受整个工作区中的类型信息，同时添加`svelte`和`vite/client`的类型信息。

**为何包含`.vscode/extensions.json`文件？**

其他模板通常通过README间接推荐插件，但此文件可以让VS Code在用户打开项目时提示安装推荐的插件。

**为何在TS模板中启用`allowJs`选项？**

尽管将`allowJs`设置为`false`确实可以阻止项目中使用`.js`文件，但它无法阻止在`.svelte`文件中使用JavaScript语法。此外，这样做还会强制设置`checkJs: false`，导致最糟糕的情况：无法确保整个代码库都是TypeScript，同时对现有的JavaScript进行更差的类型检查。另外，存在一些合理的应用场景，其中混合代码库可能是相关的。

**为何HMR不保留我的组件局部状态？**

HMR状态保持有一些需要注意的问题！由于其行为常常出乎意料，因此在`svelte-hmr`和`@sveltejs/vite-plugin-svelte`中，默认已禁用此功能。您可以在此处[详细了解](https://github.com/rixo/svelte-hmr#svelte-hmr)。

如果您有需要在组件内部保持的重要状态，请考虑创建一个外部存储，该存储不会被HMR替换。

```ts
// store.ts
// 一个极其简单的外部存储
import { writable } from 'svelte/store'
export default writable(0)
```
