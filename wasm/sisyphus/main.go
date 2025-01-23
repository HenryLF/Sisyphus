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
		Jump:  Input.Get("Jump").Bool(),
	}
	BoulderP := physic.Boulder.C
	for i := 0; i < 15; i++ {
		Colision := physic.ColidingObjectMap(physic.ObjectS)
		physic.Sisyphus.PFD(input, worldmap.Map.Generator, Colision[0], 1e-3)
		physic.Boulder.PFD(physic.UserInput{}, worldmap.Map.Generator, Colision[1], 1e-3)
	}
	playerview.P.SlowCenter(physic.Sisyphus.C)
	RotateSisyphus()
	RotateBoulder(BoulderP)
	ScaleBoulder()
	return js.ValueOf(map[string]interface{}{
		"X":        playerview.P.X * playerview.P.ScaleX,
		"Y":        playerview.P.Y * playerview.P.ScaleY,
		"Sisyphus": ObjectToJS(physic.Sisyphus),
		"Boulder":  ObjectToJS(physic.Boulder),
		"Floor":    worldmap.Map.FloorMap(playerview.P),
		"Compass":  CompasDirection(physic.Sisyphus, physic.Boulder),
	})
}

func RotateSisyphus() {
	if physic.Sisyphus.IsGrounded(worldmap.Generator) {
		v := physic.VectOf(worldmap.Generator, real(physic.Sisyphus.C))
		physic.Sisyphus.A = math.Atan(imag(v) / real(v))
	} else {
		physic.Sisyphus.A -= physic.Sisyphus.A * dA
	}
}

func RotateBoulder(st complex128) {
	d := physic.FloorDistance(st, physic.Boulder.C, worldmap.Map.Generator)
	physic.Boulder.A += d / physic.Boulder.R
	// log.Println(physic.Boulder.A)
}

func ScaleBoulder() {
	physic.Boulder.R = 1./10. + math.Abs(real(physic.Boulder.C))/BoulderSizeScale
}

func CompasDirection(A, B physic.Object) map[string]interface{} {
	out := map[string]interface{}{}

	out["D"] = physic.DotProduct(1, B.C-A.C) / playerview.P.ScaleX

	ax, _ := physic.Normalize(B.C - A.C)
	ang := math.Atan(imag(ax) / real(ax))
	if physic.DotProduct(1, ax) > 0 {
		out["A"] = ang
	} else {
		out["A"] = math.Pi - ang
	}
	return out
}

func ObjectToJS(A physic.Object) js.Value {
	out := make(map[string]interface{})
	x, y := playerview.P.ScreenTransform(A.C)
	out["X"] = x
	out["Y"] = y
	out["D"] = objectDir(A)
	out["A"] = A.A
	out["R"] = A.R * playerview.P.ScaleX
	A.UpdateMeta(playerview.P.ScaleX, playerview.P.ScaleY)
	out["Meta"] = A.Meta
	return js.ValueOf(out)
}

func objectDir(A physic.Object) string {
	x := real(A.S)
	y := imag(A.S)
	switch {
	case A.Meta["jump"]:
		return "J"
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
