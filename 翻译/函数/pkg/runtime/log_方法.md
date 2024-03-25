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

[func LogPrint(ctx context.Context, message string) {]
ff=日志
message=消息
ctx=上下文

[func LogTrace(ctx context.Context, message string) {]
ff=日志追踪
message=消息
ctx=上下文

[func LogDebug(ctx context.Context, message string) {]
ff=日志调试
message=消息
ctx=上下文

[func LogInfo(ctx context.Context, message string) {]
ff=日志信息
message=消息
ctx=上下文

[func LogWarning(ctx context.Context, message string) {]
ff=日志警告
message=消息
ctx=上下文

[func LogError(ctx context.Context, message string) {]
ff=日志错误
message=消息
ctx=上下文

[func LogFatal(ctx context.Context, message string) {]
ff=日志致命
message=消息
ctx=上下文

[func LogPrintf(ctx context.Context, format string, args ...interface{}) {]
ff=日志消息F
args=参数
format=格式化
ctx=上下文

[func LogTracef(ctx context.Context, format string, args ...interface{}) {]
ff=日志追踪消息F
args=参数
format=格式化
ctx=上下文

[func LogDebugf(ctx context.Context, format string, args ...interface{}) {]
ff=日志调试消息F
args=参数
format=格式化
ctx=上下文

[func LogInfof(ctx context.Context, format string, args ...interface{}) {]
ff=日志信息
args=参数
format=格式化
ctx=上下文

[func LogWarningf(ctx context.Context, format string, args ...interface{}) {]
ff=日志警告F
args=参数
format=格式化
ctx=上下文

[func LogErrorf(ctx context.Context, format string, args ...interface{}) {]
ff=日志错误F
args=参数
format=格式化
ctx=上下文

[func LogFatalf(ctx context.Context, format string, args ...interface{}) {]
ff=日志致命F
args=参数
format=格式化
ctx=上下文

[func LogSetLogLevel(ctx context.Context, level logger.LogLevel) {]
ff=设置日志级别
level=常量_日志级别
ctx=上下文
