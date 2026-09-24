package base_config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Load reads config from one file or from a comma-separated list of files, then
// unmarshals and validates the result. A list lets public settings live in a
// plaintext file next to an encrypted file that only holds secrets, for example
// "/config/public.yaml,/config/secret.yaml". Maps are merged across files, but
// each setting may only be defined in one file, so no file can silently
// override another.
func Load(file string, config interface{}) error {

	files := configFiles(file)
	if len(files) == 0 {
		return fmt.Errorf("failed to read config file: no config file given")
	}

	vp := viper.New()
	if len(files) == 1 {
		vp.SetConfigFile(files[0])
		if err := vp.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file: %v", err)
		}
	} else if err := mergeConfigFiles(vp, files); err != nil {
		return err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshall the config: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	return nil
}

func configFiles(file string) []string {
	var files []string
	for _, path := range strings.Split(file, ",") {
		if path = strings.TrimSpace(path); path != "" {
			files = append(files, path)
		}
	}
	return files
}

// mergeConfigFiles merges every file into vp and fails when two files define the
// same setting, or when one defines a value where another defines a map.
// Errors name settings and files, never values, since some files hold secrets.
func mergeConfigFiles(vp *viper.Viper, files []string) error {
	definedIn := map[string]string{}
	for _, path := range files {
		fileConfig := viper.New()
		fileConfig.SetConfigFile(path)
		if err := fileConfig.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file %s: %v", path, err)
		}
		for _, key := range fileConfig.AllKeys() {
			for existing, other := range definedIn {
				if key == existing || strings.HasPrefix(key, existing+".") || strings.HasPrefix(existing, key+".") {
					return fmt.Errorf("config setting %q in %s conflicts with %q in %s: define each setting in only one file", key, path, existing, other)
				}
			}
		}
		for _, key := range fileConfig.AllKeys() {
			definedIn[key] = path
		}
		if err := vp.MergeConfigMap(fileConfig.AllSettings()); err != nil {
			return fmt.Errorf("failed to merge config file %s: %v", path, err)
		}
	}
	return nil
}
