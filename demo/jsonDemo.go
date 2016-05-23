package main

import (
	"encoding/json"
	"fmt"
)

type TokenStr struct {
	Id          string   `json:"id"`
	Type        string   `json:"type"`
	AccessToken string   `json:"access_token"`
	Ids         []string `json:"ids"`
}

func main() {
	var token TokenStr
	text := `{"type": "facebook", "id": "5718732097", "access_token": "sdfsdtsetse", "ids": ["5718732097_10153714754837098", "5718732097_10153714372452098", "5718732097_10153711314922098"]}`
	json.Unmarshal([]byte(text), &token)
	fmt.Println("------------------")
	fmt.Println("no mark: ", token)
	fmt.Println("& mark: ", &token)

	fmt.Println("------------------")
	var token1 *TokenStr
	json.Unmarshal([]byte(text), &token1)
	fmt.Println("no mark: ", token1)
	fmt.Println("& mark: ", &token1)
	fmt.Println("* mark: ", *token1)
	fmt.Println("------------------")

}
