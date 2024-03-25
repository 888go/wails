package binding

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"github.com/888go/wails/internal/typescriptify"

	"github.com/leaanthony/slicer"
	"github.com/888go/wails/internal/logger"
)

type Bindings struct {
	db         *DB
	logger     logger.CustomLogger
	exemptions slicer.StringSlicer

	structsToGenerateTS map[string]map[string]interface{}
	enumsToGenerateTS   map[string]map[string]interface{}
	tsPrefix            string
	tsSuffix            string
	tsInterface         bool
	obfuscate           bool
}

// NewBindings 返回一个新的 Bindings 对象
func NewBindings(logger *logger.Logger, structPointersToBind []interface{}, exemptions []interface{}, obfuscate bool, enumsToBind []interface{}) *Bindings {
	result := &Bindings{
		db:                  newDB(),
		logger:              logger.CustomLogger("Bindings"),
		structsToGenerateTS: make(map[string]map[string]interface{}),
		enumsToGenerateTS:   make(map[string]map[string]interface{}),
		obfuscate:           obfuscate,
	}

	for _, exemption := range exemptions {
		if exemption == nil {
			continue
		}
		name := runtime.FuncForPC(reflect.ValueOf(exemption).Pointer()).Name()
		// 呼呼呼！难道没有更好的方法吗？
		name = strings.TrimSuffix(name, "-fm")
		result.exemptions.Add(name)
	}

	for _, enum := range enumsToBind {
		result.AddEnumToGenerateTS(enum)
	}

	// 将结构体添加到绑定
	for _, ptr := range structPointersToBind {
		err := result.Add(ptr)
		if err != nil {
			logger.Fatal("Error during binding: " + err.Error())
		}
	}

	return result
}

// 将给定的结构体方法添加到 Bindings 中
func (b *Bindings) Add(structPtr interface{}) error {
	methods, err := b.getMethods(structPtr)
	if err != nil {
		return fmt.Errorf("cannot bind value to app: %s", err.Error())
	}

	for _, method := range methods {
		splitName := strings.Split(method.Name, ".")
		packageName := splitName[0]
		structName := splitName[1]
		methodName := splitName[2]

		// 将其作为常规方法添加
		b.db.AddMethod(packageName, structName, methodName, method)
	}
	return nil
}

func (b *Bindings) DB() *DB {
	return b.db
}

func (b *Bindings) ToJSON() (string, error) {
	return b.db.ToJSON()
}

func (b *Bindings) GenerateModels() ([]byte, error) {
	models := map[string]string{}
	var seen slicer.StringSlicer
	var seenEnumsPackages slicer.StringSlicer
	allStructNames := b.getAllStructNames()
	allStructNames.Sort()
	allEnumNames := b.getAllEnumNames()
	allEnumNames.Sort()
	for packageName, structsToGenerate := range b.structsToGenerateTS {
		thisPackageCode := ""
		w := typescriptify.New()
		w.WithPrefix(b.tsPrefix)
		w.WithSuffix(b.tsSuffix)
		w.WithInterface(b.tsInterface)
		w.Namespace = packageName
		w.WithBackupDir("")
		w.KnownStructs = allStructNames
		w.KnownEnums = allEnumNames
		// sort the structs
		var structNames []string
		for structName := range structsToGenerate {
			structNames = append(structNames, structName)
		}
		sort.Strings(structNames)
		for _, structName := range structNames {
			fqstructname := packageName + "." + structName
			if seen.Contains(fqstructname) {
				continue
			}
			structInterface := structsToGenerate[structName]
			w.Add(structInterface)
		}

		// 如果我们为此包定义了枚举类型，也要一并添加它们
		var enums, enumsExist = b.enumsToGenerateTS[packageName]
		if enumsExist {
			for enumName, enum := range enums {
				fqemumname := packageName + "." + enumName
				if seen.Contains(fqemumname) {
					continue
				}
				w.AddEnum(enum)
			}
			seenEnumsPackages.Add(packageName)
		}

		str, err := w.Convert(nil)
		if err != nil {
			return nil, err
		}
		thisPackageCode += str
		seen.AddSlice(w.GetGeneratedStructs())
		models[packageName] = thisPackageCode
	}

	// 将未在包含结构体的包中的枚举添加到模型中
	for packageName, enumsToGenerate := range b.enumsToGenerateTS {
		if seenEnumsPackages.Contains(packageName) {
			continue
		}

		thisPackageCode := ""
		w := typescriptify.New()
		w.WithPrefix(b.tsPrefix)
		w.WithSuffix(b.tsSuffix)
		w.WithInterface(b.tsInterface)
		w.Namespace = packageName
		w.WithBackupDir("")

		for enumName, enum := range enumsToGenerate {
			fqemumname := packageName + "." + enumName
			if seen.Contains(fqemumname) {
				continue
			}
			w.AddEnum(enum)
		}
		str, err := w.Convert(nil)
		if err != nil {
			return nil, err
		}
		thisPackageCode += str
		models[packageName] = thisPackageCode
	}

	// 首先对包名进行排序，以确保输出结果的确定性
	sortedPackageNames := make([]string, 0)
	for packageName := range models {
		sortedPackageNames = append(sortedPackageNames, packageName)
	}
	sort.Strings(sortedPackageNames)

	var modelsData bytes.Buffer
	for _, packageName := range sortedPackageNames {
		modelData := models[packageName]
		if strings.TrimSpace(modelData) == "" {
			continue
		}
		modelsData.WriteString("export namespace " + packageName + " {\n")
		sc := bufio.NewScanner(strings.NewReader(modelData))
		for sc.Scan() {
			modelsData.WriteString("\t" + sc.Text() + "\n")
		}
		modelsData.WriteString("\n}\n\n")
	}
	return modelsData.Bytes(), nil
}

