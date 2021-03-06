//go:build !windows
// +build !windows

package io

import "fmt"

// Whether the file is hidden-file
func IsHidden(path string) (bool, error) {
	if len(path) == 0 {
		return false, fmt.Errorf("Invalid path")
	}
	return path[0] == '.', nil
}
