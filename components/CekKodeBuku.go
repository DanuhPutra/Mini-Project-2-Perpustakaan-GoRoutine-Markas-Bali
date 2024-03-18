package components

import (
	"fmt"
	"path/filepath"
	"strings"
)

func kodeBukuExists(kodebuku string) bool {
	files, err := filepath.Glob("books/*.json")
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return true
	}

	for _, file := range files {
		filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		if filename == kodebuku {
			return true
		}
	}
	return false
}