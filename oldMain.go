/* package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// add  the result to the  file
func AddToFile(data []byte, name string) {
	err := os.WriteFile(name, data, 0o644)
	if err != nil {
		fmt.Println(" Error ")
	}
	fmt.Println("file  has been created successfully")
}

// function to handle the ) in the exampels of (up, 2)
func MyAtoi(word string) int {
	correct := ""

	if word[len(word)-1] >= ')'    && word[len(word)-2] >= '1' && word[len(word)-2] <= '9' {
		for i := 0; i < len(word)-1; i++ {
			correct += string(word[i])
		}
	}
	num1, err := strconv.Atoi(correct)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	return num1
}

// The name explane isSelfe
func TableToString(table []string) string {
	text := ""
	for i := 0; i < len(table); i++ {
		if table[i] != "" {

			text += string(table[i])
			if i < len(table)-1 {
				text += " "
			}
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

// Process all functrions
func ProcessAll() {
	/*table := addToTable()
		table = Capitalized(table)
		table = Lower(table)
		table = Upper(table)
		table = BinaryToDecimal(table)
		table = HexadecimalToDecimal(table)
		table = withNumber(table)
		table = punctuations(table)
		table = Avowel(table)
		table = marks(table)
	fmt.Println(MarksWords(table))

		text := TableToString(table)
		data := []byte(text)
		AddToFile(data, os.Args[2])
	*/
/*
	table := addToTable()
	table = Capitalized(table)
	table = Lower(table)
	table = Upper(table)
	table = withNumber(table)
	table = punctuations(table)

	text := TableToString(table)
	data := []byte(text)
	AddToFile(data, os.Args[2])
}

func HexadecimalToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(hex)" {
				decimalNumber, err := strconv.ParseInt(table[i-1], 16, 0)
				if err != nil {
					fmt.Println("Error:", err)
					return table
				}
				table[i-1] = fmt.Sprintf("%d", decimalNumber) // Convert to string and replace
			}
		}
	}
	// ! if i have just (hex) i chould to delete the case or not
	result := DeleteCases("(hex)", table)

	return result
}

func BinaryToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(bin)" {
				decimalNumber, err := strconv.ParseInt(table[i-1], 2, 0)
				if err != nil {
					fmt.Println("Error:", err)
					return table
				}
				table[i-1] = fmt.Sprintf("%d", decimalNumber) // Convert to string and replace
			}
		}
	}
	// ! if i have just (bin) i chould to delete the case or not
	result := DeleteCases("(bin)", table)
	return result
}

func Capitalized(table []string) []string {
	for index, value := range table {
		if index > 0 && value == "(cap)" {
			word := table[index-1]
			// Capitalize the first letter
			if word != "" {
				runes := []rune(word)
				runes[0] = unicode.ToUpper(runes[0])
				for i := 1; i < len(runes); i++ {
					runes[i] = unicode.ToLower(runes[i])
				}
				table[index-1] = string(runes)
			}
		}
	}

	// Remove all instances of "(cap)" from the table
	result := DeleteCases("(cap)", table)
	return result
}

func Upper(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(up)" {
				if table[i-1] != "" {
					table[i-1] = strings.ToUpper(table[i-1])
				}
			}
		}
	}

	result := DeleteCases("(up)", table)
	return result
}

func Lower(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(low)" {
				if table[i-1] != "" {
					table[i-1] = strings.ToLower(table[i-1])
				}
			}
		}
	}

	result := DeleteCases("(low)", table)

	return result
}

// to the delete the casses like up, low ...
func DeleteCases(cas string, table []string) []string {
	var result []string
	for i := 0; i < len(table); i++ {
		if table[i] != cas {
			result = append(result, table[i])
		}
	}
	return result
}

// TODO: this is not working
func withNumber(table []string) []string {
	fmt.Println(table)
	/*
		!  ====> clean the code and make it easy
	*/
	//	var result []string

