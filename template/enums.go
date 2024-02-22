package template

import (
	"embed"
	"fmt"
	"io"
)

//go:embed enums.tpl
var FS embed.FS

func Load() (string, error) {
	file, err := FS.Open("enums.tpl")
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	return string(content), nil
}
