
<原文开始>
// NewBindings returns a new Bindings object
<原文结束>

# <翻译开始>
// NewBindings 返回一个新的 Bindings 对象
# <翻译结束>


<原文开始>
// Yuk yuk yuk! Is there a better way?
<原文结束>

# <翻译开始>
// 呼呼呼！难道没有更好的方法吗？
# <翻译结束>


<原文开始>
// Add the structs to bind
<原文结束>

# <翻译开始>
// 将结构体添加到绑定
# <翻译结束>


<原文开始>
// Add the given struct methods to the Bindings
<原文结束>

# <翻译开始>
// 将给定的结构体方法添加到 Bindings 中
# <翻译结束>


<原文开始>
// Add it as a regular method
<原文结束>

# <翻译开始>
// 将其作为常规方法添加
# <翻译结束>


<原文开始>
// if we have enums for this package, add them as well
<原文结束>

# <翻译开始>
// 如果我们为此包定义了枚举类型，也要一并添加它们
# <翻译结束>


<原文开始>
// Add outstanding enums to the models that were not in packages with structs
<原文结束>

# <翻译开始>
// 将未在包含结构体的包中的枚举添加到模型中
# <翻译结束>


<原文开始>
// Sort the package names first to make the output deterministic
<原文结束>

# <翻译开始>
// 首先对包名进行排序，以确保输出结果的确定性
# <翻译结束>


<原文开始>
// Don't write if we don't have anything
<原文结束>

# <翻译开始>
// 如果没有任何内容，就不要写
# <翻译结束>


<原文开始>
// enums should be represented as array of all possible values
<原文结束>

# <翻译开始>
// 枚举类型应表示为所有可能值的数组
# <翻译结束>


<原文开始>
// simple enum represented by struct with Value/TSName fields
<原文结束>

# <翻译开始>
// 通过具有Value和TSName字段的结构体表示的简单枚举
# <翻译结束>


<原文开始>
// otherwise expecting implementation with TSName() https://github.com/tkrajina/typescriptify-golang-structs#enums-with-tsname
<原文结束>

# <翻译开始>
// 否则期望通过TSName()方法实现，参考：https://github.com/tkrajina/typescriptify-golang-structs#enums-with-tsname
// （该注释含义：在其他情况下，希望按照 TypeScriptify-golang-structs 项目中关于“带有 TSName() 的枚举”部分的说明，采用 TSName() 方法进行实现。）
# <翻译结束>


<原文开始>
// Iterate this struct and add any struct field references
<原文结束>

# <翻译开始>
// 遍历此结构体，并添加任何结构体字段引用
# <翻译结束>

