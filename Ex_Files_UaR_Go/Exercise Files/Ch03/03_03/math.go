package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	// invalid operation
	var int1 = 5
	var float1 float64 = 43
	// sum := int1 + float1 ( will throw "mismatched types int and float64" error )
	sum := float64(int1) + float1
	fmt.Printf("Sum: %v, Type:%T\n", sum, sum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.16g\n", &bigSum)

	// For floating-point values, width sets the minimum width of the field
	// and precision sets the number of places after the decimal, if appropriate,
	// except that for %g/%G precision sets the total number of significant digits.
	// For example, given 12.345 the format %6.3f prints 12.345 while %.3g prints 12.3.
	// The default precision for %e, %f and %#g is 6; for %g it is the smallest number
	// of digits necessary to identify the value uniquely.
	fmt.Printf("%.4g\n", 12.323)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
