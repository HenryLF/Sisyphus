//go:build js && wasm

package main

import (
	"fussball/physic"
	"fussball/playerview"
	"log"
	"math"
	"strconv"
	"syscall/js"
)

func main() {
	js.Global().Set("InitView", js.FuncOf(InitView))
	js.Global().Set("GetUpdate", js.FuncOf(GetUpdate))
	log.Println(physic.Ball.A)
	ch := make(chan bool)
	<-ch
}

var Gamestate = js.Global().Get("gameState")
var UserInput = js.Global().Get("userInput")

func InitView(this js.Value, n []js.Value) any {
	playerview.P.Width = n[0].Int()
	playerview.P.Height = n[1].Int()
	log.Println(js.Global().Get("userInput"))
	Gamestate.Get("View").Set("W", playerview.P.Width)
	Gamestate.Get("View").Set("H", playerview.P.Height)
	return js.ValueOf(true)
}

func ParseUserInput() physic.UserInput {
	return physic.UserInput{
		Up:    UserInput.Get("Up").Bool(),
		Left:  UserInput.Get("Left").Bool(),
		Down:  UserInput.Get("Down").Bool(),
		Right: UserInput.Get("Right").Bool(),
		Hit:   UserInput.Get("Hit").Bool(),
	}
}

func GetUpdate(this js.Value, n []js.Value) any {
	Input := ParseUserInput()
	st := physic.Ball.C

	for range 10 {
		Colide := physic.ColidingObjectMap(physic.ObjectS)
		for k, obj := range physic.ObjectS {
			if k == 0 {
				obj.PFD(Input, Arena, Colide[k], 1e-3)
			} else {
				obj.PFD(physic.UserInput{}, Arena, Colide[k], 1e-3)

			}
		}
	}
	RotateBall(st)
	RotatePlayer(&physic.PlayerA)
	RotatePlayer(&physic.PlayerB)

	playerview.P.SlowCenter(physic.PlayerA.C)
	Gamestate.Get("View").Set("X", playerview.P.X*playerview.P.ScaleX)
	Gamestate.Get("View").Set("Y", playerview.P.Y*playerview.P.ScaleY)

	UpdateFloorMap()

	UpdateObject("PlayerA", physic.PlayerA)
	UpdateObject("PlayerB", physic.PlayerB)
	UpdateObject("Ball", physic.Ball)
	UpdateObject("Net", physic.Net)

	return js.ValueOf(true)
}

func RotatePlayer(A *physic.Object) {
	if A.IsGrounded(Arena) {
		v := physic.VectOf(Arena, real(A.C))
		A.A = math.Atan(imag(v) / real(v))
	} else {
		A.A -= A.A * dA
	}
}

func RotateBall(st complex128) {
	d := physic.FloorDistance(st, physic.Ball.C, Arena)
	physic.Ball.A += d / physic.Ball.R
	// log.Println(d, physic.Ball.A, d/physic.Ball.R, physic.Ball.R)
}

func UpdateObject(key string, A physic.Object) {
	x, y := playerview.P.ScreenTransform(A.C)
	r := A.R * playerview.P.ScaleY
	Gamestate.Get(key).Set("X", x)
	Gamestate.Get(key).Set("Y", y)
	Gamestate.Get(key).Set("A", A.A)
	Gamestate.Get(key).Set("R", r)
	Gamestate.Get(key).Set("D", GetObjectDir(A))

}

func GetObjectDir(A physic.Object) string {
	x := real(A.S)
	y := imag(A.S)
	switch {
	case A.Meta["Hit"]:
		return "J"
	case x > DirThresh:
		return "R"
	case x < -DirThresh:
		return "L"
	case y < -DirThreshV:
		return "U"
	case y > DirThreshV && !A.IsGrounded(Arena):
		return "D"
	default:
		return "I"
	}
}

const W float64 = 1.5

func Arena(x float64) float64 {
	if x > 5 {
		return -math.Exp((x - W)) * 1e-7
	}
	if x < -5 {
		return -math.Exp((-x + W)) * 1e-7
	}
	return 0
}

func UpdateFloorMap() {
	out := map[string]interface{}{}
	for x := -playerview.P.Deresolve; x <= playerview.P.Width+playerview.P.Deresolve; x += playerview.P.Deresolve {
		rx := float64(x)/playerview.P.ScaleX + playerview.P.X
		y := (Arena(rx) - playerview.P.Y) * playerview.P.ScaleY
		out[strconv.Itoa(x)] = y
	}
	Gamestate.Set("Floor", js.ValueOf(out))
}
