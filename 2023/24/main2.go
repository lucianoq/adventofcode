package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Given:
// (1) p¹ = p₀¹ + v¹⋅t
// (2) p² = p₀² + v²⋅t
// (3) p³ = p₀³ + v³⋅t
// ...
//
// Find:
// p = p₀ + v⋅t
//
// such as:
// p¹=p when t¹=t
// p²=p when t²=t
// p³=p when t³=t
// ...
//
// where
// pⁿ = [pₓⁿ, pᵧⁿ, pₛⁿ]
// p₀ⁿ = [pₓ₀ⁿ, pᵧⁿ, pₛ₀ⁿ]
// vⁿ = [vₓⁿ, vᵧⁿ, vₛⁿ]

// p¹ + v¹⋅t = p + v⋅t
// p² + v²⋅t = p + v⋅t
// p³ + v³⋅t = p + v⋅t
// pⁿ + vⁿ⋅t = p + v⋅t
//
// (p - p¹) = (v¹ - v)t¹
// (p - p²) = (v² - v)t²
// (p - p³) = (v³ - v)t³
// (p - pⁿ) = (vⁿ - v)tⁿ
// and so:
// (p - pⁿ) ⋅ (v-vⁿ) = 0
func main() {
	hail := parse()

	// Ax = b
	//
	// A = [         0,  -v⁰ₛ +v¹ₛ,   v⁰ᵧ - v¹ᵧ,          0,  -p¹ₛ +p⁰ₛ,   p¹ᵧ -p⁰ᵧ
	//		  v⁰ₛ -v¹ₛ,          0,  -v⁰ₓ + v¹ₓ,   p¹ₛ -p⁰ₛ,          0,  -p¹ₓ +p⁰ₓ
	//		 -v⁰ᵧ +v¹ᵧ,   v⁰ₓ -v¹ₓ,           0,  -p¹ᵧ +p⁰ᵧ,   p¹ₓ -p⁰ₓ,          0
	//               0,  -v⁰ₛ +v²ₛ,   v⁰ᵧ - v²ᵧ,          0,  -p²ₛ +p⁰ₛ,   p²ᵧ -p⁰ᵧ
	//		  v⁰ₛ -v²ₛ,          0,  -v⁰ₓ + v²ₓ,   p²ₛ -p⁰ₛ,          0,  -p²ₓ +p⁰ₓ
	//		 -v⁰ᵧ +v²ᵧ,   v⁰ₓ -v²ₓ,           0,  -p²ᵧ +p⁰ᵧ,   p²ₓ -p⁰ₓ,          0 ]
	//
	// x = [ pₓ
	//       pᵧ
	//       pₛ
	//       vₓ
	//       vᵧ
	//       vₛ ]
	//
	// b = [ p¹ᵧ⋅v¹ₛ -p¹ₛ⋅v¹ᵧ -p⁰ᵧ⋅v⁰ₛ +p⁰ₛ⋅v⁰ᵧ,
	//		 p¹ₛ⋅v¹ₓ -p¹ₓ⋅v¹ₛ -p⁰ₛ⋅v⁰ₓ +p⁰ₓ⋅v⁰ₛ,
	//		 p¹ₓ⋅v¹ᵧ -p¹ᵧ⋅v¹ₓ -p⁰ₓ⋅v⁰ᵧ +p⁰ᵧ⋅v⁰ₓ,
	//       p²ᵧ⋅v²ₛ -p²ₛ⋅v²ᵧ -p⁰ᵧ⋅v⁰ₛ +p⁰ₛ⋅v⁰ᵧ,
	//		 p²ₛ⋅v²ₓ -p²ₓ⋅v²ₛ -p⁰ₛ⋅v⁰ₓ +p⁰ₓ⋅v⁰ₛ,
	//		 p²ₓ⋅v²ᵧ -p²ᵧ⋅v²ₓ -p⁰ₓ⋅v⁰ᵧ +p⁰ᵧ⋅v⁰ₓ ]
	//

	// hail[0] with hail[1]
	// (p0 - p[1]) x (v0 - v[1]) == 0
	A00 := diff(crossMatrix(hail[0].Vel), crossMatrix(hail[1].Vel))
	A03 := diff(crossMatrix(hail[1].Pos), crossMatrix(hail[0].Pos))

	// hail[0] with hail[2]
	// (p0 - p[2]) x (v0 - v[2]) == 0
	A30 := diff(crossMatrix(hail[0].Vel), crossMatrix(hail[2].Vel))
	A33 := diff(crossMatrix(hail[2].Pos), crossMatrix(hail[0].Pos))

	A := mat.NewDense(6, 6, []float64{
		A00[0], A00[1], A00[2], A03[0], A03[1], A03[2],
		A00[3], A00[4], A00[5], A03[3], A03[4], A03[5],
		A00[6], A00[7], A00[8], A03[6], A03[7], A03[8],
		A30[0], A30[1], A30[2], A33[0], A33[1], A33[2],
		A30[3], A30[4], A30[5], A33[3], A33[4], A33[5],
		A30[6], A30[7], A30[8], A33[6], A33[7], A33[8],
	})

	b0 := diff(hail[1].Pos.cross(hail[1].Vel).toF(), hail[0].Pos.cross(hail[0].Vel).toF())
	b3 := diff(hail[2].Pos.cross(hail[2].Vel).toF(), hail[0].Pos.cross(hail[0].Vel).toF())

	b := mat.NewVecDense(6, []float64{b0[0], b0[1], b0[2], b3[0], b3[1], b3[2]})

	var x mat.VecDense
	_ = x.SolveVec(A, b)

	rock := HailStone{
		Pos: P{x.At(0, 0), x.At(1, 0), x.At(2, 0)},
		Vel: P{x.At(3, 0), x.At(4, 0), x.At(5, 0)},
	}

	fmt.Printf("%.0f\n", rock.Pos.x+rock.Pos.y+rock.Pos.z)
}

func crossMatrix(p P) []float64 {
	return []float64{
		0, -p.z, p.y,
		p.z, 0, -p.x,
		-p.y, p.x, 0,
	}
}

func diff(a, b []float64) []float64 {
	res := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] - b[i]
	}
	return res
}

// https://wikimedia.org/api/rest_v1/media/math/render/svg/3242bd71d63c393d02302c7dbe517cd0ec352d31
// https://en.wikipedia.org/wiki/Cross_product#Coordinate_notation
func (p P) cross(p2 P) P {
	return P{
		p.y*p2.z - p.z*p2.y,
		p.z*p2.x - p.x*p2.z,
		p.x*p2.y - p.y*p2.x,
	}
}

func (p P) toF() []float64 {
	return []float64{p.x, p.y, p.z}
}
