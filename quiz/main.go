package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer (default 'problems.csv')")

	flag.Parse()

	fmt.Println("file:", *filePtr)
	f, err := os.Open(*filePtr)
	if err != nil {
		fmt.Printf("Invalid file")
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	reader := bufio.NewReader(os.Stdin)

	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	score := 0
	for _, d := range data {
		fmt.Printf("%s = ", d[0])
		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if strings.TrimRight(in, "\n") == d[1] {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d", score, len(data))
}
