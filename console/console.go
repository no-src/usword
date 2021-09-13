package console

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear 清空控制台
func Clear() {
	if runtime.GOOS == "windows" {
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	} else {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
