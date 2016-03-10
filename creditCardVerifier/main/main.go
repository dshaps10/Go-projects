package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckCard(card int) {
	stringCard := strconv.Itoa(card)
	stringSep := strings.Split(stringCard, "")
	fmt.Println(stringSep)
}

func main() {
	CheckCard(111111)
}
