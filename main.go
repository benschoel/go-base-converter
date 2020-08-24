package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convert(num, originalBase, newBase float64) (string, error) {
	n := 0

	if int(originalBase) <= 1 || int(newBase) <= 1 || int(originalBase) > 36 || int(newBase) > 36 {
		return "", fmt.Errorf("Invalid base")
	}

	numAsStr := strconv.Itoa(int(num))

	chars := []rune(numAsStr)

	for i := 0; i < len(chars); i++ {
		r, _ := strconv.Atoi(string(chars[i]))

		if r > int(originalBase)-1 {
			return "", fmt.Errorf("%v is not of base %v", num, originalBase)
		}
	}

	for math.Pow(newBase, float64(n)) <= num {
		n++
	}

	converted := ""
	remainder := num
	for n > 0 {
		n--
		j := math.Pow(newBase, float64(n))

		k := 0
		for j*float64(k) <= remainder {
			k++
		}
		k--
		remainder -= j * float64(k)

		if k > 9 {
			converted = converted + fmt.Sprintf("%c", rune(k+55))
		} else {
			converted = converted + strconv.Itoa(k)

		}

	}

	return converted, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nA Simple Base Converter by Benjamin Schoelkopf")
	fmt.Println()
	fmt.Println("Enter the number you'd like to convert")
	fmt.Print(": ")

	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSuffix(numStr, "\n")
	num, _ := strconv.ParseFloat(numStr, 64)

	fmt.Println("\nEnter the base of that number")
	fmt.Print(": ")
	base1Str, _ := reader.ReadString('\n')
	base1Str = strings.TrimSuffix(base1Str, "\n")
	base1, _ := strconv.ParseFloat(base1Str, 64)

	fmt.Println("\nEnter the base you'd like to convert to")
	fmt.Print(": ")

	base2Str, _ := reader.ReadString('\n')
	base2Str = strings.TrimSuffix(base2Str, "\n")
	base2, _ := strconv.ParseFloat(strings.TrimSpace(base2Str), 64)

	fmt.Println()

	converted, err := convert(num, base1, base2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v(%v) in base %v is %v\n", numStr, base1Str, base2Str, converted)
	// fmt.Println("", numStr, "(", base1Str, ") in base ", base2str, " is ", convert(num, base1, base2))
}
