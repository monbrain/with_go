package datatype

// import (
//     "fmt"
//     // "strings"
// )

func Strings2Interfaces(strs []string) []interface{} {
	ret := make([]interface{}, len(strs))
	for i, v := range strs {
		ret[i] = v
	}
	return ret
}

func Strings2Interfaces2D(strs [][]string) [][]interface{} {
	var rets [][]interface{}
	for _, s := range strs {
		rets = append(rets, Strings2Interfaces(s))
	}

	return rets
}
