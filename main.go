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
	text := TableToString(table)
	// تحويل الجدول إلى نص
	data := []byte(text) // كتابة النص إلى الملف
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
func punctuations(table []string) []string {
	

for i := 0; i < len(table); i++ {
	if  i >0  {
		
		if table[i] == "," {
			table[i-1] = table[i-1]+table[i]
			table[i]= ""
		}else if strings.HasPrefix(table[i],",")  {

			

table[i-1] =  table[i-1]  + string(table[i][0])
table[i] = table[i][1:]

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

	fmt.Println(MyAtoi("55)"))
}
