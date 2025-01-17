package main

import "fmt"

const LUCKYNUMBER = 67

// Number Guessing Game
func main() {
	fmt.Print("Guess a number btwn 1 and 100:")

	var guessedNumber int
	fmt.Scanf("%d", &guessedNumber)
	fmt.Println("guessedNumber=", guessedNumber)

	if guessedNumber < 1 || guessedNumber > 100 {
		fmt.Println("You entered an invalid number")
	} else {

		// // Logic 1
		// if guessedNumber == LUCKYNUMBER {
		// 	fmt.Println("You guessed correctly!!!")
		// }

		// // Logic 2
		// if guessedNumber == LUCKYNUMBER {
		// 	fmt.Println("You guessed correctly!!!")
		// } else {
		// 	println("Please Try Again!!!")
		// }

		// Logic 3
		// 1 ------ 100
		//     67

		// 34	   --- increase their guess
		// 78		-- decrese their guess
		if guessedNumber == LUCKYNUMBER {
			fmt.Println("You guessed correctly!!!")
		} else if guessedNumber > LUCKYNUMBER { // 89 > 67
			fmt.Println("HINT:Please reduce your guessing number")
		} else { // 56 < 67
			fmt.Println("HINT:Please increase your guessing number")
		}

	}
}
