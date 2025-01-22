//go:build js && wasm

package worldmap

import (
	"math/rand/v2"
	"strconv"
	"wasm/playerview"

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

func Generator(x float64) float64 {
	return 5 * Perlin.Noise1D(x/5)
}

func (W *WorldMap) NewGenerator(a, b float64, n int32, sx, sy float64) {
	P := perlin.NewPerlin(a, b, n, rand.Int64())
	Gen := func(x float64) float64 {
		return sy * P.Noise1D(x/sx)
	}
	W.Generator = Gen
}

var Map = WorldMap{
	Generator: Generator,
}
