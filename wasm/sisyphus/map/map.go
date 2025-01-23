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

const a float64 = 1e3
const b float64 = 2
const c float64 = 1e3
const d float64 = 1e2
const e float64 = 1e1

func Generator(x float64) float64 {
	if x < 0 {
		return 0
	}
	return -math.Sin(math.Pow(x, 1+1/a)/b)*math.Pow(1+1/c, x) + Perlin.Noise1D(x)*math.Pow(1+1/d, x)/e
}

var Map = WorldMap{
	Generator: Generator,
}
