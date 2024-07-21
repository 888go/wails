
<原文开始>
// Parse parses the given tags string and returns
// a cleaned slice of strings. Both comma and space delimeted
// tags are supported but not mixed. If mixed, an error is returned.
<原文结束>

# <翻译开始>
// Parse函数用于解析给定的标签字符串，并返回
// 一个清理过的字符串切片。同时支持逗号和空格作为分隔符，
// 但不支持混合使用，若混合使用则会返回错误。
# <翻译结束>


<原文开始>
		// We couldn't find any separator, so the whole string is used as user tag
		// Otherwise we would end up with a list of every single character of the tags string,
		// e.g.: `t,e,s,t`
<原文结束>

# <翻译开始>
		// 我们未能找到任何分隔符，因此将整个字符串作为用户标签使用
		// 否则最终我们将得到一个包含标签字符串中每个单字符的列表，
		// 例如：`t,e,s,t`
# <翻译结束>


<原文开始>
// Stringify converts the given tags slice to a string compatible
// with the go build -tags flag
<原文结束>

# <翻译开始>
// Stringify 将给定的标签切片转换为与 go build -tags 标志兼容的字符串
# <翻译结束>

