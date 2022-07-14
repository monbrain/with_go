package datatype_test

import (
	"fmt"
	// "log"
	"testing"
	"with_go/datatype"
)

func TestStrings2Interfaces(t *testing.T) {
	temp := []string{"s1", "s2", "s3"}
	s := datatype.Strings2Interfaces(temp)
	if s == nil {
		t.Error("Wrong result")
	}
	// fmt.Printf("%T, %v", s, s)
}

func TestStrings2Interfaces2D(t *testing.T) {
	temp := [][]string{
		{"s1", "s2", "s3"},
		{"t1", "t2", "t3"},
	}
	s := datatype.Strings2Interfaces2D(temp)
	if s == nil {
		t.Error("Wrong result")
	}
	fmt.Printf("%T, %v", s, s)
}
