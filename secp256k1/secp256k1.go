package secp256k1

import (
	"github.com/vanntrong/go-ethkit/utils"
	"math/big"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

var CURVE_P, _ = utils.HexToBigInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F")
var CURVE_A, _ = utils.HexToBigInt("0000000000000000000000000000000000000000000000000000000000000000")

// var CURVE_N, _ = utils.HexToBigInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141")
var G_X, _ = utils.HexToBigInt("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798")
var G_Y, _ = utils.HexToBigInt("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8")

var GPoint = Point{X: G_X, Y: G_Y}

func PointDoubling(P Point) Point {
	// Copy X and Y coordinates of the input point
	x := new(big.Int).Set(P.X)
	y := new(big.Int).Set(P.Y)

	// Constants
	two := big.NewInt(2)
	three := big.NewInt(3)

	// Step 1: Calculate slope = (3 * X^2 + CURVE_A) / (2 * Y) mod CURVE_P
	// Part 1: Calculate 3 * X^2 + CURVE_A
	xSquared := new(big.Int).Mul(x, x)               // X^2
	threeXsq := new(big.Int).Mul(xSquared, three)    // 3 * X^2
	numerator := new(big.Int).Add(threeXsq, CURVE_A) // 3 * X^2 + CURVE_A

	// Part 2: Calculate 2 * Y
	denominator := new(big.Int).Mul(y, two) // 2 * Y

	// Part 3: Calculate the modular inverse of 2 * Y
	denominatorInv := new(big.Int).ModInverse(denominator, CURVE_P)

	// Part 4: Calculate the slope
	slope := new(big.Int).Mul(numerator, denominatorInv) // (3 * X^2 + CURVE_A) * (2 * Y)^-1
	slope.Mod(slope, CURVE_P)                            // mod CURVE_P

	// Step 2: Calculate x3 = slope^2 - 2 * X mod CURVE_P
	slopeSquared := new(big.Int).Mul(slope, slope) // slope^2
	twoX := new(big.Int).Mul(x, two)               // 2 * X
	x3 := new(big.Int).Sub(slopeSquared, twoX)     // slope^2 - 2 * X
	x3.Mod(x3, CURVE_P)                            // mod CURVE_P

	// Step 3: Calculate y3 = slope * (X - x3) - Y mod CURVE_P
	xMinusX3 := new(big.Int).Sub(x, x3)                 // X - x3
	slopeTimesDiff := new(big.Int).Mul(slope, xMinusX3) // slope * (X - x3)
	y3 := new(big.Int).Sub(slopeTimesDiff, y)           // slope * (X - x3) - Y
	y3.Mod(y3, CURVE_P)                                 // mod CURVE_P

	// Return the new point
	return Point{X: x3, Y: y3}
}

func PointAddition(P, Q Point) Point {
	// Create new big.Int instances to avoid modifying the original values
	x1 := new(big.Int).Set(P.X)
	y1 := new(big.Int).Set(P.Y)
	x2 := new(big.Int).Set(Q.X)
	y2 := new(big.Int).Set(Q.Y)

	// If P == Q, return point doubling of P
	if x1.Cmp(x2) == 0 && y1.Cmp(y2) == 0 {
		return PointDoubling(P)
	}

	// Step 1: Calculate slope = (y2 - y1) / (x2 - x1) mod CURVE_P
	// Part 1: Calculate y2 - y1
	yDiff := new(big.Int).Sub(y2, y1)

	// Part 2: Calculate x2 - x1
	xDiff := new(big.Int).Sub(x2, x1)

	// Part 3: Calculate the modular inverse of x2 - x1
	xDiffInv := new(big.Int).ModInverse(xDiff, CURVE_P)

	// Part 4: Calculate the slope
	slope := new(big.Int).Mul(yDiff, xDiffInv)
	slope.Mod(slope, CURVE_P)

	// Step 2: Calculate x3 = slope^2 - x1 - x2 mod CURVE_P
	slopeSquared := new(big.Int).Mul(slope, slope) // slope^2
	x3 := new(big.Int).Sub(slopeSquared, x1)       // slope^2 - x1
	x3.Sub(x3, x2)                                 // slope^2 - x1 - x2
	x3.Mod(x3, CURVE_P)                            // mod CURVE_P

	// Step 3: Calculate y3 = slope * (x1 - x3) - y1 mod CURVE_P
	x1MinusX3 := new(big.Int).Sub(x1, x3)                // x1 - x3
	slopeTimesDiff := new(big.Int).Mul(slope, x1MinusX3) // slope * (x1 - x3)
	y3 := new(big.Int).Sub(slopeTimesDiff, y1)           // slope * (x1 - x3) - y1
	y3.Mod(y3, CURVE_P)                                  // mod CURVE_P

	// Return the new point
	return Point{X: x3, Y: y3}
}

func ScalarMultiplication(privateKey string) Point {
	privateKeyBigInt := new(big.Int)
	privateKeyBigInt.SetString(privateKey, 16)
	privateKeyBinary := privateKeyBigInt.Text(2)

	result := GPoint

	for i := 1; i < len(privateKeyBinary); i++ {
		result = PointDoubling(result)
		if privateKeyBinary[i] == '1' {
			result = PointAddition(result, GPoint)
		}
	}

	return result

}
