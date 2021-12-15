package project

import "fmt"

//How would you write a while statement in Go?
func While() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

//#####################################################################################

// What does the keyword defer do?

//It implements an inverse order. The deferred sentence will be executed last. Its a LIFO system (Last In First Out)

//#####################################################################################

// Does Go support pointers? How do arguments get passed around(by value or by reference)?

//Yes. By reference.

//#####################################################################################

// Are arrays in Go fixed length? How about slices?

//Go has a fixed size for Array (anArray [3]int) but we also have slices. Slices are dinamically sized, in reality they dont contain the data. They are actually structs with 3 fields: Length, Capactiy and Data where Data points to the holding array.

//#####################################################################################

//Say you have a map: map[string]int, how would do a lookup and check to see if the map holds the value of the key you were looking for?

//Lets give an example:

func KeySearchStringInt(a map[string]int, key string) {
	for k := range a {
		if k == key {
			fmt.Printf("Your key '%v' was founded\n", k)
			fmt.Printf("[\"%v\"] = %v\n", k, a["k"])
		}
	}
}

//#####################################################################################

// How does Go structure programs? What is the difference between a library and a program that executes?

//As I understand. Go has a "main" package where you execute your program and then it has its "libraries" or "packages". The main difference is that, as every go's files first line must be "package x", your libraries will have a different package name than main.

//#####################################################################################

// How To make a function or a type public? And how do you make it private?

//Go has a pretty simple way of doing this. If you start the signature's name of a function or variable with lowercase then its private by default, if you want it to be exported instead you'ld be using an Uppercase.

//#####################################################################################

// You are going to be building a simple calculator with 4 basic operations(add, subtract, multiply and divide). First build a library that provides those 4 methods. After that implement a program that reads from the command line the operation to be done and prints the result(by calling the library you implemented previously). The operation should be read however you'd like, but for simplicity sake limit yourself to 2 operands and 1 operation character. Something like ./program 1 + 2.

//#####################################################################################

//Suppose you are building a web server that needs a DB that can do a set of simple operations. You know that the requirements of what DB to use will change. You also know that it will be easier for testing purposes to not have to setup something like MySQL. How would you solve this problem using the feature that Go provide?

//I dont think I can answer this just yet.

//#####################################################################################

// How would you build a simple function that can receive any type of argument and prints the if that argument is of a primitive type. Limit to just int, string, float and bool.

func PrintType(x interface{}) string {
	switch x.(type) {
	case string, int, float64, bool:
		return "The argument is of a primitive type."
	default:
		return "The argument is not of an expected primitive type."

	}
}

//#####################################################################################

//How are errors defined in Go?

//As simple values using the error interface

//#####################################################################################

// Ok, you know how errors are defined in Go now. Time to build a simple errors package that allows you to build errors that specify what kind of error is it, limit yourself to 3 kinds: Internal, ThirdParty and Other. Then provide a function in that package for users to check if the error they have is of the kind they care about. NOTE: remember not to break with the way errors are defined in Go, take advantage of that.

//A revisar el paquete errors... no está bien hecho.

//#####################################################################################

// What do you use to make two functions concurrent?

//Cuando estudie concurrency...

//#####################################################################################

// How would you synchronize two concurrent functions?

//Cuando estudie concurrency...

//#####################################################################################

// Write a program with three functions. One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it. The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds. NOTE: the program can not end until the sender and receiver have returned.

//Cuando estudie concurrency...
