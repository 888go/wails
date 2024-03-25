
<原文开始>
// SemanticVersion is a struct containing a semantic version
<原文结束>

# <翻译开始>
// SemanticVersion 是一个结构体，包含语义化版本信息
# <翻译结束>


<原文开始>
// NewSemanticVersion creates a new SemanticVersion object with the given version string
<原文结束>

# <翻译开始>
// NewSemanticVersion 根据给定的版本字符串创建一个新的 SemanticVersion 对象
# <翻译结束>


<原文开始>
// IsRelease returns true if it's a release version
<原文结束>

# <翻译开始>
// IsRelease 返回 true，如果它是一个发布版本
# <翻译结束>


<原文开始>
// IsPreRelease returns true if it's a prerelease version
<原文结束>

# <翻译开始>
// IsPreRelease 判断是否为预发布版本，如果是则返回 true
# <翻译结束>


<原文开始>
// IsGreaterThan returns true if this version is greater than the given version
<原文结束>

# <翻译开始>
// IsGreaterThan 返回一个布尔值，表示如果当前版本大于给定版本，则返回true
# <翻译结束>


<原文开始>
// Check if the desired one is greater than the requested on
<原文结束>

# <翻译开始>
// 检查期望值是否大于请求值
# <翻译结束>


<原文开始>
// IsGreaterThanOrEqual returns true if this version is greater than or equal the given version
<原文结束>

# <翻译开始>
// IsGreaterThanOrEqual 返回一个布尔值，若当前版本大于或等于给定版本，则返回true
# <翻译结束>


<原文开始>
// MainVersion returns the main version of any version+prerelease+metadata
// EG: MainVersion("1.2.3-pre") => "1.2.3"
<原文结束>

# <翻译开始>
// MainVersion 函数返回任何包含主版本+预发布版本+元数据的版本号中的主版本部分
// 例如：MainVersion("1.2.3-pre") => "1.2.3"
# <翻译结束>


<原文开始>
// SemverCollection is a collection of SemanticVersion objects
<原文结束>

# <翻译开始>
// SemverCollection 是 SemanticVersion 对象的集合
# <翻译结束>


<原文开始>
// Len returns the length of a collection. The number of Version instances
// on the slice.
<原文结束>

# <翻译开始>
// Len 返回一个集合的长度。即切片中 Version 实例的数量。
# <翻译结束>


<原文开始>
// Less is needed for the sort interface to compare two Version objects on the
// slice. If checks if one is less than the other.
<原文结束>

# <翻译开始>
// Less 是为了满足 sort 接口的需求，以便在切片上比较两个 Version 对象。它用于检查一个版本是否小于另一个版本。
# <翻译结束>


<原文开始>
// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
<原文结束>

# <翻译开始>
// Swap 是为了满足 sort 接口的需求，用于在切片中交换两个不同位置的 Version 对象。
# <翻译结束>

