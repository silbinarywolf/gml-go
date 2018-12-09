// +build headless

package gml

type fontData struct {
}

func StringWidth(text string) float64 {
	return 0
}

func DrawSetFont(font FontIndex) {
}

// FontInitializeIndexToName is not used for headless builds
func FontInitializeIndexToName(indexToName []string, nameToIndex map[string]FontIndex) {
}
