package euler

import (
	// "fmt"
	"math"
)

/**
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
func problemOne() int {
	sol := 0
	for x := 0; x < 1000; x++{
		if x % 3 == 0 || x % 5 == 0 {
			sol += x
		}
	}
	return sol
}

/**
Each new term in the Fibonacci sequence is generated by adding the previous two terms. 
By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func problemTwo() int {
	sum := 0;
	prev := 1
	curr := 2
	for curr < 4000000{
		curr += prev // 3
		prev = curr - prev //
		if (prev % 2 == 0) {
			sum += prev
		}
	}
	return sum
}

/**
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
func problemThree() int {
	n := 600851475143
	// Solve using...divide & conquer? IDK its just division
	// factor := 3
	// lastFactor := 1
	// for n > 1 {
	// 	if n % factor == 0 {
	// 		n = n / factor
	// 		lastFactor = factor
	// 		for n % factor == 0 {
	// 			n = n / factor
	// 		}
	// 	}
	// 	factor +=2
	// }

	// Solve using Fermat's factorization
	lastFactor := int(fermatFactorRecursive(float64(n)))
	// fmt.Println(lastFactor)

	return lastFactor
}
func fermatFactorRecursive(n float64) float64 {
	x, y := fermatFactor(n)
	if x == 1 {
		// y is prime
		return y
	}
	return math.Max(fermatFactorRecursive(x), fermatFactorRecursive(y))
}

// Return (a-b) (a+b). a-b is the smallest factor.
func fermatFactor(n float64) (float64, float64) {
	a := math.Ceil(math.Sqrt(n))
	b2 := a*a - n
	for math.Sqrt(b2) != math.Floor(math.Sqrt(b2)) {
		a = a+1
		b2 = a*a - n
	}
	return a - math.Sqrt(b2), a + math.Sqrt(b2)
}

/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
func problemFour() int {
	return 1
}

/**
Problem 5: Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func problemFive() int {
	return 1
}

/**
Problem 6: Sum square difference
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func problemSix() int {
	return 1
}

/**
Problem 7: 10001st prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
func problemSeven() int {
	return 1
}

/**
Problem 8: Largest product in a series
The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450

Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
*/
func problemEight() int {
	return 1
}

/**
Problem 9: Special Pythagorean triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
func problemNine() int {
	return 1
}

/**
Problem 10: Summation of primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
func problemTen() int {
	return 1
}