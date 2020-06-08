// +build headless

package gml

type cameraSurface struct {
	// no surface for headless mode
}

func cameraClearSurface(index int) {
}

func cameraMaybeAllocSurface(index int) {
}

func cameraDraw(index int) {
}
