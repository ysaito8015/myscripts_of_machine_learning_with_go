package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// CSVファイルを開く
	f, err := os.Open("myfile.csv")
	if err != nil {
		log.Fatal(err)
	}

	// CSVレコードを読み込む
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// integer column においての最大値を取得する
	var intMax int
	for _, record := range records {

		// integer column をパースする
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// 値を評価して最大値とする
		if intVal > intMax {
			intMax = intVal
		}
	}

	// 最大値を表示する
	fmt.Println(intMax)
}
