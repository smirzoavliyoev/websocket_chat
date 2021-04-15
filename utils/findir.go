package utils

import (
	"errors"
	"io/ioutil"
	"strings"
)

var (
	ErrNoSuchDirOrFile = errors.New("ERR: no such file or direcory")
)

func GetDirOrFilePathFromRoot(root string, target string) (string, error) {

	RootDirs := strings.Split(root, "/")

	if len(RootDirs) != 0 {
		if RootDirs[len(RootDirs)-1] == target {
			return root, nil
		}
	}

	dirsAndFiles, err := ioutil.ReadDir(root)

	if err != nil {
		return "", err
	}

	for _, v := range dirsAndFiles {
		if v.IsDir() {
			path, err := GetDirOrFilePathFromRoot(root+"/"+v.Name(), target)
			if err == nil {
				return path, nil
			}
		}
	}
	return "", ErrNoSuchDirOrFile
}
