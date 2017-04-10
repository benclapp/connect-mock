package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"

)


func main() {
	r := gin.Default()

	ticketTypesResponseData := readResponseData("TicketTypesResponse")
	orderCreatedResponseData := readResponseData("OrderCreatedResponse")
	checkoutResponseData := readResponseData("CheckoutResponse")

	r.GET("/RESTData.svc/cinemas/:cinema/sessions/:session/tickets", func(c *gin.Context) {
		tt := make([]interface{}, 0)
		if err := json.Unmarshal(ticketTypesResponseData, &tt); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Printf("Results: %v\n", tt)
		c.JSON(200, tt)
	})

	r.POST("/RESTTicketing.svc/order/tickets", func(c *gin.Context) {
		var oc interface{};
		if err := json.Unmarshal(orderCreatedResponseData, &oc); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Printf("Results: %v\n", oc)
		c.JSON(200, oc)
	})

	r.POST("/RESTTicketing.svc/order", func(c *gin.Context) {
		var co interface{};
		if err := json.Unmarshal(checkoutResponseData, &co); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Printf("Results: %v\n", co)
		c.JSON(200, co)
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}



func readResponseData(responseFileName string) []byte {
	responseData, e := ioutil.ReadFile("/tmp/" + responseFileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return responseData
}