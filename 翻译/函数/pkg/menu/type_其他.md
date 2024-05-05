# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[TextType Type = "Text"]
qm=常量_菜单项类型_文本

[SeparatorType Type = "Separator"]
qm=常量_菜单项类型_分隔符

[SubmenuType Type = "Submenu"]
qm=常量_菜单项类型_子菜单

[CheckboxType Type = "Checkbox"]
qm=常量_菜单项类型_复选框

[RadioType Type = "Radio"]
qm=常量_菜单项类型_单选框
