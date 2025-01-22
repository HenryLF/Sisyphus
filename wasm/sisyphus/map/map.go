//go:build js && wasm

package worldmap

import (
	"math"
	"math/rand/v2"
	"sisyphus/playerview"
	"strconv"

	"github.com/aquilax/go-perlin"
)

type WorldMap struct {
	Generator func(float64) float64
}

func (W *WorldMap) SetGenerator(f func(float64) float64) {
	W.Generator = f
}

func (W WorldMap) FloorMap(P playerview.PlayerView) map[string]interface{} {
	out := map[string]interface{}{}
	for x := -P.Deresolve; x <= P.Width+P.Deresolve; x += P.Deresolve {
		rx := float64(x)/P.ScaleX + P.X
		y := (W.Generator(rx) - P.Y) * P.ScaleY
		out[strconv.Itoa(x)] = y
	}
	return out
}

var Perlin = perlin.NewPerlin(3.8, 1.5, 2, rand.Int64())

const a float64 = 20
const b float64 = 1
const c float64 = 5

func Generator(x float64) float64 {
	return -math.Pow(1+1/a, x/b) + Perlin.Noise1D(x)/c
}

var Map = WorldMap{
	Generator: Generator,
}
