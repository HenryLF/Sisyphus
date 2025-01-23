package physic

import (
	"math"
)

type ColisionMap = map[int][]Object

func ColisionForce(A, B Object) complex128 {
	ax, _ := Normalize((A.C - B.C))
	E := math.Pow(Overlap(A.Hitbox(A), B.Hitbox(B))*CollisionTransfert*(B.M+A.M)/(A.M), 4)
	out := CmplxMul(ax, E)

	if E > CapImpulse {
		out = CmplxMul(out, CapImpulse/E)
	}
	return out

}

func ColidingObjectMap(L []*Object) ColisionMap {
	out := make(ColisionMap)
	n := len(L)
	if n == 0 {
		return out
	}
	if n == 1 {
		out[0] = []Object{}
	}
	for i, a := range L[:n-1] {
		for k := 1; k < n-i; k++ {
			b := L[i+k]
			if Colide(a.Hitbox(*a), b.Hitbox(*b)) {
				out[i] = append(out[i], *b)
				out[k+i] = append(out[k+i], *a)
			}
		}
	}
	return out
}
