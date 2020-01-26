package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	dup := map[string][]string{}
	for i, record := range records {
		for _, inside := range records[i+1:] {
			if inside[3] == record[3] {
				dup[inside[3]] = inside
			}
		}
	}

	for k, _ := range dup {
		fmt.Println(k)
	}
}
