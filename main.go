package main

import (
	"fmt"
	"os/exec"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	var out, _ = exec.Command("C://Windows//System32//netsh.exe", "wlan", "show", "networks").Output()
	fmt.Println(transform.String(japanese.ShiftJIS.NewDecoder(), string(out)))
}
