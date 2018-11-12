package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {

	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		// prc := exec.Command("cmd /c tasklist", "/FI", "PID eq", strconv.Itoa(1768))
		prc := exec.Command("cmd", "/C", "tasklist", "/FI", "PID eq "+strconv.Itoa(1768))
		out, err := prc.Output()

		if err != nil {
			panic(err)
		}
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(out)
		str := string(decodeBytes)
		fmt.Println(str)

		for width, start := 0, 0; start < len(out); start += width {
			var r rune
			r, width = utf8.DecodeRune(out[start:])
			fmt.Printf("r=%d(%s), width=%d\n", r, string(r), width)
		}

		// r, _ := utf8.DecodeRune(out)
		// fmt.Printf("%c", r)
		// fmt.Println(utf8.FullRune(out))

	}

	// prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	// out, err := prc.Output()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(out))

}
