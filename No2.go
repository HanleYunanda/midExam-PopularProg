package main

import "fmt"

func compareArrays(arr1, arr2 [100]string) [100]bool {
	var arrCheck [100]bool
	for i := 0; i < 100; i++ {
		if arr1[i] != arr2[i] {
			arrCheck[i] = true
		}
	}

	return arrCheck
}

func main() {
	var arr1 [100]string
	var arr2 [100]string
	i := 0
	j := 0

	fmt.Println("Array 1 :")
	for {
		fmt.Scan(&arr1[i])
		if arr1[i] == "." {
			break
		}
		i++
	}

	fmt.Println("Array 2 :")
	for {
		fmt.Scan(&arr2[j])
		if arr2[j] == "." {
			break
		}
		j++
	}

	arrCheck := compareArrays(arr1, arr2)
	if i > j {
		for x := 0; x < i; x++ {
			if arrCheck[x] == true {
				fmt.Println("index ke-", x, "berbeda")
			}
		}
	} else {
		for x := 0; x < j; x++ {
			if arrCheck[x] == true {
				fmt.Println("index ke-", x, "berbeda")
			}
		}
	}
}

// go run compareArrays.go
