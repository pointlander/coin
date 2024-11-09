// Copyright 2024 The Coin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("bitcoin.csv")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := csv.NewReader(input)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	d := make([]int, 100)
	for _, record := range records[1:] {
		fmt.Println(record[0], record[1])
		price := []rune(strings.Trim(record[1], ","))
		context := (price[0] - '0') * 10
		context += price[1] - '0'
		d[context]++
	}
	fmt.Println(d)
}
