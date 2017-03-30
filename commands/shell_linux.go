package commands

import (
	"os/exec"
	"syscall"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
	"strings"
)

func Command(shell string,outchan chan<- []byte) {
	defer close(outchan)

	cmd := exec.Command("/bin/bash", "-c", shell)

	cmd.SysProcAttr = &syscall.SysProcAttr{ Setpgid : true}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		outchan <- []byte("Error: StdoutPipe error => " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		outchan <- []byte("Error: StderrPipe error => " + err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error: Start => ", err.Error())
		outchan <- []byte("Error: cmd.Start error => " + err.Error())
		return
	}

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line ,err2 := reader.ReadString('\n')

		if err2 != nil || io.EOF == err2 {
			break
		}

		b := []byte(line)
		outchan <- b
	}

	bytesErr, err := ioutil.ReadAll(stderr)

	if err == nil {
		outchan <- bytesErr
	}else{
		outchan <- []byte("Error: Stderr => " + err.Error())
	}
	if err := cmd.Wait(); err != nil {

		fmt.Println("Error: ", err.Error())
		out := []byte("Error: " +err.Error())
		outchan <- out
		return
	}

	return
}


func ResolveShellFilePath(p string) string {
	if strings.Count(p,"") -1 <= 2 {
		return p
	}
	if p[0:2] != "./" {
		command := "./" + p
		return command
	}
	return p
}


