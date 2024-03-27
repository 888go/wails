package binding

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// BoundMethod 定义了与绑定到 Wails 应用的所有 Go 方法相关的所有数据
type BoundMethod struct {
	Name     string        `json:"name"`
	Inputs   []*Parameter  `json:"inputs,omitempty"`
	Outputs  []*Parameter  `json:"outputs,omitempty"`
	Comments string        `json:"comments,omitempty"`
	Method   reflect.Value `json:"-"`
}

// InputCount 返回此绑定方法的输入参数个数

// ff:
func (b *BoundMethod) InputCount() int {
	return len(b.Inputs)
}

// OutputCount 返回该绑定方法的输出数量

// ff:
func (b *BoundMethod) OutputCount() int {
	return len(b.Outputs)
}

// ParseArgs 方法将输入的 JSON 转换为该方法期望的类型

// ff:
// args:
func (b *BoundMethod) ParseArgs(args []json.RawMessage) ([]interface{}, error) {
	result := make([]interface{}, b.InputCount())
	if len(args) != b.InputCount() {
		return nil, fmt.Errorf("received %d arguments to method '%s', expected %d", len(args), b.Name, b.InputCount())
	}
	for index, arg := range args {
		typ := b.Inputs[index].reflectType
		inputValue := reflect.New(typ).Interface()
		err := json.Unmarshal(arg, inputValue)
		if err != nil {
			return nil, err
		}
		if inputValue == nil {
			result[index] = reflect.Zero(typ).Interface()
		} else {
			result[index] = reflect.ValueOf(inputValue).Elem().Interface()
		}
	}
	return result, nil
}

// Call 将尝试使用给定的参数调用此绑定方法

// ff:
// args:
func (b *BoundMethod) Call(args []interface{}) (interface{}, error) {
	// Check inputs
	expectedInputLength := len(b.Inputs)
	actualInputLength := len(args)
	if expectedInputLength != actualInputLength {
		return nil, fmt.Errorf("%s takes %d inputs. Received %d", b.Name, expectedInputLength, actualInputLength)
	}

	/** Convert inputs to reflect values **/

	// 创建一个切片，用于存储对方法调用的输入参数
	callArgs := make([]reflect.Value, expectedInputLength)

	// 遍历给定的参数
	for index, arg := range args {
		// 保存转换后的参数
		callArgs[index] = reflect.ValueOf(arg)
	}

	// Do the call
	callResults := b.Method.Call(callArgs)

	//** Check results **//
	var returnValue interface{}
	var err error

	switch b.OutputCount() {
	case 1:
// 遍历结果并判断结果是否为错误
		for _, result := range callResults {
			interfac := result.Interface()
			temp, ok := interfac.(error)
			if ok {
				err = temp
			} else {
				returnValue = interfac
			}
		}
	case 2:
		returnValue = callResults[0].Interface()
		if temp, ok := callResults[1].Interface().(error); ok {
			err = temp
		}
	}

	return returnValue, err
}
