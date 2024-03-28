## # 关于

这是一个极为简单的Wails模板，包含了基本的网页组件（HTML、CSS、JS），并且有意地没有引入任何前端框架、依赖项或Node包管理工具。因此，最终你会得到一个极其轻量级的源文件夹（1 - 5MB）。

也就是说，对于类似“Hello World”应用这样不需要存储约200-300MB源文件的情况，这是一个很好的模板选择。
## # 指令

1. 完成[Wails](https://wails.io/docs/gettingstarted/installation)的所有安装和设置。
2. 在您选择的目录下打开命令提示符。
3. 输入 ``> wails init -n [您的应用名称] -t https://github.com/KiddoV/wails-pure-js-template``
4. 输入 ``> cd ./[您的应用名称]``
5. 输入 ``> wails dev``
6. 继续开发...
## # 实时开发

要以实时开发模式运行，请在项目目录中执行 `wails dev`。前端开发服务器将在 http://localhost:34115 上运行。在浏览器中连接到这个地址，并连接到您的应用程序。
## # 构建

要构建一个可再分发的、生产模式的包，请使用 `wails build` 命令。

或者使用 [UPX](https://upx.github.io/) 进行构建以获得更好的打包体积：``wails build -upx -upxflags="--best --lzma"``

若要使用 ``UPX``，您需要先下载并将其路径至少添加到“系统环境变量”中：

*Windows*
![](https://user-images.githubusercontent.com/28552977/191490618-b84d307e-f783-4c68-bd90-3f484de25478.PNG)
## # 添加依赖项

您不必依赖 ``npm`` 来添加依赖项。

如果您的应用程序需要互联网访问，可以通过 ``CDN`` 链接添加依赖项。

如果您的应用程序在离线环境下使用，只需下载依赖项并将其保存在 ``src/libs`` 文件夹中，然后在 index.html 文件中导入它们。

例如：

```html
<script src="../libs/jquery/jquery-3.7.1.js"></script>
```
