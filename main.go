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

func MyAtoi(word string) int {
	correct := ""
	if word[len(word)-1] >= ')' && word[len(word)-2] >= '1' && word[len(word)-2] <= '9' {
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

/*
func cap(word string) string {

for i := 0; i < len(word); i++ {
	if (word[i]  >= 'A'   &&   word[i]  <= 'Z' )||  (word[i]  >= 'a'   &&   word[i]  <= 'z' ){
		 string(word[i]) =
	}
}


}
*/
// The name explanation isSelfe
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

// Process all things
func ProcessAll() {
	table := addToTable()               // قراءة الكلمات من الملف
	table = Capitalized(table)          // تطبيق التعديلات
	table = Upper(table)                // تطبيق (up)
	table = Lower(table)                // تطبيق (low)
	table = BinaryToDecimal(table)      // تحويل (bin)
	table = HexadecimalToDecimal(table) // تحويل (hex)
	table = withNumber(table)           // معالجة التعديلات بالأرقام
	table = punctuations(table)
	table = Avowel(table)
	//table = marks(table)
fmt.Println(MarksWords(table))  

	text := TableToString(table)
	// تحويل الجدول إلى نص
	data := []byte(text) // كتابة النص إلى الملف
	AddToFile(data, os.Args[2])
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
	result := DeleteCases("(hex)", table)

	return result
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

	result := DeleteCases("(bin)", table)

	return result
}

func Capitalized(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(cap)" {
			word := table[i-1]
			// Capitalize the first letter
			if word[0] >= 'a' && word[0] <= 'z' {
				table[i-1] = strings.ToUpper(string(word[0])) + word[1:]
			} else {
				table[i-1] = word
			}
		}
	}

	result := DeleteCases("(cap)", table)
	return result
}

func Upper(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(up)" {
			table[i-1] = strings.ToUpper(table[i-1])
		}
	}

	result := DeleteCases("(up)", table)
	return result
}

func Lower(table []string) []string {
	for i := 0; i < len(table); i++ {
		if table[i] == "(low)" {
			table[i-1] = strings.ToLower(table[i-1])
			// 	table[i+1]= ""
		}
	}

	result := DeleteCases("(low)", table)

	return result
}

func DeleteCases(cas string, table []string) []string {
	var result []string
	for i := 0; i < len(table); i++ {
		if table[i] != cas {
			result = append(result, table[i])
		}
	}
	return result
}

func withNumber(table []string) []string {
	fmt.Println(table)
	/*
		! makhdamach 3ndi prob fhad ex: hh hh hh hh hh hh (low, 6) (up, 4) ervreverv (7it kay7sb tal empty )
		! + khsni n9ad lout of range fakhir string
		! + lout of range fach kat7et gha (up, 2
		!  ====> clean the code and make it easy
	*/
	//	var result []string

	count := 0
	for i := 0; i < len(table)-1; i++ {

		count++

		num1 := 0
		if string(table[i]) == "(low," {
			num1 = MyAtoi(table[i+1])
			if count-1 < num1 {
				num1 = count - 1
			}

			for j := 1; j <= num1; j++ {
				table[i-j] = strings.ToLower(table[i-j])

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

		}
		if string(table[i]) == "(up," {
			num1 := 0
			num1 = MyAtoi(string(table[i+1]))

			if count-1 < num1 {
				num1 = count - 1
			}

			for j := 1; j <= num1; j++ {

				table[i-j] = strings.ToUpper(table[i-j])
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

		}
		if string(table[i]) == "(cap," {

			num1 = MyAtoi(string(table[i+1]))

			if count-1 < num1 {
				num1 = count - 1
			}

			for j := 1; j <= num1; j++ {
				word := table[i-j]
				fmt.Println(table)
				table[i-j] = strings.ToUpper(string(word[0])) + word[1:]
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

	return table
}

func findPunctuationIndex(word string) int {
	punctuations := []string{",", ".", "!", "?", ":", ";"}

	for _, punctuation := range punctuations {
		if strings.Contains(word, punctuation) {
			return strings.Index(word, punctuation)
		}
	}
	return -1
}

func punctuations(table []string) []string {
	var result []string
	corrWord := ""

	for i := 0; i < len(table); i++ {
		if string(table[i]) == "," {
			continue
		}
		if i >= 0 {
			if strings.Contains(table[i], ",") || strings.Contains(table[i], ".") || strings.Contains(table[i], "!") || strings.Contains(table[i], "?") || strings.Contains(table[i], ":") || strings.Contains(table[i], ";") {
				index := findPunctuationIndex(table[i])
				fmt.Println(corrWord)
				for k := 0; k < len(table[i]); k++ {
					corrWord += string(table[i][k])
					if k == index {
						corrWord += " "
					}
				}
				table[i] = corrWord
			} else if table[i] == "," || table[i] == "." || table[i] == "!" || table[i] == "?" || table[i] == ":" || table[i] == ";" {

				table[i-1] = table[i-1] + table[i]
				table[i] = ""
			} else if strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";") {
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

			}
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
	// ? you should to fix the error if you have just one marks  like=> hello 'reda 
	for i := 0; i < len(table); i++ {
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
				table[i+1] = "'" + strings.TrimSpace(table[i+1]) + "'"
			} else if strings.HasSuffix(table[i], "'") && strings.HasPrefix(table[i+2], "'") {
				table[i] = strings.TrimSuffix(table[i], "'")
				table[i+2] = strings.TrimPrefix(table[i+2], "'")
				table[i+1] = "'" + strings.TrimSpace(table[i+1]) + "'"
			} else if strings.HasSuffix(table[i], "'") && strings.HasPrefix(table[i+1], "'") {
				if len(table[i]) == 1 {
					table[i] = ""
				} else {
					table[i] = strings.TrimSuffix(table[i], "'")
				}
				table[i+1] = "'" + strings.TrimSpace(table[i+1][1:])
			}
		} else if strings.HasPrefix(table[i], "'") {
			table[i] = "'" + strings.TrimSpace(table[i][1:])
		} else if strings.HasSuffix(table[i], "'") {
			table[i] = "'" + strings.TrimSpace(table[i][:len(table[i])-1]) + "'"
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
