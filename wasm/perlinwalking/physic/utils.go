package physic

import (
	"math"
	"math/cmplx"
	"slices"
)

func CmplxMul(c complex128, k float64) complex128 {
	return complex(real(c)*k, imag(c)*k)
}

func Normalize(c complex128) (complex128, float64) {
	if c == 0 {
		return 0i, 0
	}
	n := cmplx.Abs(c)
	return CmplxMul(c, 1/n), n
}

func Dist(A, B complex128) float64 {
	return cmplx.Abs(A - B)
}

func DotProduct(A, B complex128) float64 {
	return real(A)*real(B) + imag(A)*imag(B)
}

func Unwrap(C complex128) (float64, float64) {
	return real(C), imag(C)
}

func VectOf(fun func(float64) float64, x float64) complex128 {
	dy := fun(x+dx) - fun(x-dx)
	return cmplx.Rect(1, math.Atan(dy/(2*dx)))
}

// w = v - 2 * (v ∙ n) * n
func Reflect(A, N complex128) complex128 {
	N, _ = Normalize(N)
	return A - 2*CmplxMul(N, DotProduct(N, A))
}

type Shape interface {
	Center() complex128
}

type Circle struct {
	C complex128
	R float64
}

func (C Circle) Center() complex128 {
	return C.C
}

type Rect struct {
	S, W, H complex128
}

func (R Rect) Center() complex128 {
	return R.S + (R.W+R.H)/2
}
func (R Rect) Corner() []complex128 {
	return []complex128{R.S, R.S + R.W, R.S + R.H, R.S + R.W + R.H}
}

func Colide(A, B Shape) bool {
	_, cA := A.(Circle)
	_, cB := B.(Circle)
	switch {
	case cA && cB:
		a := A.(Circle)
		b := B.(Circle)
		return Dist(a.C, b.C) <= a.R+b.R

	case cA && !cB:
		a := A.(Circle)
		b := B.(Rect)
		return ColisionCircleRect(a, b)

	case cB && !cA:
		a := A.(Rect)
		b := B.(Circle)
		return ColisionCircleRect(b, a)
	default:
		a := A.(Rect)
		b := B.(Rect)
		return ColisionRectRect(b, a)

	}
	// return false
}

func ColisionRectRect(A, B Rect) bool {
	return rectAxeProjection(A.Center(), B.H, A) &&
		rectAxeProjection(A.Center(), B.W, A) //&&
	// rectAxeProjection(A.H, B) &&
	// rectAxeProjection(A.W, B)
}
func rectAxeProjection(C, D complex128, R Rect) bool {
	ax, d := Normalize(D)
	for _, c := range R.Corner() {
		if d/2 >= math.Abs(DotProduct(c, ax)-DotProduct(C, ax)) {

			return true
		}
	}
	return false
}

func ColisionCircleRect(C Circle, R Rect) bool {
	aH, h := Normalize(R.H)
	aW, w := Normalize(R.W)
	return h/2 > math.Abs(DotProduct(aH, C.C)-DotProduct(aH, R.Center()))-C.R &&
		w/2 > math.Abs(DotProduct(aW, C.C)-DotProduct(aW, R.Center()))-C.R
}

func ColideFloor(A Shape, Floor func(float64) float64) bool {
	_, c := A.(Circle)
	x := real(A.Center())
	y := imag(A.Center())
	switch c {
	case true:
		a := A.(Circle)
		return a.R+y >= Floor(x)
	default:
		a := A.(Rect)
		min_c := slices.MaxFunc(a.Corner(), func(a, b complex128) int { return int(imag(a) - imag(b)) })
		return imag(min_c) >= Floor(real(min_c))
	}
}
