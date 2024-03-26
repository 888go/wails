package binding

import "reflect"

// Parameter 定义了一个 Go 方法的参数
type Parameter struct {
	Name        string `json:"name,omitempty"`
	TypeName    string `json:"type"`
	reflectType reflect.Type
}

func newParameter(Name string, Type reflect.Type) *Parameter {
	return &Parameter{
		Name:        Name,
		TypeName:    Type.String(),
		reflectType: Type,
	}
}

// IsType返回true，如果给定的
func (p *Parameter) IsType(typename string) bool {
	return p.TypeName == typename
}

// IsError 函数返回 true，如果参数类型是 error 类型
func (p *Parameter) IsError() bool {
	return p.IsType("error")
}
