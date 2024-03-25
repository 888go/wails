
<原文开始>
// serveFile will try to load the file from the fs.FS and write it to the response
<原文结束>

# <翻译开始>
// serveFile 尝试从 fs.FS 加载文件并将其写入响应
# <翻译结束>


<原文开始>
			// If the URL doesn't end in a slash normally a http.redirect should be done, but that currently doesn't work on
			// WebKit WebViews (macOS/Linux).
			// So we handle this as a specific error
<原文结束>

# <翻译开始>
// 如果URL末尾通常没有斜杠，正常情况下应当执行http重定向，但在当前WebKit WebViews（macOS/Linux）环境下无法正常工作。
// 因此，我们将这种情况视为一个特定错误进行处理
# <翻译结束>


<原文开始>
// Detect MimeType by sniffing the first 512 bytes
<原文结束>

# <翻译开始>
// 通过嗅探前512个字节检测MimeType
# <翻译结束>


<原文开始>
		// Do the custom MimeType sniffing even though http.ServeContent would do it in case
		// of an io.ReadSeeker. We would like to have a consistent behaviour in both cases.
<原文结束>

# <翻译开始>
// 即使在 io.ReadSeeker 的情况下 http.ServeContent 会执行自定义 MimeType 探测，我们也进行自定义 MimeType 探测操作。我们希望在这两种情况下都具有一致的行为。
# <翻译结束>


<原文开始>
// Write the first 512 bytes used for MimeType sniffing
<原文结束>

# <翻译开始>
// 写入前512字节，用于MimeType检测
# <翻译结束>


<原文开始>
// Copy the remaining content of the file
<原文结束>

# <翻译开始>
// 复制文件剩余内容
# <翻译结束>

