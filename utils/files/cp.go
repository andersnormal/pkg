package files

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// CopyFile ...
func CopyFile(src, dst string) (int64, error) {
	src, err := AbsolutePath(src)
	if err != nil {
		return 0, err
	}

	dst, err = AbsolutePath(dst)
	if err != nil {
		return 0, err
	}

	return copy(src, dst)
}

// AbsolutePath ...
func AbsolutePath(path string) (string, error) {
	path, err := replaceHomeFolder(path)
	if err != nil {
		return "", err
	}

	return filepath.Abs(path)
}

func copy(src, dst string) (int64, error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sfi.Mode().IsRegular() {
		return 0, fmt.Errorf("copyfile: %s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	n, err := io.Copy(destination, source)

	return n, err
}

func replaceHomeFolder(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	var buffer bytes.Buffer
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	_, err = buffer.WriteString(usr.HomeDir)
	if err != nil {
		return "", err
	}

	_, err = buffer.WriteString(strings.TrimPrefix(path, "~"))
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
