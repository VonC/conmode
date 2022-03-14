package main

import (
	"fmt"
	"log"

	"github.com/erikgeiser/coninput"
	"golang.org/x/sys/windows"
)

func getConsoleMode(handle uint32) (uintptr, uint32) {
	con, err := windows.GetStdHandle(handle)
	if err != nil {
		log.Fatalf("get stdin handle: %s", err)
	}

	var consoleMode uint32
	err = windows.GetConsoleMode(con, &consoleMode)
	if err != nil {
		log.Fatalf("get console mode: %s", err)
	}
	return uintptr(con), consoleMode
}

func printDefaultConsoleMode() {
	// https://github.com/charmbracelet/bubbletea/issues/121
	// https://github.com/erikgeiser/coninput/blob/main/example/main.go
	// https://github.com/microsoft/terminal/issues/8750#issuecomment-759088381
	conHandle, conMode := getConsoleMode(windows.STD_INPUT_HANDLE)
	fmt.Printf("StdInput consoleMode (handle %d) %d: %s\n", conHandle, conMode, coninput.DescribeInputMode(conMode))

	conHandle, conMode = getConsoleMode(windows.STD_OUTPUT_HANDLE)
	fmt.Printf("StdOutput consoleMode (handle %d) %d: %s\n", conHandle, conMode, DescribeOutputMode(conMode))

	conHandle, conMode = getConsoleMode(windows.STD_ERROR_HANDLE)
	fmt.Printf("StdError consoleMode (handle %d) %d: %s\n", conHandle, conMode, DescribeOutputMode(conMode))

}
