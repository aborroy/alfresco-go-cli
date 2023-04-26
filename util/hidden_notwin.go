//go:build !windows

package util

func IsHiddenFile(filename string) (bool, error) {
	return filename[0] == '.', nil
}
