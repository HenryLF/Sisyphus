//go:build js && wasm

package main

import (
	"strconv"
	"syscall/js"
	worldmap "wasm/map"
	"wasm/physic"
	"wasm/playerview"
)

func main() {
	js.Global().Set("InitView", js.FuncOf(InitView))
	js.Global().Set("GetView", js.FuncOf(GetView))
	js.Global().Set("SetView", js.FuncOf(SetView))
	js.Global().Set("GetUpdate", js.FuncOf(GetUpdate))
	js.Global().Set("GetMap", js.FuncOf(GetMap))
	js.Global().Set("Center", js.FuncOf(Center))
	js.Global().Set("SlowCenter", js.FuncOf(SlowCenter))
	js.Global().Set("GetPlayer", js.FuncOf(GetPlayer))
	js.Global().Set("NewBall", js.FuncOf(NewBall))
	js.Global().Set("GetBalls", js.FuncOf(GetBalls))
	js.Global().Set("SetMapParams", js.FuncOf(SetMapParams))
	ch := make(chan bool)
	<-ch
}

func GetView(this js.Value, n []js.Value) any {
	return js.ValueOf(playerview.P.ToJS())
}

func SetView(this js.Value, n []js.Value) any {
	playerview.P.X = n[0].Float()
	playerview.P.Y = n[1].Float()
	return js.ValueOf(playerview.P.ToJS())
}

func Center(this js.Value, n []js.Value) any {
	playerview.P.Center(physic.Player.C)
	return GetView(this, n)
}

func SlowCenter(this js.Value, n []js.Value) any {
	playerview.P.SlowCenter(physic.Player.C)
	return GetView(this, n)
}
func InitView(this js.Value, n []js.Value) any {
	W := n[0].Int()
	H := n[1].Int()
	playerview.P.Set(W, H)
	return GetView(this, n)
}

func GetUpdate(this js.Value, n []js.Value) any {
	Input := physic.UserInput{
		Up:    n[0].Bool(),
		Left:  n[1].Bool(),
		Down:  n[2].Bool(),
		Right: n[3].Bool(),
	}
	for i := 0; i < 15; i++ {
		Colision := physic.ColidingObjectMap(physic.ObjectS)
		for k, obj := range physic.ObjectS {
			if k == 0 {
				obj.PFD(Input, worldmap.Map.Generator, Colision[k], 1e-3)
			} else {
				obj.PFD(physic.UserInput{}, worldmap.Map.Generator, Colision[k], 1e-3)
			}
		}
	}
	return js.ValueOf(playerview.P.ToJS())
}

func GetMap(this js.Value, n []js.Value) any {
	M := worldmap.Map.FloorMap(playerview.P)
	return js.ValueOf(M)
}

func GetPlayer(this js.Value, n []js.Value) any {
	out := map[string]interface{}{}
	x, y := playerview.P.ScreenTransform(physic.Player.C)
	out["X"] = x
	out["Y"] = y
	out["D"] = objectDir(*physic.Player)
	return js.ValueOf(out)
}

const DirThresh = .5
const DirThreshV = .8

func objectDir(A physic.Object) string {
	x := real(A.S)
	y := imag(A.S)
	switch {
	case x > DirThresh:
		return "R"
	case x < -DirThresh:
		return "L"
	case y < -DirThreshV:
		return "U"
	case y > DirThreshV && !A.IsGrounded(worldmap.Map.Generator):
		return "D"
	default:
		return "I"
	}
}

func NewBall(this js.Value, n []js.Value) any {
	x := n[0].Int()
	y := n[1].Int()
	c := playerview.P.MapTransform(x, y)
	ball := physic.NewObject(c, 1.5/20, .5)
	physic.ObjectS = append(physic.ObjectS, ball)
	return js.ValueOf(len(physic.ObjectS) - 1)
}

func GetBalls(this js.Value, n []js.Value) any {
	out := map[string]interface{}{}
	if len(physic.ObjectS) <= 1 {
		return js.ValueOf(out)
	}
	for i, obj := range physic.ObjectS[1:] {

		if playerview.P.In(obj.C, obj.R) {

			x, y := playerview.P.ScreenTransform(obj.C)
			out[strconv.Itoa(i)] = map[string]interface{}{
				"X": x,
				"Y": y,
			}
		}
	}
	return js.ValueOf(out)
}

func SetMapParams(this js.Value, f []js.Value) any {
	a := f[0].Float()
	b := f[1].Float()
	n := int32(f[2].Int())
	sx := f[3].Float()
	sy := f[4].Float()
	worldmap.Map.NewGenerator(a, b, n, sx, sy)
	return js.ValueOf(true)
}
