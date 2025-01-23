package physic

import "math"

const G = 9

const MovementAcc = 5
const MovementAirAcc = 3
const JumpV = 3
const AirFriction = 1
const FloorFriction = .8

const CollisionTransfert = 100
const CapImpulse = 800

const HitRadius = 1e-3
const HitStrenght = 2

var Dx = 1e-4

var Rouding = math.Pow(10, 5)

var PlayerA = Object{
	C:             -1,
	R:             1.0 / 20.0,
	M:             1,
	A:             0,
	Hitbox:        CircleHitbox,
	FloorFriction: 2,
	AirFriction:   1.5,
	FloorReaction: 0.05,
	Bounds:        NoBounds,
	Meta:          make(map[string]interface{}),
}
var PlayerB = Object{
	C:             1,
	R:             1.0 / 20.0,
	M:             1,
	A:             0,
	Hitbox:        RectHitbox(1),
	FloorFriction: 2,
	AirFriction:   1.5,
	FloorReaction: 0.05,
	Bounds:        NoBounds,
	Meta:          make(map[string]interface{}),
}

var Ball = Object{
	C:             0 - 1.2i,
	R:             1.0 / 5.0,
	M:             .3,
	A:             0,
	Hitbox:        CircleHitbox,
	FloorFriction: 0.8,
	FloorReaction: 1.5,
	AirFriction:   0.5,
	Bounds:        NoBounds,
	Meta:          make(map[string]interface{}),
}

var Net = Object{
	C:             0 + 5i,
	R:             .2,
	M:             1000,
	A:             0,
	Hitbox:        RectHitbox(1),
	FloorFriction: 0.8,
	FloorReaction: 1.5,
	AirFriction:   0.5,
	Bounds:        NoBounds,
	Meta:          make(map[string]interface{}),
}

var ObjectS = []*Object{&PlayerA, &PlayerB, &Ball}
