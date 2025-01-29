package physic

import "math"

const G = 9

const MovementAcc = 5
const MovementAirAcc = 3
const JumpV = 3
const AirFriction = 1
const FloorFriction = .8

const HitIncrement = 1e-3
const MaxHit = 2
const HitSlowness = 1.5

const ColisionTransfert = 100
const FloorElasticity = 1
const CapImpulse = 800

const HitRadius = 3
const HitNoRecoil = 2

var Dx = 1e-4
var Eps = 5e-1

var MaxForce = 2e2

var Rouding = math.Pow(10, 5)

var Sisyphus = Object{
	C:             complex(-2.8, -1./10),
	R:             1.0 / 10.0,
	M:             1,
	A:             0,
	Hitbox:        RectHitbox(0.85),
	FloorFriction: 2,
	AirFriction:   1.5,
	FloorReaction: 0.05,
	Bounds:        InBounds(-3, math.Inf(-1), math.Inf(1), math.Inf(1)),
	Meta:          make(map[string]interface{}),
}

var Boulder = Object{
	C:             complex(-2, -1./5.),
	R:             1.0 / 5.0,
	M:             .3,
	A:             0,
	Hitbox:        CircleHitbox,
	FloorFriction: 0.8,
	FloorReaction: 1.5,
	AirFriction:   0.5,
	Bounds:        InBounds(-3, math.Inf(-1), math.Inf(1), math.Inf(1)),
	Meta:          make(map[string]interface{}),
}

var Hades = Object{
	C:             -2 - 1i,
	R:             1.0 / 20.0,
	M:             .3,
	A:             0,
	Hitbox:        NoHitbox,
	FloorFriction: 0.8,
	FloorReaction: 1.5,
	AirFriction:   0.5,
	Bounds:        NoBounds,
	Meta:          make(map[string]interface{}),
}

var ObjectS = []*Object{&Sisyphus, &Boulder}
