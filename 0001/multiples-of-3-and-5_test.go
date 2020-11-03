/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []func() int{v1, v2, v3, v4}

func TestVersions(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf(FunctionName(tc)), func(t *testing.T) {
			assert.Equal(t, 233168, tc())
		})
	}
}

func BenchmarkVersions(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf(FunctionName(tc)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				assert.Equal(b, 233168, tc())
			}
		})
	}
}

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Simple implemenation
func v1() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}

// Keep a running total of threes and fives and keep adding the respective
// values to them and then adding each to the total when that happens. But
// ensuring that we don't double add in the case that both threes and fives
// have the same value.
func v2() int {
	threes := 0
	fives := 0
	sum := 0
	complete := 0
	for {
		if threes <= fives && threes < 1000 {
			threes += 3
			if threes >= 1000 {
				complete++
				if complete == 2 {
					break
				}
				continue
			}
			if threes != fives {
				sum += threes
			}
		} else if fives < 1000 {
			fives += 5
			if fives >= 1000 {
				complete++
				if complete == 2 {
					break
				}
				continue
			}
			if threes != fives {
				sum += fives
			}
		}
	}
	return sum
}

// Here I try to write v2 in a more concise way, id didn't work. It's also
// slower, I guess because of the slice indexing.
func v3() int {
	increments := []int{3, 5}
	totals := []int{0, 0}
	length := len(totals)
	sum := 0
	curr := 0
	for {
		// Swap if other total is less, if we have length 1 nothing changes
		if totals[curr] > totals[(curr+1)%length] {
			curr = (curr + 1) % length
		}

		// Add increment
		totals[curr] += increments[curr]

		// Remove when exceeding limit
		if totals[curr] >= 1000 {
			if len(totals) == 1 {
				goto END
			}
			// Delete at curr
			totals = append(totals[:curr], totals[curr+1:]...)
			increments = append(increments[:curr], increments[curr+1:]...)
			length -= 1
			curr = (curr + 1) % length
			continue
		}

		// If both values are the same that means that the lesser value "caught
		// up" so don't add it because we already added it.
		if len(totals) == 2 && totals[curr] == totals[(curr+1)%length] {
			continue
		}
		// Add to the sum
		sum += totals[curr]
	}
END:
	return sum
}

// A second attempt at simplifying v2
func v4() int {
	threes := 0
	fives := 0
	sum := 0
	for threes < 1000 || fives < 1000 {
		var result int
		if threes < fives {
			threes += 3
			if threes < 1000 {
				result = threes
			}

		} else {
			fives += 5
			if fives < 1000 {
				result = fives
			}
		}
		// If they are equal then we must have already summed this value.
		if threes == fives {
			result = 0
		}
		sum += result
	}
	return sum
}

// Now having looked at the forum, I will attempt the neat solution using
// arithmetic progressions. Basically we can use the formula for summing
// arithmetic progressions to sum the progression of 3s and the progression of
// 5s and subtract the progression of 15s.
//
// The formula is n/2(2a + (n – 1)d). 'a' is the first term and 'd' is the
// difference between the first and second term
//
// So for 3s there are 333 terms in the sequence (3*333 == 999) and for 5s
// there are 199 (5*199 == 995) and for 15s there are 66 (66*15 == 990)
//
// So this gives us:
// 333/2(6 + 332*3) = 166833
// 199/2(10 + 198*5) = 99500
// 66/2(30 + 65*15) = 33165
// 166833 + 99500 - 33165 = 233168
