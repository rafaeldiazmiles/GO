package week1

import "fmt"

func MapIterate(kvs map[string]string) {
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}
}
