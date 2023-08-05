package gos

import (
	"os/exec"
)

func Exec(cmdStr string, args ...string) (string, error) {

	// 使用命令行参数
	cmd := exec.Command(cmdStr, args...)

	// 或者通过管道获取输出
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
