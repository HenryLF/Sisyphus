package physic

import "math/cmplx"

type Object struct {
	C      complex128
	S      complex128
	R      float64
	M      float64
	A      float64
	Hitbox func(Object) Shape
}

var Sisyphus = Object{
	C:      0,
	R:      1.0 / 20.0,
	M:      1,
	Hitbox: SisyphusHitbox,
}
var Boulder = Object{
	C:      1.2,
	R:      1.0 / 20.0,
	M:      1,
	Hitbox: BoulderHitbox,
}
var ObjectS = []*Object{&Sisyphus, &Boulder}

func BoulderHitbox(A Object) Shape {
	return Circle{A.C, A.R}
}

func SisyphusHitbox(A Object) Shape {
	w := CmplxMul(1, A.R/2) * cmplx.Rect(1, A.A)
	h := CmplxMul(0-1i, A.R/2) * cmplx.Rect(1, A.A)
	return Rect{A.C, w, h}
}

func (A Object) ToJS() map[string]interface{} {
	out := map[string]interface{}{}
	out["X"] = real(A.C)
	out["Y"] = imag(A.C)
	return out
}

func (P Object) IsGrounded(Floor func(float64) float64) bool {
	return ColideFloor(Circle{P.C, P.R}, Floor)
}
func (P Object) IsBellow(Floor func(float64) float64) bool {
	return imag(P.C)+P.R > Floor(real(P.C))
}
func (P *Object) Rotate(c float64) {
	P.A += c
}

func (P *Object) PFD(Input UserInput, Floor func(float64) float64, Colider []Object, delay float64) {
	//Calculate new Acceleration
	grounded := P.IsGrounded(Floor)
	floorAngle := VectOf(Floor, real(P.C))
	F := Gravity(P.M)
	F += FloorReaction(F, floorAngle, grounded)
	F += Fricton(P.S, grounded)
	for _, obj := range Colider {
		F += ColisionForce(*P, obj)
	}
	F += Movement(Input, grounded, floorAngle, P.M)
	F = CmplxMul(F, 1/P.M)
	//Calculate new speed
	P.S += CmplxMul(F, delay)

	//Handle Colision

	P.S += Jump(grounded, Input)
	//Calculate new coord
	P.C += CmplxMul(P.S, delay)
	if P.IsGrounded(Floor) {
		P.C = complex(real(P.C), Floor(real(P.C))-P.R)
		P.S = complex(real(P.S), 0)
	}
}

func Gravity(m float64) complex128 {
	return complex(0, m*G)
}

type UserInput struct {
	Up, Left, Down, Right bool
}

func Movement(Input UserInput, grounded bool, floorAngle complex128, mass float64) complex128 {
	F := complex(0, 0)
	if Input.Left && grounded {
		F -= floorAngle * MovementAcc
	} else if Input.Left {
		F -= MovementAirAcc
	}
	if Input.Right && grounded {
		F += floorAngle * MovementAcc
	} else if Input.Right {
		F += MovementAirAcc
	}
	if Input.Down && !grounded {
		F += complex(0, JumpV)
	}
	return CmplxMul(F, mass)
}

func Jump(grounded bool, Input UserInput) complex128 {
	if Input.Up && grounded {
		return complex(0, -JumpV)
	}
	return 0

}

func Fricton(V complex128, grounded bool) complex128 {
	if grounded {
		return CmplxMul(V, -FloorFriction)
	}
	return CmplxMul(V, -AirFriction)
}

func FloorReaction(F, floorAngle complex128, grounded bool) complex128 {
	if grounded {
		return CmplxMul(floorAngle*(0-1i), DotProduct(F, 0+1i))
	}
	return 0 + 0i
}
