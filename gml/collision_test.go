package gml

import (
	"testing"
)

// NOTE(Jake): 2018-07-08
//
// Native: ("go test --bench=.")
// -------
// BenchmarkPlaceFree250-4                                          3000000               424 ns/op
// BenchmarkPlaceFree500-4                                          2000000               877 ns/op
// BenchmarkPlaceFreeMMOCase_250SolidWalls_1024MovingEntities-4         100          23510338 ns/op
//
// JS: ("GOOS=linux gjbt --bench=."")
// ---
// BenchmarkPlaceFree250                                             500000              2326 ns/op
// BenchmarkPlaceFree500                                             300000              3910 ns/op
// BenchmarkPlaceFreeMMOCase_250SolidWalls_1024MovingEntities            10         103200000 ns/op
//

// NOTE(Jake): 2018-07-07
//
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		PlaceFree(playerInstance, V(32, 32))
	}
}

// NOTE(Jake): 2018-07-07
//
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		PlaceFree(playerInstance, V(32, 32))
	}
}

// NOTE(Jake): 2018-07-07
//
// Entities:
// - 250 "wall" solid entities
// - 1024 moving/non-trivial entities
// - All 1024 moving entities calling "PlaceFree" 10 times.
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
		roomInstance.InstanceCreate(V(float64(i*32), 0), ObjDummyPlayer)
	}
	movingEntityInstances := make([]*DummyPlayer, 1024)
	for i := 0; i < len(movingEntityInstances); i++ {
		movingEntityInstances[i] = roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer).(*DummyPlayer)
	}

	b.ResetTimer()
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
