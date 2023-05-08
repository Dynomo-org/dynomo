package repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/viper"
)

const (
	commitBranch = "main"
	orgName      = "Dynapgen"
	ownerName    = "Dynapgen"
	baseBranch   = "main"

	githubRawURLFormat = "https://raw.githubusercontent.com/%s/%s/%s/%s"
)

var (
	authorName  = "Dynapgen"
	authorEmail = "dynapgen@gmail.com"
)

func (r *Repository) getRef(ctx context.Context) (*github.Reference, error) {
	repoName := viper.GetString("GITHUB_REPO_NAME")
	if ref, _, err := r.github.Git.GetRef(ctx, ownerName, repoName, "refs/heads/"+commitBranch); err == nil {
		return ref, nil
	}

	ref, _, err := r.github.Git.GetRef(ctx, ownerName, repoName, "refs/heads/"+baseBranch)
	if err != nil {
		return nil, err
	}

	newRef := &github.Reference{Ref: github.String("refs/heads/" + commitBranch), Object: &github.GitObject{SHA: ref.Object.SHA}}
	ref, _, err = r.github.Git.CreateRef(ctx, ownerName, repoName, newRef)
	if err != nil {
		return nil, err
	}

	return ref, nil
}

func (r *Repository) UploadToGithub(ctx context.Context, param UploadFileParam) (string, error) {
	repoName := viper.GetString("GITHUB_REPO_NAME")
	ref, err := r.getRef(ctx)
	if err != nil {
		return "", err
	}

	fileNameSegments := strings.Split(param.FileName, ".")
	filePathRemote := param.DestinationFolderPath + param.FileName

	tree, _, err := r.github.Git.GetTree(ctx, ownerName, repoName, *ref.Object.SHA, true)
	if err != nil {
		return "", err
	}

	entries := []*github.TreeEntry{}
	for _, entry := range tree.Entries {
		path := entry.GetPath()
		if param.ReplaceIfNameExists && strings.HasPrefix(path, param.DestinationFolderPath) && strings.Contains(path, fileNameSegments[0]) {
			entry.SHA = nil
			entries = append(entries, entry)
		}
	}

	fileByte, err := os.ReadFile(param.FilePathLocal)
	if err != nil {
		return "", err
	}

	fileBase64 := base64.StdEncoding.EncodeToString(fileByte)
	blob, _, err := r.github.Git.CreateBlob(ctx, ownerName, repoName, &github.Blob{Content: github.String(fileBase64), Encoding: github.String("base64")})
	if err != nil {
		return "", err
	}

	entries = append(entries, &github.TreeEntry{Path: github.String(filePathRemote), Type: github.String("blob"), SHA: blob.SHA, Mode: github.String("100644")})
	newTree, _, err := r.github.Git.CreateTree(ctx, ownerName, repoName, *ref.Object.SHA, entries)
	if err != nil {
		return "", err
	}

	parent, _, err := r.github.Repositories.GetCommit(ctx, ownerName, repoName, *ref.Object.SHA, nil)
	if err != nil {
		return "", err
	}

	parent.Commit.SHA = parent.SHA
	date := time.Now()
	commitMessage := "Add " + filePathRemote
	author := &github.CommitAuthor{Date: &github.Timestamp{date}, Name: &authorName, Email: &authorEmail}
	commit := &github.Commit{Author: author, Message: &commitMessage, Tree: newTree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := r.github.Git.CreateCommit(ctx, ownerName, repoName, commit)
	if err != nil {
		return "", err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = r.github.Git.UpdateRef(ctx, ownerName, repoName, ref, false)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(githubRawURLFormat, ownerName, repoName, commitBranch, filePathRemote), nil
}
