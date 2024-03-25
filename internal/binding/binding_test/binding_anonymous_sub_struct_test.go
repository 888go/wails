package binding_test

type StructWithAnonymousSubStruct struct {
	Name string `json:"name"`
	Meta struct {
		Age int `json:"age"`
	} `json:"meta"`
}

func (s StructWithAnonymousSubStruct) Get() StructWithAnonymousSubStruct {
	return s
}

var AnonymousSubStructTest = BindingTest{
	name: "StructWithAnonymousSubStruct",
	structs: []interface{}{
		&StructWithAnonymousSubStruct{},
	},
	exemptions:  nil,
	shouldError: false,
	want: `
export namespace binding_test {
	export class StructWithAnonymousSubStruct {
		name: string;
		// Go 类型：结构体 { Age int "json:\"age\"" }
// 这是一个Go语言的结构体类型定义，其中包含一个字段：
// Age：整数类型，其标签为"json:\"age\""，表示在进行JSON格式编解码时，该字段对应的JSON键名为"age"。
		meta: any;
	
		static createFrom(source: any = {}) {
			return new StructWithAnonymousSubStruct(source);
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
