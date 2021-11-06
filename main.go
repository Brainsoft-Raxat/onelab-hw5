package main

import (
	"encoding/json"
)



func main() {

}

type User struct {
	ID      json.Number `json:"id" xml:"id"`
	Address Address     `json:"address" xml:"address"`
	Age     json.Number `json:"age" xml:"age"`
}

type Address struct {
	CityID json.Number `json:"city_id" xml:"city_id"`
	Street string      `json:"street" xml:"street"`
}
