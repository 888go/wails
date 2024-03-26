package binding

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/888go/wails/internal/fs"

	"github.com/leaanthony/slicer"
)

var (
	mapRegex          *regexp.Regexp
	keyPackageIndex   int
	keyTypeIndex      int
	valueArrayIndex   int
	valuePackageIndex int
	valueTypeIndex    int
)

func init() {
	mapRegex = regexp.MustCompile(`(?:map\[(?:(?P<keyPackage>\w+)\.)?(?P<keyType>\w+)])?(?P<valueArray>\[])?(?:\*?(?P<valuePackage>\w+)\.)?(?P<valueType>.+)`)
	keyPackageIndex = mapRegex.SubexpIndex("keyPackage")
	keyTypeIndex = mapRegex.SubexpIndex("keyType")
	valueArrayIndex = mapRegex.SubexpIndex("valueArray")
	valuePackageIndex = mapRegex.SubexpIndex("valuePackage")
	valueTypeIndex = mapRegex.SubexpIndex("valueType")
}

func (b *Bindings) GenerateGoBindings(baseDir string) error {
	store := b.db.store
	var obfuscatedBindings map[string]int
	if b.obfuscate {
		obfuscatedBindings = b.db.UpdateObfuscatedCallMap()
	}
	for packageName, structs := range store {
		packageDir := filepath.Join(baseDir, packageName)
		err := fs.Mkdir(packageDir)
		if err != nil {
			return err
		}
		for structName, methods := range structs {
			var jsoutput bytes.Buffer
			jsoutput.WriteString(`// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// 这个文件是自动生成的。请勿编辑
`)
			var tsBody bytes.Buffer
			var tsContent bytes.Buffer
			tsContent.WriteString(`// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// 这个文件是自动生成的。请勿编辑
`)
			// 按字母顺序对方法名称进行排序
			methodNames := make([]string, 0, len(methods))
			for methodName := range methods {
				methodNames = append(methodNames, methodName)
			}
			sort.Strings(methodNames)

			var importNamespaces slicer.StringSlicer
			for _, methodName := range methodNames {
				// Get the method details
				methodDetails := methods[methodName]

				// Generate JS
				var args slicer.StringSlicer
				for count := range methodDetails.Inputs {
					arg := fmt.Sprintf("arg%d", count+1)
					args.Add(arg)
				}
				argsString := args.Join(", ")
				jsoutput.WriteString(fmt.Sprintf("\nexport function %s(%s) {", methodName, argsString))
				jsoutput.WriteString("\n")
				if b.obfuscate {
					id := obfuscatedBindings[strings.Join([]string{packageName, structName, methodName}, ".")]
					jsoutput.WriteString(fmt.Sprintf("  return ObfuscatedCall(%d, [%s]);", id, argsString))
				} else {
					jsoutput.WriteString(fmt.Sprintf("  return window['go']['%s']['%s']['%s'](%s);", packageName, structName, methodName, argsString))
				}
				jsoutput.WriteString("\n}\n")

				// Generate TS
				tsBody.WriteString(fmt.Sprintf("\nexport function %s(", methodName))

				args.Clear()
				for count, input := range methodDetails.Inputs {
					arg := fmt.Sprintf("arg%d", count+1)
					entityName := entityFullReturnType(input.TypeName, b.tsPrefix, b.tsSuffix, &importNamespaces)
					args.Add(arg + ":" + goTypeToTypescriptType(entityName, &importNamespaces))
				}
				tsBody.WriteString(args.Join(",") + "):")
// 现在构建Typescript返回类型
// 如果没有返回值或仅返回错误，TS 返回 Promise<void>
// 如果返回单个值，TS 返回 Promise<type>
// 如果返回单个值或错误，TS 返回 Promise<type>
// 如果返回两个值，TS 返回 Promise<type1|type2>
// 否则，TS 返回 Promise<type1>（而不是抛出Go错误？）
// 翻译成中文：
// ```go
// 现在构建 TypeScript 返回类型
// 如果没有返回值或只返回错误，TypeScript 返回 Promise<void>
// 若返回单一类型值，TypeScript 返回 Promise<type>
// 若返回单一类型值或错误，TypeScript 返回 Promise<type>
// 若返回两个值，TypeScript 返回 Promise<type1|type2>
// 否则情况下，TypeScript 返回 Promise<type1>（而非抛出 Go 语言中的错误？）
// 注意：对于 "If returning single value or error, TS returns Promise<type>" 这一句，原文描述可能有误，根据 Typescript 的习惯用法，应为 "如果返回一个值和一个错误，TS 返回 Promise<[type, Error]> 或 Promise<ReturnType>".
				var returnType string
				if methodDetails.OutputCount() == 0 {
					returnType = "Promise<void>"
				} else if methodDetails.OutputCount() == 1 && methodDetails.Outputs[0].TypeName == "error" {
					returnType = "Promise<void>"
				} else {
					outputTypeName := entityFullReturnType(methodDetails.Outputs[0].TypeName, b.tsPrefix, b.tsSuffix, &importNamespaces)
					firstType := goTypeToTypescriptType(outputTypeName, &importNamespaces)
					returnType = "Promise<" + firstType
					if methodDetails.OutputCount() == 2 && methodDetails.Outputs[1].TypeName != "error" {
						outputTypeName = entityFullReturnType(methodDetails.Outputs[1].TypeName, b.tsPrefix, b.tsSuffix, &importNamespaces)
						secondType := goTypeToTypescriptType(outputTypeName, &importNamespaces)
						returnType += "|" + secondType
					}
					returnType += ">"
				}
				tsBody.WriteString(returnType + ";\n")
			}

			importNamespaces.Deduplicate()
			importNamespaces.Each(func(namespace string) {
				tsContent.WriteString("import {" + namespace + "} from '../models';\n")
			})
			tsContent.WriteString(tsBody.String())

			jsfilename := filepath.Join(packageDir, structName+".js")
			err = os.WriteFile(jsfilename, jsoutput.Bytes(), 0o755)
			if err != nil {
				return err
			}
			tsfilename := filepath.Join(packageDir, structName+".d.ts")
			err = os.WriteFile(tsfilename, tsContent.Bytes(), 0o755)
			if err != nil {
				return err
			}
		}
	}
	err := b.WriteModels(baseDir)
	if err != nil {
		return err
	}
	return nil
}

