package gml

import (
	"testing"
)

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

func (_ *DummyPlayer) Create() {}

func (_ *DummyPlayer) Update() {}

func (_ *DummyPlayer) Draw() {}

func init() {
	// Setup
	ObjectInitTypes([]ObjectType{
		ObjDummyPlayer: new(DummyPlayer),
	})
}

// NOTE(Jake): 2018-07-07
//
// Ran:
// - go test -bench=.
// Results are:
// - 440 ns/op
// - 457 ns/op
// - 457 ns/op
// Entities:
// - 250 "wall" solid entities
// - 1 player entity
// - Player entity calling "PlaceFree"
//
func BenchmarkPlaceFree250(b *testing.B) {
	roomInstance := RoomInstanceEmptyCreate()
	// Create solid instances to test against
	// NOTE(Jake): 2018-07-07
	//
	// Haven't written collision types for objects yet, so
	// everything is considered solid.
	//
	for i := 0; i < 250; i++ {
		roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
	}
	playerInstance := roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer).(*DummyPlayer)

	for n := 0; n < b.N; n++ {
		PlaceFree(playerInstance, V(32, 32))
	}
}

// NOTE(Jake): 2018-07-07
//
// Ran:
// - go test -bench=.
// Results are:
// - 863 ns/op
// - 906 ns/op
// - 894 ns/op
// Entities:
// - 500 "wall" solid entities
// - 1 player entity
// - Player entity calling "PlaceFree"
//
func BenchmarkPlaceFree500(b *testing.B) {
	roomInstance := RoomInstanceEmptyCreate()
	// Create solid instances to test against
	// NOTE(Jake): 2018-07-07
	//
	// Haven't written collision types for objects yet, so
	// everything is considered solid.
	//
	for i := 0; i < 500; i++ {
		roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
	}
	playerInstance := roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer).(*DummyPlayer)

	for n := 0; n < b.N; n++ {
		PlaceFree(playerInstance, V(32, 32))
	}
}

// NOTE(Jake): 2018-07-07
//
// Ran:
// - go test -bench=.
// Results are:
// - 22620706 ns/op
// - 23755288 ns/op
// - 23313041 ns/op
// Entities:
// - 250 "wall" solid entities
// - 1024 moving/non-trivial entities
// - All 1024 moving entities calling "PlaceFree" 10 times.
//
// This means PlaceFree() blows the entire 16ms by 7.313041ms
//
func BenchmarkPlaceFreeMMOCase_250SolidWalls_1024MovingEntities(b *testing.B) {
	roomInstance := RoomInstanceEmptyCreate()
	// Create solid instances to test against
	// NOTE(Jake): 2018-07-07
	//
	// Haven't written collision types for objects yet, so
	// everything is considered solid.
	//
	for i := 0; i < 250; i++ {
		roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
	}
	movingEntityInstances := make([]*DummyPlayer, 1024)
	for i := 0; i < len(movingEntityInstances); i++ {
		movingEntityInstances[i] = roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer).(*DummyPlayer)
	}

	for n := 0; n < b.N; n++ {
		for _, movingEntityInstance := range movingEntityInstances {
			// NOTE(Jake): 2018-07-07
			//
			// Assume each entity would call PlaceFree() at least 10 times each.
			//
			for i := 0; i < 10; i++ {
				PlaceFree(movingEntityInstance, V(32, 32))
			}
		}
	}
}
