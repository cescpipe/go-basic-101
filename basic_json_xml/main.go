package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"time"
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

	fmt.Println("START CALL HTTP GET")

	type commentStruct struct {
		PostId int    `json:"postId"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		fmt.Println("error call https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read  https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	var csGet []commentStruct

	if err := json.Unmarshal(body, &csGet); err != nil {
		fmt.Println("error unmarshall https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	fmt.Printf("%#v : ", csGet)
	fmt.Println("END CALL HTTP GET")

	fmt.Println("START CALL HTTP POST ")

	var inputData struct {
		UserId int `json:"userId"`
	}

	idb, err := json.Marshal(inputData)
	if err != nil {
		fmt.Println("error json marshall https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	resp2, err := client.Post("https://jsonplaceholder.typicode.com/comments", "application/json", bytes.NewBuffer(idb))
	if err != nil {
		fmt.Println("error call https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("error read  https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	var csPost commentStruct

	if err := json.Unmarshal(body2, &csPost); err != nil {
		fmt.Println("error unmarshall https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	fmt.Printf("%#v : ", csPost)
	fmt.Println("END CALL HTTP POST")

}
