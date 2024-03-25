## # 对话

注意：目前，这是一项仅限Mac的功能。

将任何PNG文件放入此目录中，以便在消息对话框中使用它们。文件应采用以下格式命名：`name[-(light|dark)][2x].png`

示例：

* `mypic.png` - 标准定义图标，ID 为 `mypic`
* `mypic-light.png` - 标准定义图标，ID 为 `mypic`，当系统主题为浅色时使用
* `mypic-dark.png` - 标准定义图标，ID 为 `mypic`，当系统主题为深色时使用
* `mypic2x.png` - 高清定义图标，ID 为 `mypic`
* `mypic-light2x.png` - 高清定义图标，ID 为 `mypic`，当系统主题为浅色时使用
* `mypic-dark2x.png` - 高清定义图标，ID 为 `mypic`，当系统主题为深色时使用

（注：由于上下文不完整，这里未提供Markdown格式的标题部分翻译）
## # 优先级顺序

图标的选择遵循以下优先级顺序：

对于高清显示器：
1. name-(theme)2x.png
2. name2x.png
3. name-(theme).png
4. name.png

对于标准清晰度显示器：
1. name-(theme).png
2. name.png
