package main

import "testing"

func TestV1(t *testing.T) {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	println(sum)
}

func TestV2(t *testing.T) {
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
	println(sum)
}

func TestV3(t *testing.T) {
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
	println(sum)
}
