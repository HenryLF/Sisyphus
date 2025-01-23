package playerview

import (
	"math"
)

type PlayerView struct {
	Width, Height, Deresolve int
	X, Y, ScaleX, ScaleY     float64
}

var P = PlayerView{X: 0, Y: 0, Width: 300, Height: 300, Deresolve: 2, ScaleX: 200, ScaleY: 200}

func (P PlayerView) ScreenTransform(c complex128) (int, int) {
	return int(math.Round((real(c) - P.X) * P.ScaleX)), int(math.Round((imag(c) - P.Y) * P.ScaleY))
}

func (P PlayerView) MapTransform(x int, y int) complex128 {
	return complex(float64(x)/P.ScaleX+P.X, float64(y)/P.ScaleY+P.Y)
}

func (P *PlayerView) Center(c complex128) {
	P.X = real(c) - float64(P.Width)/P.ScaleX/2
	P.Y = imag(c) - float64(P.Height)/P.ScaleY/2
}

func (P PlayerView) In(c complex128, r float64) bool {
	x, y := P.ScreenTransform(c + complex(r, r))
	X, Y := P.ScreenTransform(c + complex(-r, -r))
	return x > 0 && y > 0 && X < P.Width && Y < P.Height
}

func ClampAbs(x float64, c float64) float64 {
	return math.Min(math.Max(x, -c), c)
}

func (P *PlayerView) SlowCenter(c complex128) {
	P.X += ClampAbs(math.Pow((real(c)-P.X-float64(P.Width)/P.ScaleX/2)*5e-1, 3), 1)
	P.Y += ClampAbs(math.Pow((imag(c)-P.Y-float64(P.Height)/P.ScaleY/2)*5e-1, 3), 1)
}
func (P *PlayerView) Set(w, h int) {
	P.Width = w
	P.Height = h
}

func (P PlayerView) ToJS() map[string]interface{} {
	out := map[string]interface{}{}
	out["X"] = P.X
	out["Y"] = P.Y
	out["W"] = P.Width
	out["H"] = P.Height
	out["D"] = P.Deresolve
	return out
}