/*	count := 0
	for i := 0; i < len(table)-1; i++ {
		if i > 0 {
			count++
			num1 := 0
			if string(table[i]) == "(low,"  &&  strings.HasSuffix(table[i+1],")"){

				num1 = MyAtoi(table[i+1])
				if count-1 < num1 {
					num1 = count - 1
				}
				counter := 0

				for l := i - 1; l >= 0; l-- {
					if table[l] == "" {
						counter++
					} else {
						break
					}
				}
				for j := 1; j <= num1; j++ {

					table[i-j-counter] = strings.ToLower(table[i-j-counter])
					table[i] = ""
					if i < len(table) {
						table[i] = ""
					}
					if i+1 < len(table) {
						table[i+1] = ""
					}
					fmt.Println(table)
					fmt.Println(len(table))
				}
				fmt.Println(counter)
			}
			if string(table[i]) == "(up,"  &&  strings.HasSuffix(table[i+1],")") {
				num1 := 0
				num1 = MyAtoi(string(table[i+1]))
				if count-1 < num1 {
					num1 = count - 1
				}
				counter := 0

				for l := i - 1; l >= 0; l-- {
					if table[l] == "" {
						counter++
					} else {
						break
					}
				}
				for j := 1; j <= num1; j++ {

					table[i-j-counter] = strings.ToUpper(table[i-j-counter])
					table[i] = ""
					if i < len(table) {
						table[i] = ""
					}
					if i+1 < len(table) {
						table[i+1] = ""
					}
					fmt.Println(table)
				}
			}
			if string(table[i]) == "(cap,"  &&  strings.HasSuffix(table[i+1],")") {
				num1 = MyAtoi(string(table[i+1]))
				if count-1 < num1 {
					num1 = count - 1
				}

				counter := 0

				for l := i - 1; l >= 0; l-- {
					if table[l] == "" {
						counter++
					} else {
						break
					}
				}
				for j := 1; j <= num1; j++ {
					word := table[i-j-counter]

					runes := []rune(word)
					runes[0] = unicode.ToUpper(runes[0])
					for i := 1; i < len(runes); i++ {
						runes[i] = unicode.ToLower(runes[i])
					}
					table[i-j-counter] = string(runes)
					table[i] = ""
					if i < len(table) {
						table[i] = ""
					}
					if i+1 < len(table) {
						table[i+1] = ""
					}
				}
			}
		}
	}

	return table
}

func findPunctuationIndex(word string) int {
	punctuations := []string{",", ".", "!", "?", ":", ";"}
index:=0
	for _, punctuation := range punctuations {
		if strings.Contains(word, punctuation) {
			index = strings.Index(word, punctuation)
			break
		}
	}
	return index
}

func punctuations(table []string) []string {
	//  !  ,jkkle,fwef  this example is  not work
	var result []string
	corrWord := ""
	for i := 0; i < len(table); i++ {
		if string(table[i]) == "," {
			continue
		}

/* word:=""
	for j := 0; j < len(table[i]); j++ {
		word += string(table[i][j])
		if j >0 && j <len(table[i])-1 &&  table[i][j] == ',' || table[i][j] == '.' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';' {
		word += string(table[i][j])
		word += " "
		}
	} */
		

			/* if   i >0 && i <len(table)-1 &&  strings.Contains(table[i], ",") || strings.Contains(table[i], ".") || strings.Contains(table[i], "!") || strings.Contains(table[i], "?") || strings.Contains(table[i], ":") || strings.Contains(table[i], ";") {
				index := findPunctuationIndex(table[i])

				fmt.Println(index)
				for k := 0; k < len(table[i]); k++ {
					corrWord += string(table[i][k])
					if k == index {
						corrWord += " "
					}
				}
				table[i]  = strings.TrimSpace(corrWord) 
			} else */  /*if i > 0 && (table[i] == "," || table[i] == "." || table[i] == "!" || table[i] == "?" || table[i] == ":" || table[i] == ";") {

				table[i-1] = table[i-1] + table[i]
				table[i] = ""
			} else if i > 0 && ( strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";") ){
				count := 0
				for j := 0; j < len(table[i]); j++ {
					if table[i][j] == ',' || table[i][j] == '.' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';' {
						count++
					}
				}
				// fmt.Println(count)
				if count == 1 {
					table[i-1] = table[i-1] + string(table[i][0])
					table[i] = table[i][1:]

				} else {
					table[i-1] = table[i-1] + string(table[i][0:count])
					table[i] = table[i][count:]
				}

			}else if   ( strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";") )  {

				index := findPunctuationIndex(table[i])

				fmt.Println(index)
				for k := 0; k < len(table[i]); k++ {
					corrWord += string(table[i][k])
					if k == index {
						corrWord += " "
					}
				}
				table[i]  = strings.TrimSpace(corrWord) 

			}
			
		
	}
	for i := 0; i < len(table); i++ {
		if table[i] != "" {
			result = append(result, table[i])
		}
	}

	return result
}

func Avowel(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i < len(table)-1 {
			if (table[i] == "a" || table[i] == "A") && (strings.HasPrefix(table[i+1], "a") || strings.HasPrefix(table[i+1], "e") || strings.HasPrefix(table[i+1], "i") || strings.HasPrefix(table[i+1], "o") || strings.HasPrefix(table[i+1], "u") || strings.HasPrefix(table[i+1], "h") || strings.HasPrefix(table[i+1], "A") || strings.HasPrefix(table[i+1], "E") || strings.HasPrefix(table[i+1], "I") || strings.HasPrefix(table[i+1], "O") || strings.HasPrefix(table[i+1], "U") || strings.HasPrefix(table[i+1], "H")) {
				table[i] += "n"
			}
		}
	}
	return table
}

func marks(table []string) []string {
	// ! i hqve just one error in this example : (' multiple' ' wor'ds 'in quotes '.)
	for i := 0; i < len(table); i++ {
		if !(strings.HasPrefix(table[i], "'") && strings.HasSuffix(table[i], "'")) {
			if i+1 < len(table) && strings.HasPrefix(table[i], "'") && strings.HasPrefix(table[i+1], "'") {
				if len(table[i]) == 1 {
					table[i] = ""
				} else {
					table[i] = table[i] + "'"
				}
				table[i+1] = table[i+1][1:]
			} else if i+1 < len(table) && strings.HasSuffix(table[i], "'") && strings.HasSuffix(table[i+1], "'") {

				if len(table[i]) == 1 {
					table[i] = ""
				} else {
					table[i] = table[i][:len(table[i])-1]
				}
				table[i+1] = "'" + table[i+1]

			} else if i+2 < len(table) {
				if table[i] == "'" && table[i+2] == "'" {
					table[i] = ""
					table[i+2] = ""
					table[i+1] = "'" + table[i+1] + "'"
				} else if strings.HasSuffix(table[i], "'") && strings.HasPrefix(table[i+2], "'") { // hello' world 'rreda
					if len(table[i]) == 1 {
						table[i] = ""
					} else {
						table[i] = table[i][:len(table[i])-1]
					}
					table[i+2] = table[i+2][1:]
					table[i+1] = "'" + table[i+1] + "'"
				}
			}
		}
	}

	result := DeleteCases("", table)
	fmt.Println(result)
	return result
}

func MarksWords(table []string) []string {
	i := 0
	for i < len(table) {
		if strings.HasSuffix(table[i], "'") {
			// fmt.Println("not")

			startIndex := i
			for j := startIndex + 1; j < len(table); j++ {
				if strings.HasPrefix(table[j], "'") {
					endIndex := j
					// Process the words between startIndex and endIndex
					for k := startIndex; k <= endIndex; k++ {
						if k == startIndex {
							if len(table[k]) == 1 {
								table[k] = ""
							} else {
								table[k] = table[k][:len(table[k])-1]
							}
							table[k+1] = "'" + table[k+1]
						} else if k == endIndex {
							if len(table[k]) == 1 {
								table[k] = ""
							} else {
								table[k] = table[k][1:]
							}
							table[k-1] = table[k-1] + "'"
						}
					}
					i = endIndex // next position
					break
				} else if strings.HasSuffix(table[j], "'") {

					endIndex := j
					for k := startIndex; k <= endIndex; k++ {
						if k == startIndex {
							if len(table[k]) == 1 {
								table[k] = ""
							} else {
								table[k] = table[k][:len(table[k])-1]
							}
							table[k+1] = "'" + table[k+1]
						}
					}
					i = endIndex // Move to the next position after the endIndex
					break

				}
			}
		} else if strings.HasPrefix(table[i], "'") { // example of "hy 'I am the most well-known myself in the world 'how"
			startIndex := i

			for j := startIndex + 1; j < len(table); j++ {
				if strings.HasPrefix(table[j], "'") {
					endIndex := j
					for k := startIndex; k <= endIndex; k++ {
						if k == endIndex {
							if len(table[k]) == 1 {
								table[k] = ""
							} else {
								table[k] = table[k][1:]
							}
							table[k-1] = table[k-1] + "'"
						}
					}
					i = endIndex // Move to the next position after the endIndex
					break
				}
			}

		}
		i++
	}

	result := DeleteCases("", table)
	fmt.Println(result)
	return result
}

func addToTable() []string {
	name := os.Args[1]
	txt := ReadFile(name)

	// Split the content of the file into words
	word := ""
	var table []string
	for _, i := range txt {
		if i != ' ' {
			word += string(i)
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
	if len(os.Args) != 3 {
		fmt.Println("bro you should to write just the sample, result file names   ")
	} else {
		if !strings.HasSuffix(os.Args[2], ".txt") {
			fmt.Println(" Nice try ('_') ")
		} else {
			ProcessAll()
		}
	}
}
 