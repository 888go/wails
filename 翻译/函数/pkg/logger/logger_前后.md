# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[TRACE LogLevel = 1]
qm=常量_日志级别_追踪
cz=TRACE LogLevel #等号# 1

[DEBUG LogLevel = 2]
qm=常量_日志级别_调试
cz=DEBUG LogLevel #等号# 2

[INFO LogLevel = 3]
qm=常量_日志级别_信息
cz=INFO LogLevel #等号# 3

[WARNING LogLevel = 4]
qm=常量_日志级别_警告
cz=WARNING LogLevel #等号# 4

[ERROR LogLevel = 5]
qm=常量_日志级别_错误
cz=ERROR LogLevel #等号# 5

[Print(message string)]
qm=日志
cz=Print(message string)

[Trace(message string)]
qm=日志追踪
cz=Trace(message string)

[Debug(message string)]
qm=日志调试
cz=Debug(message string)

[Info(message string)]
qm=日志信息
cz=Info(message string)

[Warning(message string)]
qm=日志警告
cz=Warning(message string)

[Error(message string)]
qm=日志错误
cz=Error(message string)

[Fatal(message string)]
qm=日志致命
cz=Fatal(message string)
