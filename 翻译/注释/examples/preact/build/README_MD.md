
<原文开始>
Build Directory

The build directory is used to house all the build files and assets for your application. 

The structure is:

* bin - Output directory
* darwin - macOS specific files
* windows - Windows specific files


<原文结束>

# <翻译开始>
# 构建目录

构建目录用于存储应用程序的所有构建文件和资源。

其结构如下：

* bin - 输出目录
* darwin - macOS 特定文件
* windows - Windows 特定文件

# <翻译结束>


<原文开始>
Mac

The `darwin` directory holds files specific to Mac builds.
These may be customised and used as part of the build. To return these files to the default state, simply delete them
and
build with `wails build`.

The directory contains the following files:

- `Info.plist` - the main plist file used for Mac builds. It is used when building using `wails build`.
- `Info.dev.plist` - same as the main plist file but used when building using `wails dev`.


<原文结束>

# <翻译开始>
# Mac

`darwin`目录包含针对Mac构建的特定文件。您可以自定义这些文件并在构建过程中使用它们。若要将这些文件恢复到默认状态，只需删除它们，然后使用`wails build`进行构建。

该目录包含以下文件：

- `Info.plist` - 用于Mac构建的主要plist文件，在使用`wails build`构建时会用到。
- `Info.dev.plist` - 与主plist文件相同，但在使用`wails dev`构建时使用。

# <翻译结束>


<原文开始>
Windows

The `windows` directory contains the manifest and rc files used when building with `wails build`.
These may be customised for your application. To return these files to the default state, simply delete them and
build with `wails build`.

- `icon.ico` - The icon used for the application. This is used when building using `wails build`. If you wish to
  use a different icon, simply replace this file with your own. If it is missing, a new `icon.ico` file
  will be created using the `appicon.png` file in the build directory.
- `installer/*` - The files used to create the Windows installer. These are used when building using `wails build`.
- `info.json` - Application details used for Windows builds. The data here will be used by the Windows installer,
  as well as the application itself (right click the exe -> properties -> details)
- `wails.exe.manifest` - The main application manifest file.
<原文结束>

# <翻译开始>
# Windows

`windows`目录包含了使用`wails build`构建项目时所使用的清单和rc文件。你可以根据自己的应用需求对这些文件进行定制。若要将这些文件恢复到默认状态，只需删除它们并使用`wails build`命令重新构建。

- `icon.ico` - 应用程序所使用的图标。在使用`wails build`构建应用程序时会用到这个图标。如果你想使用不同的图标，只需用你自己的图标文件替换此文件即可。如果该文件缺失，系统会自动使用构建目录下的`appicon.png`文件生成一个新的`icon.ico`文件。
- `installer/*` - 用于创建Windows安装包的文件。这些文件在使用`wails build`构建时会被使用。
- `info.json` - 用于Windows构建的应用程序详细信息。其中的数据不仅会被Windows安装程序使用，还会被应用程序自身（右键点击exe文件 -> 属性 -> 详细信息）所引用。
- `wails.exe.manifest` - 主要的应用程序清单文件。

# <翻译结束>

