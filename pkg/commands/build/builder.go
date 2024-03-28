package build

import (
	"github.com/888go/wails/internal/project"
	"github.com/888go/wails/pkg/clilogger"
)

// Builder 定义了一个构建器，该构建器能够构建 Wails 应用程序
type Builder interface {
	SetProjectData(projectData *project.Project)
	BuildFrontend(logger *clilogger.CLILogger) error
	CompileProject(options *Options) error
	OutputFilename(options *Options) string
	CleanUp()
}
