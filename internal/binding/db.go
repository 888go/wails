package binding

import (
	"encoding/json"
	"sync"
	"unsafe"
)

// DB 是我们方法绑定的数据库
type DB struct {
	// map[包名] -> map[结构体名] -> map[方法名]*方法
// 该注释描述了一个Go语言中使用的嵌套映射数据结构，其中：
// - 外层映射的键是包名称（packagename）。
// - 中间层映射的键是结构体名称（structname）。
// - 内层映射的键是方法名称（methodname），其值是指向该方法的指针（*method）。
	store map[string]map[string]map[string]*BoundMethod

// 此处使用完全限定方法名作为遍历 store 的快捷方式。
// 这是为了在运行时获取性能提升。
	methodMap map[string]*BoundMethod

	// 这段代码在运行时使用ids来引用已绑定的方法
	obfuscatedMethodArray []*ObfuscatedMethod

	// 加锁以确保对数据的同步访问
	lock sync.RWMutex
}

type ObfuscatedMethod struct {
	method     *BoundMethod
	methodName string
}

func newDB() *DB {
	return &DB{
		store:                 make(map[string]map[string]map[string]*BoundMethod),
		methodMap:             make(map[string]*BoundMethod),
		obfuscatedMethodArray: []*ObfuscatedMethod{},
	}
}

// GetMethodFromStore 根据给定的包名/结构体名/方法名返回对应的方法
// 如果其中任何一个不存在，则返回 nil

// ff:
// methodName:
// structName:
// packageName:
func (d *DB) GetMethodFromStore(packageName string, structName string, methodName string) *BoundMethod {
	// 在处理过程中锁定数据库，并在返回时解锁
	d.lock.RLock()
	defer d.lock.RUnlock()

	structMap, exists := d.store[packageName]
	if !exists {
		return nil
	}
	methodMap, exists := structMap[structName]
	if !exists {
		return nil
	}
	return methodMap[methodName]
}

// GetMethod 返回给定的完整方法名所对应的方法
// 其中，qualifiedMethodName 的格式为 "包名.结构名.方法名"

// ff:
// qualifiedMethodName:
func (d *DB) GetMethod(qualifiedMethodName string) *BoundMethod {
	// 在处理过程中锁定数据库，并在返回时解锁
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.methodMap[qualifiedMethodName]
}

// GetObfuscatedMethod 根据给定ID返回方法

// ff:
// id:
func (d *DB) GetObfuscatedMethod(id int) *BoundMethod {
	// 在处理过程中锁定数据库，并在返回时解锁
	d.lock.RLock()
	defer d.lock.RUnlock()

	if len(d.obfuscatedMethodArray) <= id {
		return nil
	}

	return d.obfuscatedMethodArray[id].method
}

// AddMethod 将给定的方法定义通过指定的完全限定路径添加到db中：packageName.structName.methodName

// ff:
// methodDefinition:
// methodName:
// structName:
// packageName:
func (d *DB) AddMethod(packageName string, structName string, methodName string, methodDefinition *BoundMethod) {
	// 在处理过程中锁定数据库，并在返回时解锁
	d.lock.Lock()
	defer d.lock.Unlock()

	// 获取与包名关联的映射
	structMap, exists := d.store[packageName]
	if !exists {
		// 为这个包名创建一个新的映射
		d.store[packageName] = make(map[string]map[string]*BoundMethod)
		structMap = d.store[packageName]
	}

	// 获取与结构体名称关联的映射
	methodMap, exists := structMap[structName]
	if !exists {
		// 为这个包名创建一个新的映射
		structMap[structName] = make(map[string]*BoundMethod)
		methodMap = structMap[structName]
	}

	// 存储方法定义
	methodMap[methodName] = methodDefinition

	// Store in the methodMap
	key := packageName + "." + structName + "." + methodName
	d.methodMap[key] = methodDefinition
	d.obfuscatedMethodArray = append(d.obfuscatedMethodArray, &ObfuscatedMethod{method: methodDefinition, methodName: key})
}

// ToJSON 将方法映射转换为 JSON

// ff:
func (d *DB) ToJSON() (string, error) {
	// 在处理过程中锁定数据库，并在返回时解锁
	d.lock.RLock()
	defer d.lock.RUnlock()

	d.UpdateObfuscatedCallMap()

	bytes, err := json.Marshal(&d.store)

	// 返回零拷贝字符串，因为该字符串将只读
	result := *(*string)(unsafe.Pointer(&bytes))
	return result, err
}

// UpdateObfuscatedCallMap 设置安全调用映射

// ff:
func (d *DB) UpdateObfuscatedCallMap() map[string]int {
	mappings := make(map[string]int)

	for id, k := range d.obfuscatedMethodArray {
		mappings[k.methodName] = id
	}

	return mappings
}
