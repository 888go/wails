
<原文开始>
// RelativeToCwd returns an absolute path based on the cwd
// and the given relative path
<原文结束>

# <翻译开始>
// RelativeToCwd 返回一个基于当前工作目录（cwd）和给定的相对路径的绝对路径
# <翻译结束>


<原文开始>
// Mkdir will create the given directory
<原文结束>

# <翻译开始>
// Mkdir 将创建给定的目录
# <翻译结束>


<原文开始>
// MkDirs creates the given nested directories.
// Returns error on failure
<原文结束>

# <翻译开始>
// MkDirs 创建给定的嵌套目录。
// 若创建失败，返回错误
# <翻译结束>


<原文开始>
// MoveFile attempts to move the source file to the target
// Target is a fully qualified path to a file *name*, not a
// directory
<原文结束>

# <翻译开始>
// MoveFile尝试将源文件移动到目标位置
// 目标是一个指向文件名的完整路径，而不是一个目录
# <翻译结束>


<原文开始>
// DeleteFile will delete the given file
<原文结束>

# <翻译开始>
// DeleteFile 将会删除给定的文件
# <翻译结束>


<原文开始>
// CopyFile from source to target
<原文结束>

# <翻译开始>
// CopyFile 从源文件复制到目标文件
# <翻译结束>


<原文开始>
// DirExists - Returns true if the given path resolves to a directory on the filesystem
<原文结束>

# <翻译开始>
// DirExists - 如果给定的路径在文件系统中解析为一个目录，则返回true
# <翻译结束>


<原文开始>
// FileExists returns a boolean value indicating whether
// the given file exists
<原文结束>

# <翻译开始>
// FileExists 返回一个布尔值，表示给定的文件是否存在
# <翻译结束>


<原文开始>
// RelativePath returns a qualified path created by joining the
// directory of the calling file and the given relative path.
//
// Example: RelativePath("..") in *this* file would give you '/path/to/wails2/v2/internal`
<原文结束>

# <翻译开始>
// RelativePath 函数返回一个由调用文件所在目录与给定的相对路径组合而成的完整路径。
//
// 示例：在 *本* 文件中调用 RelativePath("..") 将会得到 '/path/to/wails2/v2/internal'
# <翻译结束>


<原文开始>
// If we have optional paths, join them to the relativepath
<原文结束>

# <翻译开始>
// 如果我们有可选路径，将其与相对路径连接起来
# <翻译结束>


<原文开始>
		// I'm allowing this for 1 reason only: It's fatal if the path
		// supplied is wrong as it's only used internally in Wails. If we get
		// that path wrong, we should know about it immediately. The other reason is
		// that it cuts down a ton of unnecessary error handling.
<原文结束>

# <翻译开始>
		// 我仅出于一个原因允许这样做：如果提供的路径不正确，那将是致命的，因为它只在Wails内部使用。如果我们获取的路径错误，我们应立即得知。另一个原因是，这可以大量减少不必要的错误处理。
# <翻译结束>


<原文开始>
// MustLoadString attempts to load a string and will abort with a fatal message if
// something goes wrong
<原文结束>

# <翻译开始>
// MustLoadString尝试加载一个字符串，如果出现任何错误，将会输出一条致命消息并终止程序
# <翻译结束>


<原文开始>
// MD5File returns the md5sum of the given file
<原文结束>

# <翻译开始>
// MD5File 返回给定文件的 md5 哈希值
# <翻译结束>


<原文开始>
// MustMD5File will call MD5File and abort the program on error
<原文结束>

# <翻译开始>
// MustMD5File将会调用MD5File函数，并在出现错误时终止程序运行
# <翻译结束>


<原文开始>
// MustWriteString will attempt to write the given data to the given filename
// It will abort the program in the event of a failure
<原文结束>

# <翻译开始>
// MustWriteString 将尝试将给定的数据写入给定的文件名
// 如果发生失败，它将中止程序
# <翻译结束>


<原文开始>
// fatal will print the optional messages and die
<原文结束>

# <翻译开始>
// fatal会打印可选的消息并终止程序
# <翻译结束>


<原文开始>
// GetSubdirectories returns a list of subdirectories for the given root directory
<原文结束>

# <翻译开始>
// GetSubdirectories 返回给定根目录下的子目录列表
# <翻译结束>


<原文开始>
// If we have a directory, save it
<原文结束>

# <翻译开始>
// 如果我们有一个目录，保存它
# <翻译结束>


<原文开始>
// CREDIT: https://stackoverflow.com/a/30708914/8325411
<原文结束>

# <翻译开始>
// CREDIT: 代码来源：https://stackoverflow.com/a/30708914/8325411
# <翻译结束>


<原文开始>
// Either not empty or error, suits both cases
<原文结束>

# <翻译开始>
// 不为空或者存在错误，适用于这两种情况
# <翻译结束>


<原文开始>
// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
// Credit: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
<原文结束>

# <翻译开始>
// CopyDir 递归地复制一个目录树，尝试保持原始权限设置。
// 源目录必须存在，目标目录必须不存在。
// 符号链接会被忽略并跳过。
// 来源：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
# <翻译结束>


<原文开始>
// SetPermissions recursively sets file permissions on a directory
<原文结束>

# <翻译开始>
// SetPermissions 递归地设置目录及其下文件的权限
# <翻译结束>


<原文开始>
// CopyDirExtended recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist. It ignores any files or
// directories that are given through the ignore parameter.
// Symlinks are ignored and skipped.
// Credit: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
<原文结束>

# <翻译开始>
// CopyDirExtended递归地复制一个目录树，尝试保持文件权限。
// 源目录必须存在，目标目录必须不存在。它会忽略通过ignore参数给出的所有文件或目录。
// 符号链接会被忽略并跳过。
// 来源：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
# <翻译结束>


<原文开始>
// FindFileInParents searches for a file in the current directory and all parent directories.
// Returns the absolute path to the file if found, otherwise an empty string
<原文结束>

# <翻译开始>
// FindFileInParents 在当前目录及其所有父目录中搜索指定文件。
// 如果找到该文件，则返回该文件的绝对路径，否则返回一个空字符串
# <翻译结束>

