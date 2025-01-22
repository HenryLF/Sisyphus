//go:build js && wasm

package main

import (
	"math"
	worldmap "sisyphus/map"
	"sisyphus/physic"
	"sisyphus/playerview"
	"syscall/js"
)

func main() {
	js.Global().Set("InitView", js.FuncOf(InitView))
	js.Global().Set("GetView", js.FuncOf(GetView))
	js.Global().Set("GetUpdate", js.FuncOf(GetUpdate))
	ch := make(chan bool)
	<-ch
}

func InitView(this js.Value, n []js.Value) any {
	W := n[0].Int()
	H := n[1].Int()
	playerview.P.Set(W, H)
	return GetView(this, n)
}

func GetView(this js.Value, n []js.Value) any {
	return js.ValueOf(playerview.P.ToJS())
}

var Input = js.Global().Get("userInput")

func GetUpdate(this js.Value, n []js.Value) any {
	input := physic.UserInput{
		Up:    Input.Get("Up").Bool(),
		Left:  Input.Get("Left").Bool(),
		Down:  Input.Get("Down").Bool(),
		Right: Input.Get("Right").Bool(),
	}
	for i := 0; i < 100; i++ {
		Colision := physic.ColidingObjectMap(physic.ObjectS)
		physic.Sisyphus.PFD(input, worldmap.Map.Generator, Colision[0], 1e-4)
		physic.Boulder.PFD(physic.UserInput{}, worldmap.Map.Generator, Colision[1], 1e-4)
	}
	playerview.P.SlowCenter(physic.Sisyphus.C)
	RotateSisyphus()
	return js.ValueOf(map[string]interface{}{
		"Sisyphus": ObjectToJS(physic.Sisyphus),
		"Boulder":  ObjectToJS(physic.Boulder),
		"Floor":    worldmap.Map.FloorMap(playerview.P),
	})
}

func RotateSisyphus() {
	if physic.Sisyphus.IsGrounded(worldmap.Generator) {
		v := physic.VectOf(worldmap.Generator, real(physic.Sisyphus.C))
		physic.Sisyphus.A = math.Tan(imag(v) / real(v))
	} else {
		physic.Sisyphus.A = 0
	}
}

func ObjectToJS(A physic.Object) js.Value {
	out := make(map[string]interface{})
	x, y := playerview.P.ScreenTransform(A.C)
	out["X"] = x
	out["Y"] = y
	out["D"] = objectDir(A)
	out["A"] = A.A
	out["R"] = A.R * playerview.P.ScaleX
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
