package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// add  the result to the a file
func AddToFile(data []byte, name string) {
	err := os.WriteFile(name, data, 0o644)
	if err != nil {
		fmt.Println(" Error ")
	}
	fmt.Println("success")
}

// The name explanation isSelfe
func TableToString(table []string) string {
	text := ""
	for i := 0; i < len(table); i++ {
		text += string(table[i])
		if i < len(table)-1 { // Fix the condition to avoid adding extra space at the end
			text += " "
		}
	}
	return text
}

// Read the file
func ReadFile(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("There is no file with this name")
	}
	return string(data)
}

// Process all things
func ProcessAll() {
	table := addToTable()
	table = Capitalized(table)
	table = Upper(table)
	table = LowerWNbr(table)

	table = Lower(table)
	table = BinaryToDecimal(table)
	table = HexadecimalToDecimal(table)
	text := TableToString(table)
	data := []byte(text)
	AddToFile(data, "text.txt")
}

func HexadecimalToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(hex)" {
			decimalNumber, err := strconv.ParseInt(table[i-1], 16, 0)
			if err != nil {
				fmt.Println("Error:", err)
				return table
			}
			table[i-1] = fmt.Sprintf("%d", decimalNumber) // Convert to string and replace
		}
	}
	return table
}

func BinaryToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(bin)" {
			decimalNumber, err := strconv.ParseInt(table[i-1], 2, 0)
			if err != nil {
				fmt.Println("Error:", err)
				return table
			}
			table[i-1] = fmt.Sprintf("%d", decimalNumber) // Convert to string and replace
		}
	}
	return table
}

func Capitalized(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(cap)" {
			word := table[i-1]
			// Capitalize the first letter
			table[i-1] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return table
}

func Upper(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(up)" {
			table[i-1] = strings.ToUpper(table[i-1])
		}
	}
	return table
}

func Lower(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(low)" {
			table[i-1] = strings.ToLower(table[i-1])
		}
	}
	return table
}

func LowerWNbr(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i][0] == '(' {
			//	table[i-1] = strings.ToLower(table[i-1])
			fmt.Println(table[i] + "<=")
			fmt.Println(string(table[i+1][0]) + "<=")

			num1, err := strconv.Atoi(string(table[i+1][0]))
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}

			if string(table[i]) == "(low," {
				for j := 0; j < num1; j++ {
					table[i-j] = strings.ToLower(table[i-j])
				}
			}

		}
	}
	return table
}

func addToTable() []string {
	name := os.Args[1]
	txt := ReadFile(name)

	// Split the content of the file into words
	word := ""
	var table []string
	for i := 0; i < len(txt); i++ {
		if txt[i] != ' ' {
			word += string(txt[i])
		} else {
			if word != "" {
				table = append(table, word)
			}
			word = ""
		}
	}
	if word != "" {
		table = append(table, word)
	}
	return table
}

func main() {
	ProcessAll()
}
