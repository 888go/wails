package binding_test

type StructWithAnonymousSubMultiLevelStruct struct {
	Name string `json:"name"`
	Meta struct {
		Age  int `json:"age"`
		More struct {
			Info       string `json:"info"`
			MoreInMore struct {
				Demo string `json:"demo"`
			} `json:"more_in_more"`
		} `json:"more"`
	} `json:"meta"`
}

func (s StructWithAnonymousSubMultiLevelStruct) Get() StructWithAnonymousSubMultiLevelStruct {
	return s
}

var AnonymousSubStructMultiLevelTest = BindingTest{
	name: "StructWithAnonymousSubMultiLevelStruct",
	structs: []interface{}{
		&StructWithAnonymousSubMultiLevelStruct{},
	},
	exemptions:  nil,
	shouldError: false,
	want: `
export namespace binding_test {
	export class StructWithAnonymousSubMultiLevelStruct {
		name: string;
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
		meta: any;
	
		static createFrom(source: any = {}) {
			return new StructWithAnonymousSubMultiLevelStruct(source);
		}
	
		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.name = source["name"];
			this.meta = this.convertValues(source["meta"], Object);
		}
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
			if (!a) {
				return a;
			}
			if (a.slice) {
				return (a as any[]).map(elem => this.convertValues(elem, classs));
			} else if ("object" === typeof a) {
				if (asMap) {
					for (const key of Object.keys(a)) {
						a[key] = new classs(a[key]);
					}
					return a;
				}
				return new classs(a);
			}
			return a;
		}
	}

}
`,
}
