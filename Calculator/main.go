package main

import (
	"fmt"
	//"math"
)

func main () {

	choose_arith := " What kind of arithmetic do you want to use? \n" + "1: add \n" + "2: subtract \n" + "3: multiply \n" + "4: divide \n"

	fmt.Println("%s", choose_arith) 
	var userinput_arith int
	fmt.Scan(&userinput_arith)

	//if  userinput_arith > 4{
	//	fmt.Println("Please choose between 1 -4")
	//	return fa
	//}
	 fmt.Println("Please enter first digit")
	 var userinput_firstdigit int
	 fmt.Scan(&userinput_firstdigit)

	 fmt.Println("Please enter second digit")
	 var userinput_seconddigit int
	 fmt.Scan(&userinput_seconddigit)

	 fmt.Println("First digit: ", userinput_firstdigit,  "\n",  "Second digit: ", userinput_seconddigit)
	 switch {
		case userinput_arith = 1:
			//add
			var calcu_add int
			calcu_add = userinput_firstdigit + userinput_seconddigit
			fmt.Println("Your Result is: ", calcu_add)
		case userinput_arith = 2:
			//subtract
			var calcu_subtract int
			calcu_subtract = userinput_firstdigit - userinput_seconddigit
		case userinput_arith = 3:
			// multiply
			var calcu_multiply int
			calcu_multiply = userinput_firstdigit * userinput_seconddigit
		case userinput_arith = 4:
			//divide
			var calcu_divide int
			calcu_divide = userinput_firstdigit / userinput_seconddigit
		default: 
			fmt.Println("Please start again and enter an number between 1 and 4")
		}

 }