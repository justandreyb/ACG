package lab1

import (
	"main/utils"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"main/types"
)

const INITIAL_TRANSLATE = 50
var scale float64 = 228

func Make() {
	var values []types.CardioidPart

	values = append(
		values,
		types.CardioidPart{
			Color: sdl.Color{G: 255, B: 255},
			Points: getPoints(0, math.Pi / 2),
		},
	)

	values = append(
		values,
		types.CardioidPart{
			Color: sdl.Color{R: 255, B: 255},
			Points: getPoints(math.Pi / 2, math.Pi),
		},
	)

	values = append(
		values,
		types.CardioidPart{
			Color: sdl.Color{G: 255, R: 255},
			Points: getPoints(math.Pi, (3 * math.Pi) / 2),
		},
	)

	values = append(
		values,
		types.CardioidPart{
			Color: sdl.Color{G: 255},
			Points: getPoints((3 * math.Pi) / 2, 2 * math.Pi),
		},
	)

	utils.DrawCardioid(values)
}

func getPoints(leftBorder float64, rightBorder float64) []sdl.Point {
	var points []sdl.Point
	var idx float64

	for idx = leftBorder; idx < rightBorder; idx+=0.001 {
		points = append(points, createPoint(idx))
	}
	return points
}

func createPoint(value float64) sdl.Point {
	var point sdl.Point

	point.X =  2 * INITIAL_TRANSLATE + int32(scale) + int32 (scale * math.Cos(value) * (1 + math.Cos(value)))
	point.Y =  2 * INITIAL_TRANSLATE + int32(scale) + int32 (scale * math.Sin(value) * (1 + math.Cos(value)))

	return point
}



