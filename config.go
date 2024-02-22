package genenum

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	ConfigEnumHelpers struct {
		IsValid    bool `yaml:"is_valid"`
		Is         bool `yaml:"is"`
		Categories []struct {
			Name   string   `yaml:"name"`
			Values []string `yaml:"values"`
		} `yaml:"categories"`
		Validate  bool `yaml:"validate"`
		String    bool `yaml:"string"`
		AllValues struct {
			VarName  string `yaml:"var_name"`
			FuncName string `yaml:"func_name"`
		} `yaml:"all_values"`
	}

	ConfigEnum struct {
		Name    string            `yaml:"name"`
		Values  []string          `yaml:"values"`
		Helpers ConfigEnumHelpers `yaml:"helpers"`
	}
)

func LoadConfig(path string) ([]*ConfigEnum, error) {
	configContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var result []*ConfigEnum
	if err := yaml.Unmarshal(configContent, &result); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return result, nil
}
