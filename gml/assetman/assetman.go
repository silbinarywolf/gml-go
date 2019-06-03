package assetman

var assetManagers []assetManager

var isHasLoadedAll bool

func Register(manager assetManager) {
	assetManagers = append(assetManagers, manager)
}

func LoadAll() {
	if isHasLoadedAll {
		panic("Cannot call LoadAll() more than once")
	}
	for _, manager := range assetManagers {
		manager.LoadAll()
	}
	isHasLoadedAll = true
}

type assetManager interface {
	//DebugUpdateAll()
	LoadAll()
}
