package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run() // eksekusi semua unit test

	fmt.Println("Setelah Unit Test")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ali")
	require.Equal(t, "Hello Ali", result, "Result must be 'Hello Ali'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ali")
	assert.Equal(t, "Hello Ali", result, "Result must be 'Hello Ali'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak bisa jalan di Windows")
	}

	result := HelloWorld("Ali")
	require.Equal(t, "Hello Ali", result)
}

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
