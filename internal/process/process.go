package process

import (
	"os"
	"os/exec"
)

// Process 定义了一个可执行的进程
type Process struct {
	cmd         *exec.Cmd
	exitChannel chan bool
	Running     bool
}

// NewProcess 创建一个新的 process 结构体

// ff:
// args:
// cmd:
func NewProcess(cmd string, args ...string) *Process {
	result := &Process{
		cmd:         exec.Command(cmd, args...),
		exitChannel: make(chan bool, 1),
	}
	result.cmd.Stdout = os.Stdout
	result.cmd.Stderr = os.Stderr
	return result
}

// Start the process

// ff:
// exitCodeChannel:
func (p *Process) Start(exitCodeChannel chan int) error {
	err := p.cmd.Start()
	if err != nil {
		return err
	}

	p.Running = true

	go func(cmd *exec.Cmd, running *bool, exitChannel chan bool, exitCodeChannel chan int) {
		err := cmd.Wait()
		if err == nil {
			exitCodeChannel <- 0
		}
		*running = false
		exitChannel <- true
	}(p.cmd, &p.Running, p.exitChannel, exitCodeChannel)

	return nil
}

// Kill the process

// ff:
func (p *Process) Kill() error {
	if !p.Running {
		return nil
	}
	err := p.cmd.Process.Kill()
	if err != nil {
		return err
	}
	err = p.cmd.Process.Release()
	if err != nil {
		return err
	}

	// 等待命令正确退出
	<-p.exitChannel

	return err
}

// PID 返回进程的PID（进程标识符）

// ff:
func (p *Process) PID() int {
	return p.cmd.Process.Pid
}


// ff:
// dir:
func (p *Process) SetDir(dir string) {
	p.cmd.Dir = dir
}
