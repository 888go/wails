
<原文开始>
// CreateCommand returns a *Cmd struct that when run, will run the given command + args in the given directory
<原文结束>

# <翻译开始>
// CreateCommand 返回一个 *Cmd 结构体，当运行这个结构体时，将在指定目录下执行给定的命令及参数
# <翻译结束>


<原文开始>
// RunCommand will run the given command + args in the given directory
// Will return stdout, stderr and error
<原文结束>

# <翻译开始>
// RunCommand将在给定的目录中运行给定的命令及参数
// 并将返回stdout（标准输出）、stderr（标准错误输出）和error（错误信息）
# <翻译结束>


<原文开始>
// RunCommandWithEnv will run the given command + args in the given directory and using the specified env.
//
// Env specifies the environment of the process. Each entry is of the form "key=value".
// If Env is nil, the new process uses the current process's environment.
//
// Will return stdout, stderr and error
<原文结束>

# <翻译开始>
// RunCommandWithEnv 将在指定目录下使用给定的环境变量执行命令及参数。
//
// Env 指定了进程的环境变量，其格式为 "key=value"。如果 Env 为 nil，则新进程将使用当前进程的环境变量。
//
// 将返回 stdout（标准输出）、stderr（标准错误输出）和 error（错误信息）。
# <翻译结束>


<原文开始>
// RunCommandVerbose will run the given command + args in the given directory
// Will return an error if one occurs
<原文结束>

# <翻译开始>
// RunCommandVerbose会在给定的目录下执行给定的命令及参数
// 如果发生错误，将会返回错误
# <翻译结束>


<原文开始>
// CommandExists returns true if the given command can be found on the shell
<原文结束>

# <翻译开始>
// CommandExists 返回 true，如果在shell上可以找到给定的命令
# <翻译结束>

