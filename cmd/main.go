package main

import (
	"bufio"
	"fmt"
	"os"
	"parentheses_validation/pkg/parentheses"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	sb := strings.Builder{}
	for scan.Scan() {
		sb.WriteString(scan.Text())
	}
	if parentheses.IsValid(sb.String()) {
		fmt.Println("ПРАВИЛЬНО")
	} else {
		fmt.Println("НЕПРАВИЛЬНО")
	}
}
