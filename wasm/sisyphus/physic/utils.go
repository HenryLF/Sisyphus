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
	dy := fun(x+Dx) - fun(x-Dx)
	return cmplx.Rect(1, math.Atan(dy/(2*Dx)))
}

func FloorDistance(s, e complex128, Floor func(float64) float64) float64 {
	x1 := math.Min(real(s), real(e))
	x2 := math.Max(real(e), real(s))
	v := 0.0
	for x := x1; x < x2-Dx; x += Dx {
		v1 := complex(x, Floor(x))
		v2 := complex(x+Dx, Floor(x+Dx))
		v += Dist(v1, v2)
	}
	return v * (real(e) - real(s)) / math.Abs(real(s)-real(e))
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
	C, W, H complex128
}

func (R Rect) Center() complex128 {
	return R.C
}
func (R Rect) Corner() []complex128 {
	return []complex128{R.C - R.W - R.H, R.C + R.W - R.H, R.C - R.W + R.H, R.C + R.W + R.H}
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
	return rectAxeProjection(B.Center(), B.H, A) &&
		rectAxeProjection(B.Center(), B.W, A) //&&
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
		// min_c := slices.MaxFunc(a.Corner(), func(a, b complex128) int { return int(imag(a) - imag(b)) })
		// return imag(min_c) >= Floor(real(min_c))
		return ColideRectFloor(a, Floor)
	}
}

func ColideRectFloor(R Rect, Floor func(float64) float64) bool {
	x := real(R.C)
	N := VectOf(Floor, x) * (0 - 1i)
	P := []complex128{R.C + R.H, R.C - R.H, R.C + R.W, R.C - R.W}
	c := slices.MinFunc(P, func(a, b complex128) int {
		return int((DotProduct(N, a) - DotProduct(N, b)) * 1000)
	})

	return DotProduct(N, complex(x, Floor(x))) <= DotProduct(c, N)
}

func Overlap(A, B Shape) float64 {
	_, cA := A.(Circle)
	_, cB := B.(Circle)
	switch {
	case cA && cB:
		a := A.(Circle)
		b := B.(Circle)
		return math.Abs(a.R+b.R-Dist(a.C, b.C)) / Dist(a.C, b.C)

	case cA && !cB:
		a := A.(Circle)
		b := B.(Rect)
		return OverlapCircleRect(a, b)

	case cB && !cA:
		a := A.(Rect)
		b := B.(Circle)
		return OverlapCircleRect(b, a)
	default:
		a := A.(Rect)
		b := B.(Rect)
		return OverlapRectRect(b, a)

	}
	// return false
}

func OverlapCircleRect(C Circle, R Rect) float64 {
	aH, H := Normalize(R.H)
	aW, W := Normalize(R.W)
	ax, _ := Normalize(C.C - R.C)
	x := C.C + CmplxMul(ax, C.R)
	K := R.Center()
	w := DotProduct(x, aW) - DotProduct(K, aW)
	h := DotProduct(x, aH) - DotProduct(K, aH)
	return (W/2 - math.Abs(w)) * (H/2 - math.Abs(h)) / math.Pow(Dist(C.C, R.C), 2)
}

func OverlapRectRect(A, B Rect) float64 {
	aH, H := Normalize(A.H)
	aW, W := Normalize(A.W)
	aC := A.Center()
	OverlapW := math.MaxFloat64
	OverlapH := math.MaxFloat64
	for _, c := range B.Corner() {
		w := DotProduct(c, aW) - DotProduct(aC, aW)
		OverlapW = math.Min(OverlapW, math.Abs(w))

		h := DotProduct(c, aH) - DotProduct(aC, aH)
		OverlapH = math.Min(OverlapH, math.Abs(h))
	}
	return (W/2 - OverlapW) * (H/2 - OverlapH) / math.Pow(Dist(A.C, B.C), 2)
}

func OutFloor(A Shape, Floor func(float64) float64) complex128 {
	_, cA := A.(Circle)
	switch {
	case cA:
		a := A.(Circle)
		return CircleOutFloor(a, Floor)
	default:
		a := A.(Rect)
		return RectOutFloor(a, Floor)
	}
}

func CircleOutFloor(C Circle, Floor func(float64) float64) complex128 {
	x, y := Unwrap(C.C)
	return complex(0, Floor(x)-y-C.R)
}

func RectOutFloor(R Rect, Floor func(float64) float64) complex128 {
	x, yC := Unwrap(R.C)
	bot := MaxN(R.Corner(), func(a, b complex128) int { return int(1000 * (imag(a) - imag(b))) }, 2)
	_, angle := cmplx.Polar(bot[0] - bot[1])
	y := math.Tan(angle) * (real(bot[1]) - x)
	return complex(0, Floor(x)-yC-y)
}

func MaxN[T comparable](S []T, fun func(a, b T) int, N int) []T {
	m := []T{}
	for range N {
		m = append(m, slices.MaxFunc(S, fun))
		S = slices.DeleteFunc(S, func(a T) bool { return a == m[len(m)-1] })
	}
	return m
}
