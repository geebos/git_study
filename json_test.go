package study

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(len("git"))
	fmt.Println(len("git"))
}

func TestB(t *testing.T) {
	fmt.Println("TestB is tested.")
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain is called.")
	m.Run()
}
