package settings

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// FileSystem defines underlying file system behavior
type FileSystem interface {
	ReadFile(name string) ([]byte, error)
	Mkdir(name string, perm fs.FileMode) error
	UserHomeDir() (string, error)
	Stat(name string) (fs.FileInfo, error)
	IsNotExist(err error) bool
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

// Setting defines behavior to handle goku configuration file
type Setting struct {
	osfs        FileSystem // fs is the fs implementation we are using to interact with configuration data.
	userHomeDir string
}

// Configuration setup information
type Configuration struct {
	DefaultNamespace string
}

const configurationDir = ".gokucli"
const defaultNamespace = "default"
const configurationFile = ".gokucli/config"

// configuration errors
var (
	errReadingHomeDir             = errors.New("could not read user home directory")
	errWritingConfiguration       = errors.New("could not write goku configuration")
	errMakingConfigurationDir     = errors.New("could not make goku configuration directory")
	errReadingConfiguration       = errors.New("could not read goku configuration")
	errMarshallingConfiguration   = errors.New("could not interpret new configuration file")
	errUnmarshallingConfiguration = errors.New("could not interpret configuration file")
)

// New creates a new setting handler
func New(newfs FileSystem) *Setting {
	newSetting := Setting{
		osfs: newfs,
	}
	return &newSetting
}

// SetUpDefault setup goku with default parameters
func (s *Setting) SetUpDefault() error {
	defaultSettings := Configuration{
		DefaultNamespace: defaultNamespace,
	}
	return s.SetUp(defaultSettings)
}

// SetUp setup goku based on given configuration data.
func (s *Setting) SetUp(configuration Configuration) error {
	output, err := yaml.Marshal(&configuration)
	if err != nil {
		log.Printf("could marshal new configuration data %+v: %s", configuration, err)
		return errMarshallingConfiguration
	}
	configurationFilePath, err := s.getConfigurationFilePath()
	if err != nil {
		log.Printf("unexpected error getting user home dir %q: %s", configurationFilePath, err)
		return errReadingHomeDir
	}
	err = s.createConfigurationDir()
	if err != nil {
		log.Printf("could not create configuration dir %q: %s", configurationDir, err)
		return errMakingConfigurationDir
	}
	err = s.osfs.WriteFile(configurationFilePath, output, 0700)
	if err != nil {
		log.Printf("could not write configuration file %q: %s", configurationFile, err)
		return errWritingConfiguration
	}
	return nil
}

func (s *Setting) LoadConfiguration() (Configuration, error) {
	var configuration Configuration
	configurationRawData, err := s.readConfigurationFile()
	if err != nil {
		log.Println("could not read goku configuration file", err)
		return configuration, errReadingConfiguration
	}
	err = yaml.Unmarshal(configurationRawData, &configuration)
	if err != nil {
		log.Printf("could not unmarshall file %q: %s\n", configurationFile, err)
		return configuration, errUnmarshallingConfiguration
	}
	return configuration, nil
}

func (s *Setting) createConfigurationDir() error {
	if s.existConfigurationDir() {
		return nil
	}
	gokuDir, err := s.getConfigurationFileDir()
	if err != nil {
		log.Println("could not create goku configuaration dir", err)
		return err
	}
	err = s.osfs.Mkdir(gokuDir, 0700)
	if err != nil {
		log.Printf("could not make goku configuration directory %q: %s", configurationFile, err)
		return err
	}
	return nil
}

func (s *Setting) readConfigurationFile() ([]byte, error) {
	configurationFilePath, err := s.getConfigurationFilePath()
	if err != nil {
		log.Printf("unexpected error getting user home dir %q: %s", configurationFilePath, err)
		return nil, errReadingHomeDir
	}
	return s.osfs.ReadFile(configurationFilePath)
}

func (s *Setting) getConfigurationFilePath() (string, error) {
	userHomeDir, err := s.getUserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHomeDir, configurationFile), nil
}

func (s *Setting) getConfigurationFileDir() (string, error) {
	userHomeDir, err := s.getUserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHomeDir, configurationDir), nil
}

func (s *Setting) existConfigurationDir() bool {
	gokuDir, err := s.getConfigurationFileDir()
	if err != nil {
		log.Println("goku configuration dir doesn't exist", err)
		return false
	}
	_, err = s.osfs.Stat(gokuDir)
	return !s.osfs.IsNotExist(err)
}

func (s *Setting) getUserHomeDir() (string, error) {
	if s.userHomeDir != "" {
		return s.userHomeDir, nil
	}
	var err error
	s.userHomeDir, err = s.osfs.UserHomeDir()
	if err != nil {
		return "", err
	}
	return s.userHomeDir, nil
}

// SetFileSystem set a new file system provider
func (s *Setting) SetFileSystem(fsProvider FileSystem) {
	s.osfs = fsProvider
}