func fullyQualifiedName(packageName string, typeName string) string {
	if len(packageName) > 0 {
		return packageName + "." + typeName
	}

	switch true {
	case len(typeName) == 0:
		return ""
	case typeName == "interface{}" || typeName == "interface {}":
		return "any"
	case typeName == "string":
		return "string"
	case typeName == "error":
		return "Error"
	case
		strings.HasPrefix(typeName, "int"),
		strings.HasPrefix(typeName, "uint"),
		strings.HasPrefix(typeName, "float"):
		return "number"
	case typeName == "bool":
		return "boolean"
	default:
		return "any"
	}
}

func arrayifyValue(valueArray string, valueType string) string {
	if len(valueArray) == 0 {
		return valueType
	}

	return "Array<" + valueType + ">"
}

func goTypeToJSDocType(input string, importNamespaces *slicer.StringSlicer) string {
	matches := mapRegex.FindStringSubmatch(input)
	keyPackage := matches[keyPackageIndex]
	keyType := matches[keyTypeIndex]
	valueArray := matches[valueArrayIndex]
	valuePackage := matches[valuePackageIndex]
	valueType := matches[valueTypeIndex]
// 打印输出，格式化字符串内容为：
// "input=%s, keyPackage=%s, keyType=%s, valueArray=%s, valuePackage=%s, valueType=%s\n"
// 其中，
// %s 将被替换为对应的变量值：
// input、keyPackage、keyType、valueArray、valuePackage、valueType
// 分别代表输入值、键包名、键类型、值数组、值包名和值类型

	// 字节数组是特殊情况
	if valueArray == "[]" && valueType == "byte" {
		return "string"
	}

	// 如果存在任何包，确保它们已被保存
	if len(keyPackage) > 0 {
		importNamespaces.Add(keyPackage)
	}

	if len(valuePackage) > 0 {
		importNamespaces.Add(valuePackage)
	}

	key := fullyQualifiedName(keyPackage, keyType)
	var value string
	if strings.HasPrefix(valueType, "map") {
		value = goTypeToJSDocType(valueType, importNamespaces)
	} else {
		value = fullyQualifiedName(valuePackage, valueType)
	}

	if len(key) > 0 {
		return fmt.Sprintf("{[key: %s]: %s}", key, arrayifyValue(valueArray, value))
	}

	return arrayifyValue(valueArray, value)
}

func goTypeToTypescriptType(input string, importNamespaces *slicer.StringSlicer) string {
	return goTypeToJSDocType(input, importNamespaces)
}

func entityFullReturnType(input, prefix, suffix string, importNamespaces *slicer.StringSlicer) string {
	if strings.ContainsRune(input, '.') {
		nameSpace, returnType := getSplitReturn(input)
		return nameSpace + "." + prefix + returnType + suffix
	}

	return input
}
