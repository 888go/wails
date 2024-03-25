# # Go语言通用文件对话框绑定

[项目主页](https://github.com/harry1453/go-common-file-dialog)

本库包含对Windows Vista及更高版本[通用文件对话框](https://docs.microsoft.com/en-us/windows/win32/shell/common-file-dialog)的绑定，这是用于打开或保存文件或文件夹的标准系统对话框。

通用文件对话框需要通过[COM接口](https://en.wikipedia.org/wiki/Component_Object_Model)访问，通常使用C++或绑定（如在C#中）进行访问。

此库为Golang提供了绑定。**它不需要CGO**，并为非Windows平台包含了空的存根（因此在除Windows之外的平台上编译和运行是安全的，但在运行时会返回错误）。

如果你希望快速在Golang应用程序中添加文件选择器，这个库将非常有用。`cfdutil`包包含一些实用函数，只需一个调用即可打开并配置对话框，然后从中获取结果。此类示例可在[`_examples/usingutil`](_examples/usingutil)中找到。或者，如果你希望对对话框的操作有更精细的控制，可以使用基础包。此类示例可在[`_examples/notusingutil`](_examples/notusingutil)中找到。

本库采用MIT许可协议发布。

当前支持的功能：

* 打开单个文件对话框
* 打开多个文件对话框
* 打开文件夹对话框
* 保存文件对话框
* 对话框“角色”功能，允许Windows为不同类型对话框记住不同的“最后位置”
* 设置对话框标题、默认文件夹和初始文件夹
* 设置对话框文件过滤器
