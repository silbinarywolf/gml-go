package game

import "github.com/silbinarywolf/gml-go/gml/alarm"

type WormLag struct {
	LagTimer alarm.Alarm
	YLag     float64
}
