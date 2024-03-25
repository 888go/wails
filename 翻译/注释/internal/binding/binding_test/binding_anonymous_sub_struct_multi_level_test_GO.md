
<原文开始>
// Go type: struct { Age int "json:\"age\""; More struct { Info string "json:\"info\""; MoreInMore struct { Demo string "json:\"demo\"" } "json:\"more_in_more\"" } "json:\"more\"" }
<原文结束>

# <翻译开始>
// Go 类型定义：一个结构体，其字段包含嵌套的结构体，并通过标签指定 JSON 序列化时的键名
// 结构体注释翻译：
// struct {
//     Age int `json:"age"` // 整型变量 Age，在 JSON 中对应的键名为 "age"
//     More struct {      // 嵌套结构体 More
//         Info string `json:"info"` // 字符串变量 Info，在 JSON 中对应的键名为 "info"
//         MoreInMore struct {       // 更深层次的嵌套结构体 MoreInMore
//             Demo string `json:"demo"` // 字符串变量 Demo，在 JSON 中对应的键名为 "demo"
//         } `json:"more_in_more"` // 在 JSON 中对应的键名为 "more_in_more"
//     } `json:"more"` // 在 JSON 中对应的键名为 "more"
// }
# <翻译结束>

