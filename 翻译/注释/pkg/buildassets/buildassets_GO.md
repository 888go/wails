
<原文开始>
// Same as assets but chrooted into /build/
<原文结束>

# <翻译开始>
// 与 assets 相同，但以 /build/ 为根目录进行操作
# <翻译结束>


<原文开始>
// Install will install all default project assets
<原文结束>

# <翻译开始>
// Install 将安装所有默认项目资源
# <翻译结束>


<原文开始>
// GetLocalPath returns the local path of the requested build asset file
<原文结束>

# <翻译开始>
// GetLocalPath 返回请求构建资源文件的本地路径
# <翻译结束>


<原文开始>
// ReadFile reads the file from the project build folder.
// If the file does not exist it falls back to the embedded file and the file will be written
// to the disk for customisation.
<原文结束>

# <翻译开始>
// ReadFile 从项目构建文件夹中读取文件。
// 如果文件不存在，则回退到嵌入的文件，并将文件写入磁盘以便进行自定义。
# <翻译结束>


<原文开始>
// The file does not exist, let's read it from the assets FS and write it to disk
<原文结束>

# <翻译开始>
// 文件不存在，让我们从资源文件系统读取该文件并将其写入磁盘
# <翻译结束>


<原文开始>
// ReadFileWithProjectData reads the file from the project build folder and replaces ProjectInfo if necessary.
// If the file does not exist it falls back to the embedded file and the file will be written
// to the disk for customisation. The file written is the original unresolved one.
<原文结束>

# <翻译开始>
// ReadFileWithProjectData 从项目构建文件夹读取文件，并在必要时替换 ProjectInfo。
// 如果文件不存在，则回退到嵌入的文件，该文件将被写入磁盘以便进行自定义。
// 写入的文件是原始未解析的文件。
# <翻译结束>


<原文开始>
// ReadOriginalFileWithProjectDataAndSave reads the file from the embedded assets and replaces
// ProjectInfo if necessary.
// It will also write the resolved final file back to the project build folder.
<原文结束>

# <翻译开始>
// ReadOriginalFileWithProjectDataAndSave 从嵌入的资源中读取文件，并在必要时替换项目信息（ProjectInfo）。
// 同时，它还会将解析后的最终文件写回到项目的构建目录中。
# <翻译结束>

