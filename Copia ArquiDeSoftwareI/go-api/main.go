package main

import (
	"encoding/json"
	"fmt"
	"go-api/items"
)

func main() {
	fmt.Println(items.ItemStore)

	item := items.GetItem()

	itemJSON, _ := json.MarshalIndent(item, "", "    ")

	fmt.Println(string(itemJSON))
}
