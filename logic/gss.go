package logic

import (
	"fmt"
	"log"
	"math"
	"math/big"
)

// var (
// 	sqrt5   = math.Sqrt(5)
// 	invphi  = (sqrt5 - 1) / 2 //# 1/phi
// 	invphi2 = (3 - sqrt5) / 2 //# 1/phi^2
// 	nan     = math.NaN()
// )

// var (
// 	bigSqrt5   = big.NewFloat(math.Sqrt(5))
// 	bigInvphi  = new(big.Float).SetPrec(256).Quo(bigSqrt5.Sub(bigSqrt5, big.NewFloat(1.0)), big.NewFloat(2.0))
// 	bigInvphi2 = new(big.Float).SetPrec(256).Quo(big.NewFloat(3.0).Sub(bigSqrt5, big.NewFloat(1.0)), big.NewFloat(2.0))
// 	bigNan     = new(big.Float).SetInf(false)
// )

var (
	sqrt5   = math.Sqrt(5)
	invphi  = (sqrt5 - 1) / 2 //# 1/phi
	invphi2 = (3 - sqrt5) / 2 //# 1/phi^2
	nan     = math.NaN()
)

var (
	bigSqrt5   = big.NewFloat(math.Sqrt(5))
	bigInvphi  = new(big.Float).SetPrec(256).Quo(new(big.Float).Sub(bigSqrt5, big.NewFloat(1.0)), big.NewFloat(2.0))
	bigInvphi2 = new(big.Float).SetPrec(256).Quo(new(big.Float).Sub(big.NewFloat(3.0), bigSqrt5), big.NewFloat(2.0))
	bigNan     = new(big.Float).SetInf(false)
)

// Gss golden section search (recursive version)
// https://en.wikipedia.org/wiki/Golden-section_search
// '''
// Golden section search, recursive.
// Given a function f with a single local minimum in
// the interval [a,b], gss returns a subset interval
// [c,d] that contains the minimum with d-c <= tol.
//
// logger may be nil
//
// example:
// >>> f = lambda x: (x-2)**2
// >>> a = 1
// >>> b = 5
// >>> tol = 1e-5
// >>> (c,d) = gssrec(f, a, b, tol)
// >>> print (c,d)
// (1.9999959837979107, 2.0000050911830893)
// '''
func BigGss(f func(*big.Float) *big.Float, a, b, tol *big.Float, logger *log.Logger) (*big.Float, *big.Float) {

	fmt.Println("bigSqrt5", bigSqrt5.String())
	fmt.Println("bigInvphi", bigInvphi.String())
	fmt.Println("bigInvphi2", bigInvphi2.String())
	fmt.Println("bigNan", bigNan.String())
	fmt.Println("sqrt5", sqrt5)
	fmt.Println("invphi", invphi)
	fmt.Println("invphi2", invphi2)
	fmt.Println("nan", nan)

	return bigGss(f, a, b, tol, bigNan, bigNan, bigNan, bigNan, bigNan, logger)
}
func bigGss(f func(*big.Float) *big.Float, a, b, tol, h, c, d, fc, fd *big.Float, logger *log.Logger) (*big.Float, *big.Float) {
	if bigGreaterThan(a, b) { // a > b
		a, b = b, a
	}
	h = new(big.Float).Sub(b, a) //h = b - a
	it := 0
	for {
		if logger != nil {
			logger.Printf("%d\t%9.6g\t%9.6g\n", it, a, b)
		}
		it++
		if bigLessThan(h, tol) { //h < tol
			return a, b
		}
		if bigGreaterThan(a, b) { //a > b
			a, b = b, a
		}
		if c.IsInf() { //math.IsNaN(c)
			c = new(big.Float).Add(a, new(big.Float).Mul(bigInvphi2, h)) //c = a + invphi2*h
			fc = f(c)
		}
		if d.IsInf() { //math.IsNaN(d)
			d = new(big.Float).Add(a, new(big.Float).Mul(bigInvphi, h)) //d = a + invphi*h
			fd = f(d)
		}
		if bigLessThan(fc, fd) { //fc < fd
			b, h, c, fc, d, fd = d, new(big.Float).Mul(h, bigInvphi), bigNan, bigNan, c, fc //b, h, d, fd, c, fc = d, h*invphi, nan, nan, c, fc
		} else {
			a, h, c, fc, d, fd = c, new(big.Float).Mul(h, bigInvphi), d, fd, bigNan, bigNan //a, h, c, fc, d, fd = c, h*invphi, d, fd, nan, nan
		}
	}
}

func bigLessThan(a, b *big.Float) bool {
	// return a < b
	return a.Cmp(b) < 0
}

func bigGreaterThan(a, b *big.Float) bool {
	// return a > b
	return a.Cmp(b) > 0
}

func Gss(f func(float64) float64, a, b, tol float64, logger *log.Logger) (float64, float64) {
	return gss(f, a, b, tol, nan, nan, nan, nan, nan, logger)
}
func gss(f func(float64) float64, a, b, tol, h, c, d, fc, fd float64, logger *log.Logger) (float64, float64) {
	if a > b {
		a, b = b, a
	}
	h = b - a
	it := 0
	for {
		if logger != nil {
			logger.Printf("%d\t%9.6g\t%9.6g\n", it, a, b)
		}
		it++
		if h < tol {
			return a, b
		}
		if a > b {
			a, b = b, a
		}
		if math.IsNaN(c) {
			c = a + invphi2*h
			fc = f(c)
		}
		if math.IsNaN(d) {
			d = a + invphi*h
			fd = f(d)
		}
		if fc < fd {
			b, h, c, fc, d, fd = d, h*invphi, nan, nan, c, fc
		} else {
			a, h, c, fc, d, fd = c, h*invphi, d, fd, nan, nan
		}
	}
}
