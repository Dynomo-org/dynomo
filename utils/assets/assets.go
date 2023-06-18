package assets

import "path/filepath"

func GenerateAssetPath(fileName string) string {
	return filepath.Join("./assets/" + fileName)
}
