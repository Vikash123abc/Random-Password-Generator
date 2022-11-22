package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const ss = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()"

var (
	minSpecialCharacterSet = 1
	minNumberSet           = 1
	minUpperCaseSet        = 1
	minLowerCaseSet        = 1
	specialCharcterSet     = "~!@#$%^&*()_"
	numberSet              = "0123456789"
	lowerCaseSet           = "abcdefghijklmnopqrstuvwxyz"
	upperCaseSet           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandStringBytes(n int, s string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}
	return string(b)
}

func passwordGenerator(size int) string {
	var password = ""
	// Special Characters
	s := RandStringBytes(minSpecialCharacterSet, specialCharcterSet)
	password += s
	// Number
	s = RandStringBytes(minNumberSet, numberSet)
	password += s
	// UpperCase
	s = RandStringBytes(minUpperCaseSet, upperCaseSet)
	password += s
	// LowerCase
	s = RandStringBytes(minLowerCaseSet, lowerCaseSet)
	password += s
	// General scenerio
	s = RandStringBytes(size, ss)
	password += s
	return password
}

func main() {
	var s, n string
	fmt.Println("Please provide the count of passwords needed")
	fmt.Scanln(&n)
	nn, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("Please provide correct number")
		return
	}

	fmt.Println("Please provide the size of the password (between 5 to 10) ")
	fmt.Scanln(&s)

	Size, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Please provide correct number")
		return
	}

	if Size < 5 || Size > 10 {
		fmt.Printf("Please provide the size between 5 and 10")
		return
	}

	size := Size - minSpecialCharacterSet - minNumberSet - minUpperCaseSet - minLowerCaseSet
	rand.Seed(time.Now().Unix())
	fmt.Println("The unique password are below:")
	for i := 0; i < nn; i++ {
		password := passwordGenerator(size)
		fmt.Println(password)
	}
}
