package main

import "fmt"

func main() {
	s := "III"
	res := romanToInt(s)
	if res != 3 {
		fmt.Println("res1 failed:", s, res)
	}

	s = "LVIII"
	res = romanToInt(s)
	if res != 58 {
		fmt.Println("res2 failed:", s, res)
	}

	s = "MCMXCIV"
	res = romanToInt(s)
	if res != 1994 {
		fmt.Println("res3 failed:", s, res)
	}
}

func romanToInt(s string) int {

}
