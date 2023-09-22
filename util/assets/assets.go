package assets

import "path/filepath"

func GenerateWorkFilePath(fileName string) string {
	return filepath.Join("./work/" + fileName)
}
