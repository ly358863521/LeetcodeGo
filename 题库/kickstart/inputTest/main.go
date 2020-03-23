package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func bufioTest() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()
		res := strings.Split(string(input), " ")
		for _, i := range res {
			fmt.Println(strconv.Atoi(i))
		}
	}
}
func fmtTest() {
	fmt.Println("----------")
	var i, j int
	i = 0
	for {
		n, err := fmt.Scanf("%d", &j)
		if err != nil {
			log.Println(err)
		}
		i += n
		fmt.Println(j)
		if n == 0 {
			break
		}
	}
	fmt.Println("-------")
	fmt.Println(i)
}
func main() {
	bufioTest()
	// fmtTest()
}
