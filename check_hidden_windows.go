//go:build windows
// +build windows

package io

import (
	"path/filepath"
	"syscall"
)

func isHidden(path string) (bool, error) {
	// This func code is from https://freshman.tech/snippets/go/detect-hidden-file/
	// Author is Ayo
	// At 2022-04-08

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}

	// Appending `\\?\` to the absolute path helps with
	// preventing 'Path Not Specified Error' when accessing
	// long paths and filenames
	// https://docs.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation?tabs=cmd
	pointer, err := syscall.UTF16PtrFromString(`\\?\` + absPath)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}

	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
