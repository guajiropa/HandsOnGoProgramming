/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/31/2019
** SYNOPSIS	: By using the AddDate function on the  time type, we can add as many years, months,
** and days as we want, since it accepts three parameters.
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// Get the current time/date
	currdate := time.Now()
	// Lets add 14 days to the current date
	twoweeks := currdate.AddDate(0, 0, 14)
	// Lets take two weeks off of the date
	twoweeksoff := currdate.AddDate(0, 0, -14)

	fmt.Println("==============================================")
	fmt.Printf("Today's date is %s \n", currdate.Format("01-02-2006"))
	fmt.Println("==============================================")
	fmt.Printf("In two weeks it will be %s \n", twoweeks.Format("01-02-2006"))
	fmt.Println("==============================================")
	fmt.Printf("In two weeks ago it was %s \n", twoweeksoff.Format("01-02-2006"))
	fmt.Println("==============================================")

}
