
<原文开始>
// DB is our database of method bindings
<原文结束>

# <翻译开始>
// DB 是我们方法绑定的数据库
# <翻译结束>


<原文开始>
//  map[packagename] -> map[structname] -> map[methodname]*method
<原文结束>

# <翻译开始>
// map[包名] -> map[结构体名] -> map[方法名]*方法
// 该注释描述了一个Go语言中使用的嵌套映射数据结构，其中：
// - 外层映射的键是包名称（packagename）。
// - 中间层映射的键是结构体名称（structname）。
// - 内层映射的键是方法名称（methodname），其值是指向该方法的指针（*method）。
# <翻译结束>


<原文开始>
	// This uses fully qualified method names as a shortcut for store traversal.
	// It used for performance gains at runtime
<原文结束>

# <翻译开始>
// 此处使用完全限定方法名作为遍历 store 的快捷方式。
// 这是为了在运行时获取性能提升。
# <翻译结束>


<原文开始>
// This uses ids to reference bound methods at runtime
<原文结束>

# <翻译开始>
// 这段代码在运行时使用ids来引用已绑定的方法
# <翻译结束>


<原文开始>
// Lock to ensure sync access to the data
<原文结束>

# <翻译开始>
// 加锁以确保对数据的同步访问
# <翻译结束>


<原文开始>
// GetMethodFromStore returns the method for the given package/struct/method names
// nil is returned if any one of those does not exist
<原文结束>

# <翻译开始>
// GetMethodFromStore 根据给定的包名/结构体名/方法名返回对应的方法
// 如果其中任何一个不存在，则返回 nil
# <翻译结束>


<原文开始>
// Lock the db whilst processing and unlock on return
<原文结束>

# <翻译开始>
// 在处理过程中锁定数据库，并在返回时解锁
# <翻译结束>


<原文开始>
// GetMethod returns the method for the given qualified method name
// qualifiedMethodName is "packagename.structname.methodname"
<原文结束>

# <翻译开始>
// GetMethod 返回给定的完整方法名所对应的方法
// 其中，qualifiedMethodName 的格式为 "包名.结构名.方法名"
# <翻译结束>


<原文开始>
// GetObfuscatedMethod returns the method for the given ID
<原文结束>

# <翻译开始>
// GetObfuscatedMethod 根据给定ID返回方法
# <翻译结束>


<原文开始>
// AddMethod adds the given method definition to the db using the given qualified path: packageName.structName.methodName
<原文结束>

# <翻译开始>
// AddMethod 将给定的方法定义通过指定的完全限定路径添加到db中：packageName.structName.methodName
# <翻译结束>


<原文开始>
// Get the map associated with the package name
<原文结束>

# <翻译开始>
// 获取与包名关联的映射
# <翻译结束>


<原文开始>
// Create a new map for this packagename
<原文结束>

# <翻译开始>
// 为这个包名创建一个新的映射
# <翻译结束>


<原文开始>
// Get the map associated with the struct name
<原文结束>

# <翻译开始>
// 获取与结构体名称关联的映射
# <翻译结束>


<原文开始>
// Store the method definition
<原文结束>

# <翻译开始>
// 存储方法定义
# <翻译结束>


<原文开始>
// ToJSON converts the method map to JSON
<原文结束>

# <翻译开始>
// ToJSON 将方法映射转换为 JSON
# <翻译结束>


<原文开始>
// Return zero copy string as this string will be read only
<原文结束>

# <翻译开始>
// 返回零拷贝字符串，因为该字符串将只读
# <翻译结束>


<原文开始>
// UpdateObfuscatedCallMap sets up the secure call mappings
<原文结束>

# <翻译开始>
// UpdateObfuscatedCallMap 设置安全调用映射
# <翻译结束>

