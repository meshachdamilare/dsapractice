package containsduplicate

import (
	"math/rand"
	"time"
)

//func main() {
//	fmt.Println(ContainsDuplicates([]int{1, 2, 3, 4, 4, 5, 6}))
//}

// ContainsDuplicates Contains Duplicate
func ContainsDuplicates(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func RandomizedArrayB(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func RandomizedArrayA(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
