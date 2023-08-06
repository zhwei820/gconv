package gos

import (
	"os/exec"
)

func Exec(cmdStr string, args ...string) (string, error) {

	// 使用命令行参数
	cmd := exec.Command(cmdStr, args...)

	out, err := cmd.Output()
	return string(out), err
}
