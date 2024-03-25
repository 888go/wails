package shell

import (
	"bytes"
	"os"
	"os/exec"
)

type Command struct {
	command    string
	args       []string
	env        []string
	dir        string
	stdo, stde bytes.Buffer
}

// ff:
// command:
func NewCommand(command string) *Command {
	return &Command{
		command: command,
		env:     os.Environ(),
	}
}

// ff:
// dir:
func (c *Command) Dir(dir string) {
	c.dir = dir
}

// ff:
// value:
// name:
func (c *Command) Env(name string, value string) {
	c.env = append(c.env, name+"="+value)
}

// ff:运行
func (c *Command) Run() error {
	cmd := exec.Command(c.command, c.args...)
	if c.dir != "" {
		cmd.Dir = c.dir
	}
	cmd.Stdout = &c.stdo
	cmd.Stderr = &c.stde
	return cmd.Run()
}

// ff:
func (c *Command) Stdout() string {
	return c.stdo.String()
}

// ff:
func (c *Command) Stderr() string {
	return c.stde.String()
}

// ff:
// args:
func (c *Command) AddArgs(args []string) {
	c.args = append(c.args, args...)
}

// CreateCommand 返回一个 *Cmd 结构体，当运行这个结构体时，将在指定目录下执行给定的命令及参数

// ff:
// args:
// command:
// directory:
func CreateCommand(directory string, command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = directory
	return cmd
}

// RunCommand将在给定的目录中运行给定的命令及参数
// 并将返回stdout（标准输出）、stderr（标准错误输出）和error（错误信息）

// ff:
// args:
// command:
// directory:
func RunCommand(directory string, command string, args ...string) (string, string, error) {
	return RunCommandWithEnv(nil, directory, command, args...)
}

// RunCommandWithEnv 将在指定目录下使用给定的环境变量执行命令及参数。
//
// Env 指定了进程的环境变量，其格式为 "key=value"。如果 Env 为 nil，则新进程将使用当前进程的环境变量。
//
// 将返回 stdout（标准输出）、stderr（标准错误输出）和 error（错误信息）。

// ff:
// args:
// command:
// directory:
// env:
func RunCommandWithEnv(env []string, directory string, command string, args ...string) (string, string, error) {
	cmd := CreateCommand(directory, command, args...)
	cmd.Env = env

	var stdo, stde bytes.Buffer
	cmd.Stdout = &stdo
	cmd.Stderr = &stde
	err := cmd.Run()
	return stdo.String(), stde.String(), err
}

// RunCommandVerbose会在给定的目录下执行给定的命令及参数
// 如果发生错误，将会返回错误

// ff:
// args:
// command:
// directory:
func RunCommandVerbose(directory string, command string, args ...string) error {
	cmd := CreateCommand(directory, command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

// CommandExists 返回 true，如果在shell上可以找到给定的命令

// ff:
// name:
func CommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
