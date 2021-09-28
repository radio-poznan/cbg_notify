package cbg_notify

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
)

func CopyFile(from, to string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("copy", "/Y", from, to)
	} else {
		cmd = exec.Command("cp", from, to)
	}

	return cmd.Run()
}


func ReadFileIntoChan(path string, c chan string) error {
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c <- string(fileByte)

	return nil
}

func ErrHandle(e error, msg string) {
	fmt.Println(msg, e)
}