package templates

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/matryer/is"
)

func TestList(t *testing.T) {

	is2 := is.New(t)
	templateList, err := List()
	is2.NoErr(err)

	is2.Equal(len(templateList), 13)
}

func TestShortname(t *testing.T) {

	is2 := is.New(t)

	vanillaTemplate, err := getTemplateByShortname("vanilla")
	is2.NoErr(err)

	is2.Equal(vanillaTemplate.Name, "Vanilla + Vite")
}

func TestInstall(t *testing.T) {

	is2 := is.New(t)

	// 更改到此文件所在的目录
	_, filename, _, _ := runtime.Caller(0)

	err := os.Chdir(filepath.Dir(filename))
	is2.NoErr(err)

	options := &Options{
		ProjectName:  "test",
		TemplateName: "vanilla",
		AuthorName:   "Lea Anthony",
		AuthorEmail:  "lea.anthony@gmail.com",
	}

	defer func() {
		_ = os.RemoveAll(options.ProjectName)
	}()
	_, _, err = Install(options)
	is2.NoErr(err)

}
