
<原文开始>
Svelte + TS + Vite

This template should help get you started developing with Svelte and TypeScript in Vite.


<原文结束>

# <翻译开始>
# Svelte + TS + Vite

这个模板旨在帮助您在Vite中使用Svelte和TypeScript进行开发，从而快速上手。

# <翻译结束>


<原文开始>
Recommended IDE Setup

[VS Code](https://code.visualstudio.com/)

+ [Svelte](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode).


<原文结束>

# <翻译开始>
# 推荐的IDE设置

[VS Code](https://code.visualstudio.com/)

+ [Svelte](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode)

# <翻译结束>


<原文开始>
Need an official Svelte framework?

Check out [SvelteKit](https://github.com/sveltejs/kit#readme), which is also powered by Vite. Deploy anywhere with its
serverless-first approach and adapt to various platforms, with out of the box support for TypeScript, SCSS, and Less,
and easily-added support for mdsvex, GraphQL, PostCSS, Tailwind CSS, and more.


<原文结束>

# <翻译开始>
# 需要一个官方的 Svelte 框架吗？

请查看 [SvelteKit](https://github.com/sveltejs/kit#readme)。它同样由 Vite 提供支持，采用无服务器优先的方式，可部署到任何地方，并能适应各种平台。它开箱即用支持 TypeScript、SCSS 和 Less，并且可以轻松添加对 mdsvex、GraphQL、PostCSS、Tailwind CSS 等的支持。

# <翻译结束>


<原文开始>
Technical considerations

**Why use this over SvelteKit?**

- It brings its own routing solution which might not be preferable for some users.
- It is first and foremost a framework that just happens to use Vite under the hood, not a Vite app.
  `vite dev` and `vite build` wouldn't work in a SvelteKit environment, for example.

This template contains as little as possible to get started with Vite + TypeScript + Svelte, while taking into account
the developer experience with regards to HMR and intellisense. It demonstrates capabilities on par with the
other `create-vite` templates and is a good starting point for beginners dipping their toes into a Vite + Svelte
project.

Should you later need the extended capabilities and extensibility provided by SvelteKit, the template has been
structured similarly to SvelteKit so that it is easy to migrate.

**Why `global.d.ts` instead of `compilerOptions.types` inside `jsconfig.json` or `tsconfig.json`?**

Setting `compilerOptions.types` shuts out all other types not explicitly listed in the configuration. Using triple-slash
references keeps the default TypeScript setting of accepting type information from the entire workspace, while also
adding `svelte` and `vite/client` type information.

**Why include `.vscode/extensions.json`?**

Other templates indirectly recommend extensions via the README, but this file allows VS Code to prompt the user to
install the recommended extension upon opening the project.

**Why enable `allowJs` in the TS template?**

While `allowJs: false` would indeed prevent the use of `.js` files in the project, it does not prevent the use of
JavaScript syntax in `.svelte` files. In addition, it would force `checkJs: false`, bringing the worst of both worlds:
not being able to guarantee the entire codebase is TypeScript, and also having worse typechecking for the existing
JavaScript. In addition, there are valid use cases in which a mixed codebase may be relevant.

**Why is HMR not preserving my local component state?**

HMR state preservation comes with a number of gotchas! It has been disabled by default in both `svelte-hmr`
and `@sveltejs/vite-plugin-svelte` due to its often surprising behavior. You can read the
details [here](https://github.com/rixo/svelte-hmr#svelte-hmr).

If you have state that's important to retain within a component, consider creating an external store which would not be
replaced by HMR.

```ts
// store.ts
// An extremely simple external store
import { writable } from 'svelte/store'
export default writable(0)
```

<原文结束>

# <翻译开始>
# 技术考量

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

# <翻译结束>

