package settings_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/goku/internal/settings"
)

func TestSetUpApplication(t *testing.T) {
	cases := map[string]struct {
		fileSystem         fileSystemMock
		want               settings.Configuration
		errorSavingWanted  error
		errorLoadingWanted error
		input              settings.Configuration
	}{
		"save": {
			fileSystem: fileSystemMock{
				userHomeDir: "/home/dummy",
				existFile:   false,
			},
			want: settings.Configuration{
				DefaultNamespace: "default",
			},
			input: settings.Configuration{
				DefaultNamespace: "default",
			},
		},
		"save_error_no_user_home": {
			fileSystem: fileSystemMock{
				errUserHomeDir: errors.New("no user home"),
			},
			errorSavingWanted: errors.New("could not read user home directory"),
			input: settings.Configuration{
				DefaultNamespace: "default",
			},
		},
		"save_error_no_making_dir": {
			fileSystem: fileSystemMock{
				existFile: true,
				errMkdir:  errors.New("no making dir"),
			},
			errorSavingWanted: errors.New("could not make goku configuration directory"),
			input: settings.Configuration{
				DefaultNamespace: "default",
			},
		},
		"save_error_no_writing": {
			fileSystem: fileSystemMock{
				errWriteFile: errors.New("no writing"),
			},
			errorSavingWanted: errors.New("could not write goku configuration"),
			input: settings.Configuration{
				DefaultNamespace: "default",
			},
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			newSetting := settings.New(&data.fileSystem)
			// When
			errSaving := newSetting.SetUp(data.input)
			if data.errorSavingWanted == nil && errSaving != nil {
				t.Fatalf("unexpected error setting up the application: %s", errSaving)
			}
			var errLoading error
			var setupObtained settings.Configuration
			if data.errorSavingWanted == nil {
				setupObtained, errLoading = loadConfigurationHelper(st, newSetting, data.errorLoadingWanted)
			}
			// Then
			assert.Equal(t, data.errorSavingWanted, errSaving, name)
			assert.Equal(t, data.errorLoadingWanted, errLoading, name)
			assert.Equal(t, data.want, setupObtained, name)
		})
	}
}

func loadConfigurationHelper(t *testing.T, setting *settings.Setting, errorWanted error) (settings.Configuration, error) {
	t.Helper()
	setupObtained, err := setting.LoadConfiguration()
	if errorWanted == nil && err != nil {
		t.Fatalf("unexpected error reading application setup: %s", err)
	}
	return setupObtained, err
}

type fileSystemMock struct {
	fileData       []byte
	errMkdir       error
	errUserHomeDir error
	userHomeDir    string
	errStat        error
	existFile      bool
	errWriteFile   error
}

func (f *fileSystemMock) ReadFile(name string) ([]byte, error) {
	return f.fileData, nil
}

func (f *fileSystemMock) Mkdir(name string, perm fs.FileMode) error {
	if f.errMkdir != nil {
		return f.errMkdir
	}
	return nil
}
func (f *fileSystemMock) UserHomeDir() (string, error) {
	if f.errUserHomeDir != nil {
		return "", f.errUserHomeDir
	}
	return f.userHomeDir, nil
}
func (f *fileSystemMock) Stat(name string) (fs.FileInfo, error) {
	if f.errStat != nil {
		return nil, f.errStat
	}
	return nil, nil
}
func (f *fileSystemMock) IsNotExist(error) bool {
	return f.existFile
}
func (f *fileSystemMock) WriteFile(name string, data []byte, perm fs.FileMode) error {
	f.fileData = data
	return f.errWriteFile
}
