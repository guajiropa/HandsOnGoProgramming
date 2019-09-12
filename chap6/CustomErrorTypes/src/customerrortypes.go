/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/11/2019
** SYNOPSIS	: Let's start with creating your custom error type. If you come from languages such as C#
** and Java, you may find the error mechanism a little different in Go. Moreover, the way you create
** your own custom error is very simple because Go is a duck-typed language, which means that you are
** good to go as long as your struct satisfies an interface. Let's go ahead and create our own custom
** error using a new type. So, I will have two fields, ShortMessage and the DetailedMessage of string
** type. You can have as many fields as you want, to capture more information about your errors.
** Furthermore, to satisfy the error interface, I'm going to implement a new method, *MyError, which
** will return a string value, and we can output this error to either our console or some log file.
**
** Then, what I'm going to do is to return the error message. So, the way you do this is very simple:
** you can just return this error type from your method. Let's imagine that we have a doSomething
** method that returns an error. Let's imagine we did some lines of codes in that method and it
** returns an error for some reason, such as a ShortMessage instance of "Wohoo something happened!".
** Of course, you will probably need to use more meaningful messages here, and don't forget to use this
** & operator. It will get the address of your *MyError object, since we are working with a pointer
** here. If you don't do this, you will see there's a type error, and one way to fix is this is to just
** remove that * pointer and the error will be will fixed. But you probably don't want to have multiple
** copies of the same object, so instead of doing what I have just described, you can easily do this:
** send a reference back so that you have better memory management.
 */
package main

import "fmt"

type MyError struct {
	ShortMessage    string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func main() {
	err := doSomething()
	fmt.Println(err)
}

func doSomething() error {
	//Doing something here . . .
	return &MyError{ShortMessage: "Oh oh, something went wrong!",
		DetailedMessage: "These command paramater are not allowed."}
}
