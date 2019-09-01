/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/31/2019
** SYNOPSIS	: we are going to learn how to find the difference between two dates. Let's imagine that
** we have two dates, as shown in the following code block, and you will see that the signature of
** this method is self-explanatory. So, all we have to do is use the following code for the subtract
** method which subtracts the first date from the second one. As you will see that the method returns
** Duration instead of the Time type between the dates.
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// Store two diffent dates to varaibles
	first := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	second := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)

	// Calculate the differenc
	diff := second.Sub(first)

	// Display the results
	fmt.Println("=====================================")
	fmt.Printf("\"first\" holds the value of %v \n", first)
	fmt.Println("=====================================")
	fmt.Printf("\"second\" holds the value of %v \n", second)
	fmt.Println("=====================================")
	fmt.Printf("The difference is %v \n", diff)
	fmt.Println("=====================================")
}