func (b *Bindings) WriteModels(modelsDir string) error {
	modelsData, err := b.GenerateModels()
	if err != nil {
		return err
	}
	// 如果没有任何内容，就不要写
	if len(modelsData) == 0 {
		return nil
	}

	filename := filepath.Join(modelsDir, "models.ts")
	err = os.WriteFile(filename, modelsData, 0o755)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bindings) AddEnumToGenerateTS(e interface{}) {
	enumType := reflect.TypeOf(e)

	var packageName string
	var enumName string
	// 枚举类型应表示为所有可能值的数组
	if hasElements(enumType) {
		enum := enumType.Elem()
		// 通过具有Value和TSName字段的结构体表示的简单枚举
		if enum.Kind() == reflect.Struct {
			_, tsNamePresented := enum.FieldByName("TSName")
			enumT, valuePresented := enum.FieldByName("Value")
			if tsNamePresented && valuePresented {
				packageName = getPackageName(enumT.Type.String())
				enumName = enumT.Type.Name()
			} else {
				return
			}
			// 否则期望通过TSName()方法实现，参考：https://github.com/tkrajina/typescriptify-golang-structs#enums-with-tsname
// （该注释含义：在其他情况下，希望按照 TypeScriptify-golang-structs 项目中关于“带有 TSName() 的枚举”部分的说明，采用 TSName() 方法进行实现。）
		} else {
			packageName = getPackageName(enumType.Elem().String())
			enumName = enumType.Elem().Name()
		}
		if b.enumsToGenerateTS[packageName] == nil {
			b.enumsToGenerateTS[packageName] = make(map[string]interface{})
		}
		if b.enumsToGenerateTS[packageName][enumName] != nil {
			return
		}
		b.enumsToGenerateTS[packageName][enumName] = e
	}
}

func (b *Bindings) AddStructToGenerateTS(packageName string, structName string, s interface{}) {
	if b.structsToGenerateTS[packageName] == nil {
		b.structsToGenerateTS[packageName] = make(map[string]interface{})
	}
	if b.structsToGenerateTS[packageName][structName] != nil {
		return
	}
	b.structsToGenerateTS[packageName][structName] = s

	// 遍历此结构体，并添加任何结构体字段引用
	structType := reflect.TypeOf(s)
	if hasElements(structType) {
		structType = structType.Elem()
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Anonymous {
			continue
		}
		kind := field.Type.Kind()
		if kind == reflect.Struct {
			if !field.IsExported() {
				continue
			}
			fqname := field.Type.String()
			sNameSplit := strings.Split(fqname, ".")
			if len(sNameSplit) < 2 {
				continue
			}
			sName := sNameSplit[1]
			pName := getPackageName(fqname)
			a := reflect.New(field.Type)
			if b.hasExportedJSONFields(field.Type) {
				s := reflect.Indirect(a).Interface()
				b.AddStructToGenerateTS(pName, sName, s)
			}
		} else if hasElements(field.Type) && field.Type.Elem().Kind() == reflect.Struct {
			if !field.IsExported() {
				continue
			}
			fqname := field.Type.Elem().String()
			sNameSplit := strings.Split(fqname, ".")
			if len(sNameSplit) < 2 {
				continue
			}
			sName := sNameSplit[1]
			pName := getPackageName(fqname)
			typ := field.Type.Elem()
			a := reflect.New(typ)
			if b.hasExportedJSONFields(typ) {
				s := reflect.Indirect(a).Interface()
				b.AddStructToGenerateTS(pName, sName, s)
			}
		}
	}
}

func (b *Bindings) SetTsPrefix(prefix string) *Bindings {
	b.tsPrefix = prefix
	return b
}

func (b *Bindings) SetTsSuffix(postfix string) *Bindings {
	b.tsSuffix = postfix
	return b
}

func (b *Bindings) SetOutputType(outputType string) *Bindings {
	if outputType == "interfaces" {
		b.tsInterface = true
	}
	return b
}

func (b *Bindings) getAllStructNames() *slicer.StringSlicer {
	var result slicer.StringSlicer
	for packageName, structsToGenerate := range b.structsToGenerateTS {
		for structName := range structsToGenerate {
			result.Add(packageName + "." + structName)
		}
	}
	return &result
}

func (b *Bindings) getAllEnumNames() *slicer.StringSlicer {
	var result slicer.StringSlicer
	for packageName, enumsToGenerate := range b.enumsToGenerateTS {
		for enumName := range enumsToGenerate {
			result.Add(packageName + "." + enumName)
		}
	}
	return &result
}

func (b *Bindings) hasExportedJSONFields(typeOf reflect.Type) bool {
	for i := 0; i < typeOf.NumField(); i++ {
		jsonFieldName := ""
		f := typeOf.Field(i)
		jsonTag := f.Tag.Get("json")
		if len(jsonTag) == 0 {
			continue
		}
		jsonTagParts := strings.Split(jsonTag, ",")
		if len(jsonTagParts) > 0 {
			jsonFieldName = jsonTagParts[0]
		}
		for _, t := range jsonTagParts {
			if t == "-" {
				continue
			}
		}
		if jsonFieldName != "" {
			return true
		}
	}
	return false
}
