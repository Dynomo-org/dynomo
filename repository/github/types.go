package github

type UploadFileParam struct {
	FilePathLocal       string
	FilePathRemote      string
	ReplaceIfNameExists bool
}
