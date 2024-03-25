//go:build windows

package gomod

const basic string = `module changeme

go 1.17

require github.com/wailsapp/wails/v2 v2.0.0-beta.7

// 将 "github.com/wailsapp/wails/v2" 版本 "v2.0.0-beta.7" 替换为 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"
// 这段Go代码注释表明了进行包路径替换的意图，即将依赖包 "github.com/wailsapp/wails/v2" 的版本 "v2.0.0-beta.7" 替换成本地路径 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"。这通常发生在开发者需要使用自定义版本或者本地开发的库时。
`
const basicUpdated string = `module changeme

go 1.17

require github.com/wailsapp/wails/v2 v2.0.0-beta.20

// 将 "github.com/wailsapp/wails/v2" 版本 "v2.0.0-beta.7" 替换为 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"
// 这段Go代码注释表明了进行包路径替换的意图，即将依赖包 "github.com/wailsapp/wails/v2" 的版本 "v2.0.0-beta.7" 替换成本地路径 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"。这通常发生在开发者需要使用自定义版本或者本地开发的库时。
`

const multilineRequire = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7
)

// 将 "github.com/wailsapp/wails/v2" 版本 "v2.0.0-beta.7" 替换为 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"
// 这段Go代码注释表明了进行包路径替换的意图，即将依赖包 "github.com/wailsapp/wails/v2" 的版本 "v2.0.0-beta.7" 替换成本地路径 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"。这通常发生在开发者需要使用自定义版本或者本地开发的库时。
`
const multilineReplace = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7
)

replace github.com/wailsapp/wails/v2 v2.0.0-beta.7 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
`

const multilineReplaceNoVersion = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7
)

replace github.com/wailsapp/wails/v2 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
`

const multilineReplaceNoVersionBlock = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7
)

replace (
	github.com/wailsapp/wails/v2 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
)
`

const multilineReplaceBlock = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7
)

replace (
	github.com/wailsapp/wails/v2 v2.0.0-beta.7 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
)
`

const multilineRequireUpdated = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20
)

// 将 "github.com/wailsapp/wails/v2" 版本 "v2.0.0-beta.7" 替换为 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"
// 这段Go代码注释表明了进行包路径替换的意图，即将依赖包 "github.com/wailsapp/wails/v2" 的版本 "v2.0.0-beta.7" 替换成本地路径 "C:\Users\leaan\Documents\wails-v2-beta\wails\v2"。这通常发生在开发者需要使用自定义版本或者本地开发的库时。
`

const multilineReplaceUpdated = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20
)

replace github.com/wailsapp/wails/v2 v2.0.0-beta.20 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
`
const multilineReplaceNoVersionUpdated = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20
)

replace github.com/wailsapp/wails/v2 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
`
const multilineReplaceNoVersionBlockUpdated = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20
)

replace (
	github.com/wailsapp/wails/v2 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
)
`

const multilineReplaceBlockUpdated = `module changeme

go 1.17

require (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20
)

replace (
	github.com/wailsapp/wails/v2 v2.0.0-beta.20 => C:\Users\leaan\Documents\wails-v2-beta\wails\v2
)
`
