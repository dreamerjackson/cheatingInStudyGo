package tt_test

import (
	"fmt"
	"testing"
)

func kk() {
	fmt.Println("456")
}

func Test_kk(t *testing.T) {
	fmt.Println("123")
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"zjx"},
	}
	for _, _ = range tests {
		fmt.Println("123")
	}
}
