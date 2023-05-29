package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"unsafe"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	NewScanner().Scan()
}

// type Mode string

// const (
// 	SSID  Mode = "ssid"
// 	BSSID Mode = "Bssid"
// )

type Scanner struct {
	program string
	args    []string
}

func NewScanner() *Scanner {

	switch runtime.GOOS {
	case "windows":
		return &Scanner{
			program: "C://Windows//System32//netsh.exe",
			args:    []string{"wlan", "show", "networks"},
		}
	default:
		panic("")
	}

}

func (s *Scanner) Scan() {

	output, err := exec.Command(s.program, s.args...).Output()
	fmt.Println(err)
	fmt.Println(transform.String(japanese.ShiftJIS.NewDecoder(), *(*string)(unsafe.Pointer(&output))))
}
