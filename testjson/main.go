package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type FruitBasket struct {
		Name  string
		Fruit []string
		//申明对应json 的key
		Id      int64 `json:"ref"`
		Created time.Time
	}

	jsonData := []byte(`
	{
        "Name": "Standard",
        "Fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket.Created, basket.Name, basket.Fruit, basket.Id)
	bytes, err := json.Marshal(basket)
	fmt.Println(string(bytes))
}
