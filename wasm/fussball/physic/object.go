package physic

import (
	"maps"
	"math"
	"math/cmplx"
)

type Object struct {
	C             complex128
	S             complex128
	R             float64
	M             float64
	A             float64
	Hit           bool
	FloorFriction float64
	AirFriction   float64
	FloorReaction float64
	Hitbox        func(Object) Shape
	Bounds        func(complex128) bool
	Meta          map[string]interface{}
}

func CircleHitbox(A Object) Shape {
	if A.Hit {
		return Circle{A.C, A.R * 3}
	}
	return Circle{A.C, A.R}
}
func RectHitbox(aspectW_H float64) func(A Object) Shape {
	return func(A Object) Shape {
		w := CmplxMul(1, A.R*aspectW_H) * cmplx.Rect(1, A.A)
		h := CmplxMul(0-1i, A.R) * cmplx.Rect(1, A.A)
		if A.Hit {
			return Rect{A.C, w * 3, h}
		}
		return Rect{A.C, w, h}
	}
}

func InBounds(Xm, Ym, XM, YM float64) func(complex128) bool {
	return func(c complex128) bool {
		x, y := Unwrap(c)
		return x > Xm && x < XM && y > Ym && y < YM
	}
}
func NoBounds(c complex128) bool {
	return true
}

func (A Object) Copy() Object {
	cp := A
	cp.Meta = maps.Clone(A.Meta)
	return cp
}

func (A *Object) UpdateMeta(sx, sy float64) {
	A.Meta["X"] = real(A.C) * sx
	A.Meta["Y"] = imag(A.C) * sy
	A.Meta["Hit"] = A.Hit
}

func (A Object) IsGrounded(Floor func(float64) float64) bool {
	return ColideFloor(Circle{A.C, A.R}, Floor)
}
func (A Object) IsBellow(Floor func(float64) float64) bool {
	return imag(A.C)+A.R > Floor(real(A.C))
}
func (A *Object) Rotate(c float64) {
	A.A += c
}

func (A *Object) PFD(Input UserInput, Floor func(float64) float64, Colider []Object, delay float64) {
	//Calculate new Acceleration

	grounded := A.IsGrounded(Floor)
	floorAngle := VectOf(Floor, real(A.C))
	F := Gravity(A.M)
	F += FloorReaction(F, floorAngle, grounded, *A)
	if A.IsBellow(Floor) {
		F += complex(0, -FloorElasticity*(imag(A.C)+A.R-Floor(real(A.C)))/delay)
	}
	F += Fricton(*A, grounded)
	for _, obj := range Colider {
		F += ColisionForce(*A, obj)
	}
	if Input.Hit {
		A.Hit = true
	} else {
		A.Hit = false
		F += Movement(Input, grounded, floorAngle, A)
	}
	F = CmplxMul(F, 1/A.M)
	//Calculate new speed
	A.S += CmplxMul(F, delay)

	A.S += Jump(grounded, floorAngle, Input, A)
	//Calculate new coord apply in still in Bounds
	if nC := A.C + CmplxMul(A.S, delay); A.Bounds(nC) {
		A.C = nC
	}
	//Push object if bellow ground
	// if A.IsBellow(Floor) {
	// 	A.C = complex(real(A.C), Floor(real(A.C))-A.R*(A.Hit+1))
	// 	A.S = complex(real(A.S), math.Min(imag(A.S), 0))
	// }

}

func Gravity(m float64) complex128 {
	return complex(0, m*G)
}

type UserInput struct {
	Up, Left, Down, Right, Hit bool
}

func Movement(Input UserInput, grounded bool, floorAngle complex128, A *Object) complex128 {
	F := complex(0, 0)

	if _, a := cmplx.Polar(floorAngle); math.Abs(a) < 1.05 && grounded {
		if Input.Left {
			F -= CmplxMul(floorAngle, MovementAcc)
		}
		if Input.Right {
			F += CmplxMul(floorAngle, MovementAcc)
		}
	}
	if Input.Left && !grounded {
		F -= CmplxMul(floorAngle, MovementAirAcc)
	}
	if Input.Right && !grounded {
		F += CmplxMul(floorAngle, MovementAirAcc)
	}
	if Input.Down && !grounded {
		F += complex(0, JumpV)
	}
	return CmplxMul(F, A.M)
}

func Jump(grounded bool, floorAngle complex128, Input UserInput, A *Object) complex128 {
	if _, a := cmplx.Polar(floorAngle); Input.Up && grounded && math.Abs(a) < 1.05 {
		return complex(0, -JumpV)
	}
	return 0

}

func Fricton(A Object, grounded bool) complex128 {
	if grounded {
		return CmplxMul(A.S, -A.FloorFriction)
	}
	return CmplxMul(A.S, -A.AirFriction)
}

func FloorReaction(F, floorAngle complex128, grounded bool, A Object) complex128 {
	if grounded {
		return CmplxMul(floorAngle*(0-1i), A.FloorReaction*DotProduct(F, 0+1i))
	}
	return 0 + 0i
}
