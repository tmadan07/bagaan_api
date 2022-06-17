package global

import (
	"fmt"
	"math"
	"strconv"
)

// basic linear congruential generator
func lcg(a, c, m, seed uint64) func() uint64 {
	r := seed
	return func() uint64 {
		r = (a*r + c) % m
		return r
	}
}

//compute luhn value
func Luhn(pan15_string string) int {

	//fullpart := strconv.Itoa(pan_segment_format) + strconv.Itoa(pan_segment)
	//pan15_string := strconv.Itoa(pan_15digit)
	multiplierLength := len(pan15_string)
	fmt.Println("luhn...15=", pan15_string)
	totalSum := 0
	if multiplierLength%2 == 0 {
		for i := 1; i <= multiplierLength; i++ {
			t, _ := strconv.Atoi(pan15_string[i-1 : i])
			if i%2 == 0 {
				t *= 2
				if t/10 >= 1 {
					quo := t % 10
					rem := t / 10
					sum := quo + rem
					t = sum
				}
				totalSum += t
				//fmt.Println("2x", i)
			} else {
				t *= 1
				if t/10 >= 1 {
					quo := t % 10
					rem := t / 10
					sum := quo + rem
					t = sum
				}
				totalSum += t
				//fmt.Println("1x", i)
			}
		}
	} else {
		for i := 1; i <= multiplierLength; i++ {
			t, _ := strconv.Atoi(pan15_string[i-1 : i])
			if i%2 == 0 {
				t *= 1
				if t/10 >= 1 {
					quo := t % 10
					rem := t / 10
					sum := quo + rem
					t = sum
				}
				totalSum += t
				//fmt.Println("1x", t)
			} else {
				t *= 2
				if t/10 >= 1 {
					quo := t % 10
					rem := t / 10
					sum := quo + rem
					t = sum
				}
				totalSum += t
				//fmt.Println("2x", t)
			}
		}
	}
	//fmt.Println("sum = ", totalSum)
	modValue := 10 - (totalSum%10)%10
	if modValue == 10 {
		modValue = 0
	}
	fmt.Println("luhn = 16=", modValue)
	return modValue
}

func PanGenerator(seed uint64, digits int) uint64 {
	g := lcg(2042041, 510511, uint64(math.Pow10(digits)), seed)
	return g()
}

// microsoft generator has extra division step
func msg(seed uint64) func() uint64 {
	g := lcg(2042041, 510511, 1<<31, seed)
	return func() uint64 {
		return g() / (1 << 16)
	}
}

func example(seed uint64) {
	fmt.Printf("\nWith seed = %d\n", seed)
	bsd := lcg(2042041, 510511, 1<<31, seed)
	msf := msg(seed)
	fmt.Println("       BSD  Microsoft")
	for i := 0; i < 5; i++ {
		fmt.Printf("%10d    %5d\n", bsd(), msf())
	}
}

// func main() {
// 	example(0)
// 	example(1)
// }
