// +build !headless

package gml

import "github.com/hajimehoshi/ebiten"

type cameraSurface struct {
	surface *ebiten.Image
}

func cameraClearSurface(index int) {
	// NOTE(Jake): 2019-01-26
	// We don't render an offscreen image to the screen if
	// only 1 camera is enabled.
	if cameraHasMultipleEnabled() {
		view := &gCameraManager.cameras[index]
		view.surface.Clear()
	}
}

func cameraMaybeAllocSurface(index int) {
	// NOTE(Jake): 2019-01-26
	// We don't render an offscreen image to the screen if
	// only 1 camera is enabled.
	if cameraHasMultipleEnabled() {
		view := &gCameraManager.cameras[index]
		mustCreateNewRenderTarget := false
		if view.surface == nil {
			// Create new camera
			mustCreateNewRenderTarget = true
		} else {
			// Resize camera
			viewSurfaceSize := view.surface.Bounds().Max
			if int(view.Size.X) != viewSurfaceSize.X ||
				int(view.Size.Y) != viewSurfaceSize.Y {
				mustCreateNewRenderTarget = true
			}
		}
		if mustCreateNewRenderTarget {
			image, err := ebiten.NewImage(int(view.Size.X), int(view.Size.Y), ebiten.FilterDefault)
			if err != nil {
				panic(err)
			}
			view.surface = image
		}
	}
}

func cameraDraw(index int) {
	// NOTE(Jake): 2019-01-26
	// We don't render an offscreen image to the screen if
	// only 1 camera is enabled.
	if cameraHasMultipleEnabled() {
		view := &gCameraManager.cameras[index]

		// NOTE(Jake): 2019-01-26
		// op is a global variable in "draw"
		op.GeoM.Reset()
		op.GeoM.Scale(view.scale.X, view.scale.Y)
		op.GeoM.Translate(view.X, view.Y)
		gScreen.DrawImage(view.surface, op)
	}
}
