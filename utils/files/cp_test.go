package files_test

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"
	"testing"

	"github.com/andersnormal/pkg/utils/files"

	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	tempDir, err := os.MkdirTemp(os.TempDir(), "empty_test")
	assert.NoError(t, err)

	defer func() { _ = os.RemoveAll(tempDir) }()

	old := strings.Join([]string{tempDir, "example.txt"}, "/")

	f, err := os.Create(old)
	assert.NoError(t, err)

	oldBytes, err := f.Write([]byte("Hello World"))
	assert.NoError(t, err)
	f.Close()

	new := strings.Join([]string{tempDir, "example_copy.txt"}, "/")

	newBytes, err := files.CopyFile(old, new)
	assert.NoError(t, err)

	assert.Equal(t, oldBytes, int(newBytes))

	b, err := ioutil.ReadFile(new)
	assert.NoError(t, err)

	assert.Equal(t, "Hello World", string(b))
}

func TestCopyFileHomeDir(t *testing.T) {
	sr, err := user.Current()
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp(sr.HomeDir, "empty_test")
	assert.NoError(t, err)

	defer func() { _ = os.RemoveAll(tempDir) }()

	old := strings.Join([]string{tempDir, "example.txt"}, "/")

	f, err := os.Create(old)
	assert.NoError(t, err)

	oldBytes, err := f.Write([]byte("Hello World"))
	assert.NoError(t, err)
	f.Close()

	new := strings.Join([]string{tempDir, "example_copy.txt"}, "/")
	newHomeDir := strings.Replace(new, sr.HomeDir, "~", 1)

	newBytes, err := files.CopyFile(old, newHomeDir)
	assert.NoError(t, err)

	assert.Equal(t, oldBytes, int(newBytes))

	b, err := ioutil.ReadFile(new)
	assert.NoError(t, err)

	assert.Equal(t, "Hello World", string(b))
}
