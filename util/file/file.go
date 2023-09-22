package file

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/xid"
)

type GithubFolderEnum string

const (
	GithubFolderIcon     GithubFolderEnum = "icon"
	GithubFolderArtifact GithubFolderEnum = "artifact"
	GithubFolderKeystore GithubFolderEnum = "keystore"
)

func GenerateUniqueFilename(filename string) string {
	fileNameSegments := strings.Split(filename, ".")
	return fileNameSegments[0] + "-" + xid.New().String() + "." + fileNameSegments[1]
}

func GenerateUniqueGithubFilePath(folder GithubFolderEnum, filename string) string {
	return fmt.Sprintf("%s/%s/%s", time.Now().Format("200601"), folder, GenerateUniqueFilename(filename))
}

func GetFilenameFromPath(fullPath string) string {
	segments := strings.Split(fullPath, "/")
	return segments[len(segments)-1]
}

func GetFileExtensionFromPath(fullPath string) string {
	segments := strings.Split(fullPath, ".")
	return segments[len(segments)-1]
}
