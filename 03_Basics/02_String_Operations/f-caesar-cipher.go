package main

import "fmt"

// abcdefghijklmnopqrstuvwxyz
// EX: I am in mozambique now
//     L dp lq prcdpeltxh qrz

func main() {
	// fmt.Printf("\t rune       : %c \n", 'a') // a
	// fmt.Printf("\t rune       : %q \n", 'a') // 'a'
	// fmt.Printf("\t UTF8 code  : %v \n", 'a') // 97
	// fmt.Println()

	// fmt.Printf("\t rune       : %c \n", 'z')   // z
	// fmt.Printf("\t rune       : %q \n", 'z')   // 'z'
	// fmt.Printf("\t UTF8 code  : %v \n\n", 'z') // 122
	// fmt.Println()

	sentence := "planning for an attack at river bay at night 7 PM UTC"
	encryptedString := ""
	for i := 0; i < len(sentence); i++ {
		// fmt.Println(string(sentence[i]), string(sentence[i]+3))
		if sentence[i] == ' ' {
			encryptedString += string(sentence[i])

		} else {
			encryptedString += string(sentence[i] + 3)

		}

	}
	fmt.Println("encryptedString=", encryptedString)
}

//  sodqqlqj iru dq dwwdfn dw ulyhu ed| dw qljkw : SP XWF
// sodqqlqj iru dq dwwdfn dw ulyhu edb dw qljkw 7 SP XWF

// assignment -- update this code, to modify only the 
// limit of a-z and A-Z
// Also, consider working with both upper and lower cases
