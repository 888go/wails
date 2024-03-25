# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[func (m *MenuItem) Remove() {]
ff=删除

[func (m *MenuItem) Parent() *MenuItem {]
ff=取父菜单

[func (m *MenuItem) Append(item *MenuItem) bool {]
ff=加入子菜单
item=菜单项

[func (m *MenuItem) Prepend(item *MenuItem) bool {]
ff=加入子菜单最前
item=菜单项

[func (m *MenuItem) InsertAfter(item *MenuItem) bool {]
ff=插入当前后面
item=菜单项

[func (m *MenuItem) InsertBefore(item *MenuItem) bool {]
ff=插入当前前面
item=菜单项

[func (m *MenuItem) SetLabel(name string) {]
ff=设置显示名称
name=名称

[func (m *MenuItem) IsSeparator() bool {]
ff=是否为分隔符

[func (m *MenuItem) IsCheckbox() bool {]
ff=是否为复选框

[func (m *MenuItem) Disable() *MenuItem {]
ff=设置禁用

[func (m *MenuItem) Enable() *MenuItem {]
ff=取消禁用

[func (m *MenuItem) OnClick(click Callback) *MenuItem {]
ff=绑定单击事件
click=回调函数

[func (m *MenuItem) SetAccelerator(acc *keys.Accelerator) *MenuItem {]
ff=设置快捷键
acc=快捷键

[func (m *MenuItem) SetChecked(value bool) *MenuItem {]
ff=设置选中
value=选中

[func (m *MenuItem) Hide() *MenuItem {]
ff=设置隐藏

[func (m *MenuItem) Show() *MenuItem {]
ff=取消隐藏

[func (m *MenuItem) IsRadio() bool {]
ff=是否为菜单项

[func Label(label string) *MenuItem {]
ff=创建文本菜单项
label=显示名称

[func Text(label string, accelerator *keys.Accelerator, click Callback) *MenuItem {]
ff=创建文本菜单项2
click=单击回调函数
accelerator=快捷键
label=显示名称

[func Separator() *MenuItem {]
ff=创建分隔符菜单项

[func Radio(label string, selected bool, accelerator *keys.Accelerator, click Callback) *MenuItem {]
ff=创建单选框菜单项
click=单击回调函数
accelerator=快捷键
selected=选中
label=显示名称

[func Checkbox(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {]
ff=创建复选框菜单项
click=单击回调函数
accelerator=快捷键
checked=选中
label=显示名称

[func SubMenu(label string, menu *Menu) *MenuItem {]
ff=创建子菜单
menu=子菜单
label=显示名称
