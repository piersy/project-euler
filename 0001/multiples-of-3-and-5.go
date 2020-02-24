package main

func main() {
	threes := 0
	fives := 0
	sum := 0
	complete := 0
	for {
		if threes <= fives && threes < 1000{
			threes += 3
			if threes >= 1000 {
				complete ++
				if complete == 2 {
					break
				}
				continue
			}
			if threes != fives{
				sum += threes
			}
		}else if fives < 1000{
			fives += 5
			if fives >= 1000 {
				complete ++
				if complete == 2 {
					break
				}
				continue
			}
			if threes != fives{
				sum += fives
			}
		}
	}
	println(sum)
}
