
<原文开始>
// TypeOptions overrides options set by `ts_*` tags.
<原文结束>

# <翻译开始>
// TypeOptions 用于覆盖通过 `ts_*` 标签设置的选项。
# <翻译结束>


<原文开始>
// StructType stores settings for transforming one Golang struct.
<原文结束>

# <翻译开始>
// StructType 用于存储转换一个 Golang 结构体所需的设置。
# <翻译结束>


<原文开始>
// throwaway, used when converting
<原文结束>

# <翻译开始>
// 临时变量，用于转换时使用
# <翻译结束>


<原文开始>
// fmt.Println(v.Interface())
<原文结束>

# <翻译开始>
// fmt.Println(v.Interface()) 翻译为：
// 打印v的Interface()方法返回的值
# <翻译结束>


<原文开始>
// Check we have a json tag
<原文结束>

# <翻译开始>
// 检查我们是否有json标签
# <翻译结束>


<原文开始>
// ManageType can define custom options for fields of a specified type.
//
// This can be used instead of setting ts_type and ts_transform for all fields of a certain type.
<原文结束>

# <翻译开始>
// ManageType 可以为指定类型的字段定义自定义选项。
//
// 可以使用此方法替代为某种特定类型的所有字段设置 ts_type 和 ts_transform。
# <翻译结束>


<原文开始>
// Key should always be a JS primitive. JS will read it as a string either way.
<原文结束>

# <翻译开始>
// Key 应始终为 JS 原始类型。无论哪种方式，JS 都会将其读取为字符串。
# <翻译结束>


<原文开始>
// AddEnumValues is deprecated, use `AddEnum()`
<原文结束>

# <翻译开始>
// AddEnumValues 已废弃，使用 `AddEnum()`
# <翻译结束>


<原文开始>
// Put the custom imports, i.e.: `import Decimal from 'decimal.js'`
<原文结束>

# <翻译开始>
// 在此处放置自定义导入，例如：`import Decimal from 'decimal.js'`
# <翻译结束>


<原文开始>
// No neet to backup, just return:
<原文结束>

# <翻译开始>
// 无需备份，直接返回：
# <翻译结束>


<原文开始>
// By default use options defined by tags:
<原文结束>

# <翻译开始>
// 默认情况下，使用由标签定义的选项：
# <翻译结束>


<原文开始>
// But there is maybe an struct-specific override:
<原文结束>

# <翻译开始>
// 但是可能有一个针对结构体的特定覆盖：
# <翻译结束>


<原文开始>
			// Anonymous structures is ignored
			// It is possible to generate them but hard to generate correct name
<原文结束>

# <翻译开始>
			// 匿名结构体将被忽略
			// 虽然可以生成它们，但很难生成正确的名称
# <翻译结束>


<原文开始>
// Also convert map key types if needed
<原文结束>

# <翻译开始>
// 如果需要，同时转换映射键的类型
# <翻译结束>


<原文开始>
// Also convert map value types if needed
<原文结束>

# <翻译开始>
// 如果需要，同时转换映射值的类型
# <翻译结束>


<原文开始>
// Slice of simple fields:
<原文结束>

# <翻译开始>
// 简单字段的切片：
# <翻译结束>


<原文开始>
// check if type is in known enum. If so, then replace TStype with enum name to avoid missing types
<原文结束>

# <翻译开始>
// 检查类型是否在已知的枚举中。如果是，则用枚举名称替换 TStype，以避免丢失类型
# <翻译结束>

