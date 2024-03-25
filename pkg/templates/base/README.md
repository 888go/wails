# README

## # 关于

这是官方Wails $NAME 模板。

您可以通过编辑 `wails.json` 配置项目。有关项目设置的更多信息，请访问：
https://wails.io/docs/reference/project-config
## # 实时开发

要以实时开发模式运行，在项目目录下执行 `wails dev`。这将运行一个 Vite 开发服务器，能够非常快速地实现前端更改的热重载。如果你想在浏览器中进行开发并访问 Go 方法，还有一个运行在 http://localhost:34115 的开发服务器。在浏览器中连接到这个地址，你就可以通过开发者工具调用 Go 代码了。
## # 构建

要构建一个可再分发的、生产模式的包，请使用 `wails build`。
