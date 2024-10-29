package main

import (
	"fmt"
	"time"
)

const (
	easyPassword   = "P4ssw0rd007"
	mediumPassword = "P4$$ssw0rd^18AZ!*192"
	hardPassword   = "x9@RK`*LGuHMrvd5gVE?;%"
)

func main() {
	// easyPwdAvgHashingTime := 0
	// easyPwdAvgMatchingTime := 0

	// mediumPwdAvgHashingTime := 0
	// mediumPwdAvgMatchingTime := 0

	// hardPwdAvgHashingTime := 0
	// hardPwdAvgMatchingTime := 0

	fmt.Println("Easy password")
	for i := 1; i <= 15; i++ {
		measureTime(i, easyPassword)
		// hash, match := measureTime(i, easyPassword)
		// if hash != 0 && match != 0 {
		// 	easyPwdAvgHashingTime += int(hash)
		// 	easyPwdAvgMatchingTime += int(match)
		// }
	}

	fmt.Println("Medium password")
	for i := 1; i <= 15; i++ {
		measureTime(i, mediumPassword)
		// hash, match := measureTime(i, mediumPassword)
		// if hash != 0 && match != 0 {
		// 	mediumPwdAvgHashingTime += int(hash)
		// 	mediumPwdAvgMatchingTime += int(match)
		// }
	}

	fmt.Println("Hard password")
	for i := 1; i <= 15; i++ {
		measureTime(i, hardPassword)
		// hash, match := measureTime(i, hardPassword)
		// if hash != 0 && match != 0 {
		// 	hardPwdAvgHashingTime += int(hash)
		// 	hardPwdAvgMatchingTime += int(match)
		// }
	}
}

func measureTime(iterations int, password string) (time.Duration, time.Duration) {
	_iterations = uint32(iterations)
	fmt.Println("Measuring time for", iterations, "iterations")

	startTime := time.Now()
	hashedPassword, err := HashPassword(password)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	endTime := time.Now()
	diff1 := endTime.Sub(startTime)
	fmt.Println("Hashed password:", hashedPassword, "Time taken:", diff1)

	startTime = time.Now()
	match := ComparePasswordMatching(password, hashedPassword)
	endTime = time.Now()
	diff2 := endTime.Sub(startTime)
	fmt.Println("Password match:", match, "Time taken:", diff2)

	fmt.Println()
	return diff1, diff2
}
