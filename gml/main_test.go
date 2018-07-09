package gml

const (
	ObjUndefined   ObjectIndex = 0
	ObjDummyPlayer             = 1
	ObjDummyEnemy              = 2
)

type DummyPlayer struct {
	Object
}

func (_ *DummyPlayer) ObjectIndex() ObjectIndex { return ObjDummyPlayer }

func (_ *DummyPlayer) ObjectName() string { return "DummyPlayer" }

func (inst *DummyPlayer) Create() {
	inst.Size.X = 32
	inst.Size.Y = 32
}

func (_ *DummyPlayer) Update() {}

func (_ *DummyPlayer) Draw() {}

func init() {
	// Setup
	ObjectInitTypes([]ObjectType{
		ObjDummyPlayer: new(DummyPlayer),
	})
}
