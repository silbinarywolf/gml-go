package assetman

import (
	"sync"
)

var assetManagers []assetManager

var isHasLoadedAll bool

func Register(manager assetManager) {
	//if isHasLoadedAll {
	//	panic("Cannot call Register() if asset managers have already been initialized")
	//}
	assetManagers = append(assetManagers, manager)
}

// UnsafeLoadAll is called at initialization time by GML-Go
func UnsafeLoadAll() {
	if isHasLoadedAll {
		panic("Cannot call LoadAll() more than once")
	}
	var wg sync.WaitGroup
	wg.Add(len(assetManagers))
	for _, m := range assetManagers {
		go func(manager assetManager) {
			manager.LoadAll()
			wg.Done()
		}(m)
	}
	wg.Wait()
	isHasLoadedAll = true

}

type assetManager interface {
	//DebugUpdateAll()
	LoadAll()
}
