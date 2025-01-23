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

const IncrementSuperJump = 1e-3
const MaxSuperJump = 2
const JumpSlowness = 1.5

var Dx = 1e-4

var Rouding = math.Pow(10, 5)
