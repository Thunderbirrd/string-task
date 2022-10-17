package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var s string
	var s2 = ""
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	var count int
	for len(s) > 0 {
		el := s[0]
		count = strings.Count(s, string(el))
		s2 += string(el) + strconv.Itoa(count)
		s = s[count:]
	}
	fmt.Println(s2)
}
