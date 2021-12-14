package week1

import "fmt"

func MapStructDemo() {
	var mOne map[string]someStruct //nil map
	mTwo := map[string]someStruct{
		"key1": {6, false},
		"key2": {7, true},
	}
	fmt.Printf("mOne[\"key1\"].intField=%v\n", mOne["key1"].intField)
	fmt.Printf("mTwo[\"key1\"].intField=%v\n", mTwo["key1"].intField)
	mTwo["key3"] = someStruct{8, false} //Allocate new key-value in mTwo
	fmt.Printf("m[\"key3\"].boolField=%v\n", mTwo["key3"].intField)
}
