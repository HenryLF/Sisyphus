package physic

import (
	"math"
	"math/cmplx"
)

type Object struct {
	C             complex128
	S             complex128
	R             float64
	M             float64
	A             float64
	Jump          float64
	FloorFriction float64
	AirFriction   float64
	FloorReaction float64
	Hitbox        func(Object) Shape
	Meta          map[string]interface{}
}

var Sisyphus = Object{
	C:             1,
	R:             1.0 / 20.0,
	M:             1,
	Hitbox:        SisyphusHitbox,
	FloorFriction: 2,
	AirFriction:   1.5,
	FloorReaction: 0.05,
	Meta:          make(map[string]interface{}),
}
var Boulder = Object{
	C:             1.2,
	R:             1.0 / 5.0,
	M:             .3,
	Hitbox:        BoulderHitbox,
	FloorFriction: 0.8,
	FloorReaction: 1.5,
	AirFriction:   0.5,
	Meta:          make(map[string]interface{}),
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

func (A *Object) UpdateMeta(sx, sy float64) {
	A.Meta["X"] = real(A.C) * sx
	A.Meta["Y"] = imag(A.C) * sy
	A.Meta["Jump"] = A.Jump
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
	F += Fricton(*A, grounded)
	for _, obj := range Colider {
		F += ColisionForce(*A, obj)
	}
	F += Movement(Input, grounded, floorAngle, *A)
	F = CmplxMul(F, 1/A.M)
	//Calculate new speed
	A.S += CmplxMul(F, delay)

	//Handle Colision

	A.S += Jump(grounded, Input, A)
	//Calculate new coord
	A.C += CmplxMul(A.S, delay)
	if A.IsBellow(Floor) {
		A.C = complex(real(A.C), Floor(real(A.C))-A.R)
		A.S = complex(real(A.S), math.Max(imag(A.S), 0))
	}
}

func Gravity(m float64) complex128 {
	return complex(0, m*G)
}

type UserInput struct {
	Up, Left, Down, Right, Jump bool
}

func Movement(Input UserInput, grounded bool, floorAngle complex128, A Object) complex128 {
	F := complex(0, 0)
	if Input.Jump {
		return F
	}
	if Input.Left && grounded {
		F -= complex(MovementAcc-JumpSlowness*A.Jump, 0)
	} else if Input.Left {
		F -= complex(MovementAirAcc-JumpSlowness*A.Jump, 0)
	}
	if Input.Right && grounded {
		F += complex(MovementAcc-JumpSlowness*A.Jump, 0)
	} else if Input.Right {
		F += complex(MovementAirAcc-JumpSlowness*A.Jump, 0)
	}
	if Input.Down && !grounded {
		F += complex(0, JumpV)
	}
	return CmplxMul(F, A.M)
}

func Jump(grounded bool, Input UserInput, A *Object) complex128 {
	if Input.Jump {
		A.Jump = min(A.Jump+IncrementSuperJump, MaxSuperJump)
		A.Meta["jump"] = true
	} else {
		A.Meta["jump"] = false
	}
	if Input.Up && grounded {
		J := A.Jump
		A.Jump = 0
		return complex(0, -JumpV*(J+1))
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
