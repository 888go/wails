package binding_test

import (
	"github.com/888go/wails/internal/binding/binding_test/binding_test_import/int_package"
	"io/fs"
	"os"
	"testing"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/logger"
)

const expectedTypeAliasBindings = `// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// 这个文件是自动生成的。请勿编辑
import {binding_test} from '../models';
import {int_package} from '../models';

export function Map():Promise<{[key: string]: string}>;

export function MapAlias():Promise<binding_test.MapAlias>;

export function MapWithImportedStructValue():Promise<{[key: string]: int_package.SomeStruct}>;

export function Slice():Promise<Array<string>>;

export function SliceImportedStruct():Promise<Array<int_package.SomeStruct>>;
`

type AliasTest struct{}
type MapAlias map[string]string

func (h *AliasTest) Map() map[string]string                                        { return nil }
func (h *AliasTest) MapAlias() MapAlias                                            { return nil }
func (h *AliasTest) MapWithImportedStructValue() map[string]int_package.SomeStruct { return nil }
func (h *AliasTest) Slice() []string                                               { return nil }
func (h *AliasTest) SliceImportedStruct() []int_package.SomeStruct                 { return nil }

func TestAliases(t *testing.T) {
	// given
	generationDir := t.TempDir()

	// setup
	testLogger := &logger.Logger{}
	b := binding.NewBindings(testLogger, []interface{}{&AliasTest{}}, []interface{}{}, false, []interface{}{})

	// then
	err := b.GenerateGoBindings(generationDir)
	if err != nil {
		t.Fatalf("could not generate the Go bindings: %v", err)
	}

	// then
	rawGeneratedBindings, err := fs.ReadFile(os.DirFS(generationDir), "binding_test/AliasTest.d.ts")
	if err != nil {
		t.Fatalf("could not read the generated bindings: %v", err)
	}

	// then
	generatedBindings := string(rawGeneratedBindings)
	if generatedBindings != expectedTypeAliasBindings {
		t.Fatalf("the generated bindings does not match the expected ones.\nWanted:\n%s\n\nGot:\n%s", expectedTypeAliasBindings,
			generatedBindings)
	}
}
