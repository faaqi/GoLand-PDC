package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//text := "1"

	fmt.Println("Check Corona updates and suggestions here\n")

	fmt.Println("1. Print Covid 19 cases in Pakistan")

	fmt.Println("2. Print Covid 19 cases in South Korea")

	fmt.Println("3. Print Covid 19 cases in France")

	fmt.Println("4. Print a Personalized message for Covid-19")

	fmt.Print("Select an option from Menu (1-4) , otherwise Enter 0 for exiting the menu : ")

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		if input.Text() == "0" {
			break
		} else if input.Text() == "1" {

			fmt.Println("\n Covid-19 cases in pakistan are 1,526 \n")
		} else if input.Text() == "2" {

			fmt.Println("\n Covid-19 cases in South Korea are 9,583 \n ")

		} else if input.Text() == "3" {

			fmt.Println("\nCovid-19 cases in France are 40,174\n")

		} else if input.Text() == "4" {

			fmt.Println("\nCorona please shaant hojao. Go Corona Go. Corona jaa jaa Duniya c nikal jaa... Corona, time nikaal k marjaa ")

		}

		fmt.Println("\nCheck Corona updates and suggestions here\n")

		fmt.Println("1. Print Covid 19 cases in Pakistan")

		fmt.Println("2. Print Covid 19 cases in South Korea")

		fmt.Println("3. Print Covid 19 cases in France")

		fmt.Println("4. Print a Personalized message for Covid-19\n")

		fmt.Print("\nSelect an option from Menu (1-4) , otherwise Enter 0 for exiting the menu : ")

	}

}
