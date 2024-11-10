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
	d := make([][100]int, 100)
	last := 0
	for _, record := range records[1:] {
		fmt.Println(record[0], record[1])
		price := []rune(strings.Trim(record[1], ","))
		context := int((price[0] - '0') * 10)
		context += int(price[1] - '0')
		d[last][context]++
		last = context
	}
	fmt.Println(d[last])
	max, index := 0, 0
	for i, v := range d[last] {
		if v >= max {
			max, index = v, i
		}
	}
	fmt.Println(max, index)
}
