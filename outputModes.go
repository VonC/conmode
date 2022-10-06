package main

import (
	"strings"

	"golang.org/x/sys/windows"
)

var outputModes = []struct {
	mode uint32
	name string
}{
	{mode: windows.ENABLE_PROCESSED_OUTPUT, name: "ENABLE_PROCESSED_OUTPUT"},
	{mode: windows.ENABLE_WRAP_AT_EOL_OUTPUT, name: "ENABLE_WRAP_AT_EOL_OUTPUT"},
	{mode: windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING, name: "ENABLE_VIRTUAL_TERMINAL_PROCESSING"},
	{mode: windows.ENABLE_LVB_GRID_WORLDWIDE, name: "ENABLE_LVB_GRID_WORLDWIDE"},
}

// ListOutputModes returns the isolated enabled output modes as a list.
func ListOutputModes(mode uint32) []uint32 {
	modes := []uint32{}

	for _, outputMode := range outputModes {
		if mode&outputMode.mode > 0 {
			modes = append(modes, outputMode.mode)
		}
	}

	return modes
}

// ListOutputModeNames returns the isolated enabled output mode names as a list.
func ListOutputModeNames(mode uint32) []string {
	modes := []string{}

	for _, outputMode := range outputModes {
		if mode&outputMode.mode > 0 {
			modes = append(modes, outputMode.name)
		}
	}

	return modes
}

// DescribeOutputMode returns a string containing the names of each enabled output mode.
func DescribeOutputMode(mode uint32) string {
	return strings.Join(ListOutputModeNames(mode), "|")
}
