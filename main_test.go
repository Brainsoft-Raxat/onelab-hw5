package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserStruct(t *testing.T) {
	var rawJson = []byte(`[
  {
    "id": "1",
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 2,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

	var rawXml = []byte(`
	<User>
		<id>3</id>	
   		<address>
      		<city_id>4</city_id>
			<street>KBTU</street>
		</address>
		<age>20</age>
	</User>
`)
	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		panic(err)
	}

	if err := xml.Unmarshal(rawXml, &users); err != nil {
		panic(err)
	}

	wantUsers := []User{
		{
			ID: "1",
			Address: Address{
				CityID: "5",
				Street: "Satbayev",
			},
			Age: "20",
		},
		{
			ID: "2",
			Address: Address{
				CityID: "6",
				Street: "Al-Farabi",
			},
			Age: "32",
		},
		{
			ID: "3",
			Address: Address{
				CityID: "4",
				Street: "KBTU",
			},
			Age: "20",
		},
	}

	t.Run("test struct that contains field of type json.Number", func(t *testing.T) {
		assert.Equal(t, wantUsers, users)
	})
}
