package test

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getKey() string {
	f, err := os.Open(".\\key.txt")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer f.Close()

	var key string
	s := bufio.NewScanner(f)
	if s.Scan() {
		key = s.Text()
	}

	fmt.Println(key)
	return key
}
