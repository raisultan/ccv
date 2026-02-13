package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func main() {
	fmt.Println("Welcome to the Interstellar Credit Card Validator")

	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Println("Error loading bank data:", err)
		return
	}

	for {
		cardNumber := getUserInput()
		if cardNumber == "" {
			fmt.Println("Exiting the program. Thank you for using the Interstellar Credit Card Validator!")
			break
		}

		if !validateInput(cardNumber) {
			fmt.Println("Invalid input. Please enter a valid credit card number.")
			continue
		}

		isValid := validateLuhn(cardNumber)
		if !isValid {
			fmt.Println("The credit card number is invalid.")
			continue
		}

		bin := extractBIN(cardNumber)
		bank := identifyBank(bin, banks)

		fmt.Printf("The credit card number is valid.\n")
		if bank != "" {
			fmt.Printf("Issuing Bank: %s\n", bank)
		} else {
			fmt.Println("Unable to identify the issuing bank.")
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a credit card number (or press Enter to quit): ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func validateInput(input string) bool {
	if len(input) < 13 || len(input) > 19 {
		return false
	}
	for _, ch := range input {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func validateLuhn(cardNumber string) bool {
	sum := 0
	isEven := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if isEven {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isEven = !isEven
	}

	return sum%10 == 0
}

func extractBIN(cardNumber string) int {
	bin, _ := strconv.Atoi(cardNumber[:6])
	return bin
}

func loadBankData(filename string) ([]Bank, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks []Bank
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		binFrom, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		binTo, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}
		banks = append(banks, Bank{Name: parts[0], BinFrom: binFrom, BinTo: binTo})
	}

	return banks, nil
}

func identifyBank(bin int, banks []Bank) string {
	for _, bank := range banks {
		if bin >= bank.BinFrom && bin <= bank.BinTo {
			return bank.Name
		}
	}
	return ""
}
