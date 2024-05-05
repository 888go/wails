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

[TRACE LogLevel = 1]
qm=常量_日志级别_追踪

[DEBUG LogLevel = 2]
qm=常量_日志级别_调试

[INFO LogLevel = 3]
qm=常量_日志级别_信息

[WARNING LogLevel = 4]
qm=常量_日志级别_警告

[ERROR LogLevel = 5]
qm=常量_日志级别_错误

[Print(message string)]
qm=日志

[Trace(message string)]
qm=日志追踪

[Debug(message string)]
qm=日志调试

[Info(message string)]
qm=日志信息

[Warning(message string)]
qm=日志警告

[Error(message string)]
qm=日志错误

[Fatal(message string)]
qm=日志致命
