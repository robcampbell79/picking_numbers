package main

import(
    "fmt"
)

func main() {
	// a := []int32{5,4,6,5,3,3,1,2,3,9,3,4,5,5,6}
	a := []int32{1,2,3,4,5,6,7}

	fmt.Println(pickingNumbers(a))
}

func pickingNumbers(a []int32) int32 {
	matches := repeatNumbers(a)
	
	var lrg int32 = findLrg(matches) + 1
	var add int32 = 0
	
	for i := 0; i < len(matches); i++ {
		for j := i+1; j < len(matches); j++ {
			newI := fmt.Sprintf("%d", i)
			newIPl := fmt.Sprintf("%d", j)
			sub := matches["p"+newI][0] - matches["p"+newIPl][0]
			if sub < 0  {
				sub = sub * -1
			}

			if sub <= 1 {
				add = matches["p"+newI][1] + matches["p"+newIPl][1]
			}

			if add > lrg {
				lrg = add
			}
		}
	}



	fmt.Println("lrg: ", lrg)

	fmt.Println(matches)

	return 5
}

func repeatNumbers(a []int32) map[string][]int32 {
	l := len(a)
	var count int32
	var keyCount int32 = 0
	var repeater int32
	repeatMap := make(map[string][]int32)
	flag := false

	for i := 0; i < l; i++ {
		repeatArr := []int32{}
		for _, x := range repeatMap {
			if x[0] == a[i] {
				flag = true
				break
			} else {
				flag = false
			}
		}
		for j := 0; j < l; j++ {
			if flag == false {
				repeater = a[i]
				if a[j] == repeater {
					count++
				}
			} else {
				break
			}
		}

		if count > 1 && flag == false {
			repeatArr = append(repeatArr, repeater)
			repeatArr = append(repeatArr, count)
			k := fmt.Sprintf("%d", keyCount)
			repeatMap["p" + k] = repeatArr
			keyCount++
		} 

		repeatArr = nil
		count = 0
		repeater = 0
	}

	return repeatMap
}

func findLrg(m map[string][]int32) int32 {
	var lrg int32 = 0

	for i := 0; i < len(m); i++ {
		newI := fmt.Sprintf("%d", i)
		if m["p"+newI][1] > lrg {
			lrg = m["p"+newI][1]
		}
	}

	return lrg
}