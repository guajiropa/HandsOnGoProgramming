/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/29/2019
** SYNOPSIS	: We can use time.Now to get today's date, and it imports a time package, and time
** returns a time type, and so we are going to assign this to another variable and use the
** String function.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	currtime := time.Now()
	fmt.Println("===========================================")
	fmt.Printf("The current time is : %s \n", currtime.String())
	fmt.Println("===========================================")
	fmt.Printf("It is now is : %s \n", currtime.Format("01-02-2006 15:04:05"))
	fmt.Println("===========================================")
}
