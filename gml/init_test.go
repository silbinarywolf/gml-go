package gml

func init() {
	// Setup
	ObjectInitTypes([]ObjectType{
		ObjDummyPlayer: new(DummyPlayer),
	})
}
