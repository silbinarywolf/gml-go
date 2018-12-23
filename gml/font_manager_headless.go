// +build headless

package gml

func FontLoad(name string) FontIndex {
	return fntUndefined
}

// InitFontGeneratedData is not supported in headless mode
func InitFontGeneratedData(indexToName []string, nameToIndex map[string]FontIndex) {
}
