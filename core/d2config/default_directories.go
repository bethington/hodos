package d2config

import (
	"os"
	"path/filepath"
)

const (
	od2ConfigDirName  = "Nostos"
	od2ConfigFileName = "config.json"
)

// DefaultConfigPath returns the absolute path for the default config file location
func DefaultConfigPath() string {
	if configDir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(configDir, od2ConfigDirName, od2ConfigFileName)
	}

	return LocalConfigPath()
}

// LocalConfigPath returns the absolute path to the directory of the Nostos executable
func LocalConfigPath() string {
	return filepath.Join(filepath.Dir(os.Args[0]), od2ConfigFileName)
}
