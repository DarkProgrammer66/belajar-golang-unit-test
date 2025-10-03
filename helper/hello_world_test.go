package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAli(t *testing.T) {
	result := HelloWorld("Ali")
	if result != "Hello Ali" {
		// unit tes failed
		t.Error("Result must be 'Hello Ali'")
	}

	fmt.Println("TestHelloWorldAli Done")
}

func TestHelloWorldYudhati(t *testing.T) {
	result := HelloWorld("Yudhati")

	if result != "Hello Yudhati" {
		// error
		t.Fatal("Result must be 'Hello Yudhati'")
	}

	fmt.Println("TestHelloWorldYudhati Done")
}
