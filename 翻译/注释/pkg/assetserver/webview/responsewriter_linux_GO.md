
<原文开始>
	// We can't use os.Pipe here, because that returns files with a finalizer for closing the FD. But the control over the
	// read FD is given to the InputStream and will be closed there.
	// Furthermore we especially don't want to have the FD_CLOEXEC
<原文结束>

# <翻译开始>
// 在这里我们不能使用os.Pipe，因为它返回的文件带有用于关闭文件描述符的终结器。但是，读取文件描述符的控制权交给了InputStream，并将在那里被关闭。
// 此外，我们特别不希望拥有FD_CLOEXEC
# <翻译结束>

