package commands

import (
	"os/exec"
	"github.com/axgle/mahonia"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

func Command(shell string,outchan chan<- []byte) {

	defer close(outchan)

	cmd := exec.Command("cmd","/C",shell)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		outchan <- []byte("StdoutPipe: " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		outchan <- []byte("StderrPipe: " + err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		outchan <- []byte("Start: " + err.Error())
		return
	}

	reader := bufio.NewReader(stdout)
	enc := mahonia.NewDecoder("gbk")

	//实时循环读取输出流中的一行内容
	for {
		line ,err2 := reader.ReadString('\n')

		if err2 != nil || io.EOF == err2 {
			break
		}

		b := []byte(enc.ConvertString(line))
		outchan <- b
	}

	bytesErr, err := ioutil.ReadAll(stderr)

	if err == nil {
		out := []byte(enc.ConvertString(string(bytesErr)))
		outchan <- out
	}else{
		outchan <- []byte("Stderr: " + err.Error())
	}

	if err := cmd.Wait(); err != nil {

		fmt.Println("Wait: ", err.Error())
		out := []byte("Wait: " +err.Error())
		outchan <- out
		return
	}


	return
}

func ResolveShellFilePath(p string) string {
	return p
}