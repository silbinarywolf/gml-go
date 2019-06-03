package assetman

import (
	"sync"
)

var assetManagers []assetManager

var isHasLoadedAll bool

func Register(manager assetManager) {
	assetManagers = append(assetManagers, manager)
}

func LoadAll() {
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
