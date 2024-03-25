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

[func NewFileLogger(filename string) Logger {]
ff=创建文件日志
filename=文件路径

[func (l *FileLogger) Print(message string) {]
ff=日志
message=消息

[func (l *FileLogger) Println(message string) {]
ff=日志并换行
message=消息

[func (l *FileLogger) Trace(message string) {]
ff=日志追踪
message=消息

[func (l *FileLogger) Debug(message string) {]
ff=日志调试
message=消息

[func (l *FileLogger) Info(message string) {]
ff=日志信息
message=消息

[func (l *FileLogger) Warning(message string) {]
ff=日志警告
message=消息

[func (l *FileLogger) Error(message string) {]
ff=日志错误
message=消息

[func (l *FileLogger) Fatal(message string) {]
ff=日志致命
message=消息
