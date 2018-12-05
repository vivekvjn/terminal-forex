package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Get all currency rates from Open Exchange Rates API
func getallCurrency() {

	var appId, ok = os.LookupEnv("OER_APP_ID")

	if !ok {
		fmt.Println("first export OER_APP_ID to use this tool")
		return
	}

	// Get data in JSON format from Open Exchange Rates API using a Key
	response, err := http.Get("http://openexchangerates.org/api/latest.json?app_id=" + appId)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// Removing the leading/trailing white spaces in the data
	data := string(body)
	data = strings.TrimSpace(data)

	// Parsing the JSON data
	parseData := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &parseData)
	if err != nil {
		panic(err)
	}

	// Iterate, format and display the data interactively
	for index := range parseData {
		if index != "rates" {
			fmt.Println(index, ":", parseData[index])
		} else {
			currencies := parseData["rates"].(map[string]interface{})
			for currency := range currencies {
				fmt.Println(currency, ":", currencies[currency].(float64))
			}
		}
	}

}

// Get currency details from Open Exchange Rates
func main() {

	getallCurrency()

}
