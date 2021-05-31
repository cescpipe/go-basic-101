package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/shopspring/decimal"
)

type people struct {
	FirstName string          `json:"first_name" xml:"first_name"`
	LastName  string          `json:"last_name" xml:"last_name"`
	Age       int             `json:"age" xml:"age"`
	Address   address         `json:"address" xml:"address"`
	Salary    decimal.Decimal `xml:"salary>base_salary"`
}

type address struct {
	No       string `json:"no" xml:"no"`
	Moo      int    `json:"moo" xml:"moo"`
	District string `json:"district" xml:"district"`
}

func main() {
	someoneJSON := []byte(`{
					"first_name":"Gopher", 
					"last_name":"Conference",
					"age": 8,
					"address" : {"no":"904", "moo":10, "district":"Dusit", "province":"Bangkok"}
				}`)

	var p people
	if err := json.Unmarshal(someoneJSON, &p); err != nil {
		fmt.Println("error json unmarshall : ", err)
		return
	}

	fmt.Printf("JSON %+v", p)
	fmt.Println()

	someoneXML := []byte(`
				<people>
					<first_name>Gopher</first_name> 
					<last_name>Conference</last_name>
					<age>8</age>
					<address><no>904"</no><moo>10</moo><district>Dusit</district><province>Bangkok</province></address>
					<salary>
						<base_salary>5000</base_salary>
					</salary>
				</people>`)
	var p2 people
	if err := xml.Unmarshal(someoneXML, &p2); err != nil {
		fmt.Println("error xml unmarshall : ", err)
		return
	}

	fmt.Printf("XML %+v", p2)
}
