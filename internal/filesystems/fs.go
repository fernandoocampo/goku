package filesystems

import (
	"io/fs"
	"os"
)

type OSFileSystem struct {
}

func (*OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (*OSFileSystem) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}
func (*OSFileSystem) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}
func (*OSFileSystem) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}
func (*OSFileSystem) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}
func (*OSFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}
