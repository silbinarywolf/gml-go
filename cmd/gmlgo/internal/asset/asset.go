package asset

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var whitelistExt = map[string]bool{
	".data": true,
	// font
	".ttf":  true,
	".json": true,
}

// CopyAllFilesWithExt was built to copy files specific to certain OS builds
// to the correct folders
func CopyAllFilesWithExt(src, dst, ext string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())
		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}
		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir, os.ModeSymlink:
			// ignore
		default:
			if filepath.Ext(sourcePath) == ext {
				if err := CopyFile(sourcePath, destPath); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// CopyAssetDirectory recursively copies a src directory to a destination.
func CopyAssetDirectory(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		createdFileOrDir := false
		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			switch entry.Name() {
			case "custom",
				"sound",
				"sprite":
				if err := createDir(destPath, 0755); err != nil {
					return err
				}
				if err := copyDirectory(sourcePath, destPath); err != nil {
					return err
				}
				createdFileOrDir = true
			case "font":
				if err := createDir(destPath, 0755); err != nil {
					return err
				}
				if err := copyDirectoryRecursive(sourcePath, destPath); err != nil {
					return err
				}
				createdFileOrDir = true
			}
		case os.ModeSymlink:
			// ignore
		default:
			if _, ok := whitelistExt[filepath.Ext(sourcePath)]; ok {
				if err := CopyFile(sourcePath, destPath); err != nil {
					return err
				}
				createdFileOrDir = true
			}
		}

		if createdFileOrDir {
			isSymlink := entry.Mode()&os.ModeSymlink != 0
			if !isSymlink {
				if err := os.Chmod(destPath, entry.Mode()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func copyDirectory(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			// ignore
		case os.ModeSymlink:
			// ignore
		default:
			if _, ok := whitelistExt[filepath.Ext(sourcePath)]; ok {
				if err := CopyFile(sourcePath, destPath); err != nil {
					return err
				}
				isSymlink := entry.Mode()&os.ModeSymlink != 0
				if !isSymlink {
					if err := os.Chmod(destPath, entry.Mode()); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func copyDirectoryRecursive(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		createdFileOrDir := false
		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := createDir(destPath, 0755); err != nil {
				return err
			}
			if err := copyDirectoryRecursive(sourcePath, destPath); err != nil {
				return err
			}
			createdFileOrDir = true
		case os.ModeSymlink:
			// ignore
		default:
			if _, ok := whitelistExt[filepath.Ext(sourcePath)]; ok {
				if err := CopyFile(sourcePath, destPath); err != nil {
					return err
				}
				createdFileOrDir = true
			}
		}

		if createdFileOrDir {
			isSymlink := entry.Mode()&os.ModeSymlink != 0
			if !isSymlink {
				if err := os.Chmod(destPath, entry.Mode()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Copy copies a src file to a dst file where src and dst are regular files.
func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func createDir(dir string, perm os.FileMode) error {
	if exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CopySymLink copies a symbolic link from src to dst.
func copySymLink(src, dst string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}
	return os.Symlink(link, dst)
}
