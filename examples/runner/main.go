package main

import (
	"fmt"

	"github.com/aethosio/aego/utils"
)

// Run the examples

var rt utils.Runtime

func init() {
	rt = utils.Runtime{
		Packages: []string{"aego"},
	}
}

func main() {
	err := functionThatFails()
	if err != nil {
		fmt.Printf("functionThatFails: %s\n", err.Error())
	}

	err = notImplementedFunction()
	if err != nil {
		fmt.Printf("functionThatFails: %s\n", err.Error())
	}
}

func functionThatFails() error {
	return rt.Errorf("failure")
}

func notImplementedFunction() error {
	return rt.NotImplemented()
}
