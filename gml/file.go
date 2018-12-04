package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func AssetDirectory() string {
	return file.AssetDirectory
}

func ProgramDirectory() string {
	return file.ProgramDirectory
}

// todo(Jake): 2018-12-02: #21
// Deprecated. Only used in private project, can be removed after we support "data" binary files
// FileStringReadAll will read a file from the "asset" directory, used to be ReadFileAsString
/*func FileStringReadAll(path string) (string, error) {
	path = AssetDirectory() + "/" + path
	fileData, err := file.OpenFile(path)
	if err != nil {
		return "", err
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytesData)), nil
}
*/
