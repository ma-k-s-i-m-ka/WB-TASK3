package main

import (
	"bufio"
	"fmt"
	proc "github.com/shirou/gopsutil/process"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Print(dir + ">")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			err = ExecPipe(line)
		} else {
			err = Exec(line)
		}

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func Exec(line string) error {
	args := strings.Fields(line)

	if len(args) == 0 {
		return nil
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "cd":
		return cd(args)
	case "pwd":
		return pwd()
	case "echo":
		return echo(args)
	case "kill":
		return kill(args)
	case "ps":
		return ps()
	case "exec":
		path, err := exec.LookPath(args[0])
		if err != nil {
			return err
		}
		return syscall.Exec(path, args[1:], os.Environ())
	case "exit":
		os.Exit(0)
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
	return nil
}

func cd(args []string) error {
	if len(args) == 0 {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	}

	return os.Chdir(args[0])
}

func pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}

func echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func kill(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pid specified")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func ps() error {
	processes, err := proc.Processes()
	if err != nil {
		return err
	}

	fmt.Println("NAME\tPID")

	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			return err
		}

		fmt.Printf("%s\t%d\n", name, process.Pid)
	}

	return nil
}

func ExecPipe(line string) error {
	pipeArgs := strings.Split(line, "|")
	var inReader *io.PipeReader
	var outWriter *io.PipeWriter

	for _, pipeCommand := range pipeArgs {
		args := strings.Fields(pipeCommand)
		command := exec.Command(args[0], args[1:]...)

		if inReader != nil {
			command.Stdin = inReader
		}

		if outWriter != nil {
			command.Stdout = outWriter
		}

		err := command.Start()
		if err != nil {
			return err
		}

		inReader, outWriter = io.Pipe()
		command.Stdout = outWriter
		command.Stdin = inReader

		err = command.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}
