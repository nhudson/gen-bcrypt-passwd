package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	costChart = `
This chart will help you calculate the cost and interations of
generating the hash.

|------|-------------------|
| Cost | Iterations        |
|------|-------------------|
|  2   |    4 iterations   |
|  4   |    16 iterations  |
|  5   |    32 iterations  |
|  6   |    64 iterations  |
|  7   |    128 iterations |
|  8   |    256 iterations |
|  9   |    512 iterations |
| 10   |  1,024 iterations |
| 11   |  2,048 iterations |
| 12   |  4,096 iterations |
| 13   |  8,192 iterations |
| 14   | 16,384 iterations |
| 15   | 32,768 iterations |
| 16   | 65,536 iterations |
|------|-------------------|`
)

var (
	password string
	cost     int
)

// Set Password from command line
func getPasswordFromCmdLine() ([]byte, string) {
	// Lets promt for user input for password
	fmt.Println("Enter password to hash:")

	// Read the input
	_, err := fmt.Scanln(&password)
	if err != nil {
		log.Println(err)
	}

	return []byte(password), password
}

// Set cost from command line
func getCostFromCmdLine() int {
	// Prompt to enter cost
	fmt.Println(costChart)
	fmt.Println("Enter cost:")

	// Read input of cost
	_, err := fmt.Scanln(&cost)
	if err != nil {
		log.Println(err)
	}
	return int(cost)
}

// Hash password
func hashPassword(password []byte, cost int) ([]byte, error) {
	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(password, cost)

	return hashedPasswordBytes, err
}

func doPasswordsMatch(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password))
	return err == nil
}

func hashingCost(password []byte) (int, error) {
	costInt, err := bcrypt.Cost(password)

	return int(costInt), err
}

func main() {
	// Input password
	passwordBytes, password := getPasswordFromCmdLine()

	// Input Cost
	setCost := getCostFromCmdLine()

	// Hash password
	hashedPassword, err := hashPassword(passwordBytes, setCost)
	if err != nil {
		return
	}

	fmt.Println("Password Hash:", string(hashedPassword))

	// Check if passed password matches the original password
	fmt.Println("Password Match:",
		doPasswordsMatch(string(hashedPassword), password))
}
