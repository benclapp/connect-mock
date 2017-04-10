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
	tt := make([]interface{}, 0)
	tt = unmarshalJsonArray(ticketTypesResponseData, tt)

	orderCreatedResponseData := readResponseData("OrderCreatedResponse")
	var oc interface{}
	unmarshalJsonObject(orderCreatedResponseData, &oc)

	checkoutResponseData := readResponseData("CheckoutResponse")
	var co interface{}
	unmarshalJsonObject(checkoutResponseData, &co)

	r.GET("/RESTData.svc/cinemas/:cinema/sessions/:session/tickets", func(c *gin.Context) {
		c.JSON(200, tt)
	})

	r.POST("/RESTTicketing.svc/order/tickets", func(c *gin.Context) {
		c.JSON(200, oc)
	})

	r.POST("/RESTTicketing.svc/order", func(c *gin.Context) {
		c.JSON(200, co)
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
func unmarshalJsonArray(ticketTypesResponseData []byte, arr []interface{}) []interface{} {
	if err := json.Unmarshal(ticketTypesResponseData, &arr); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return arr
}

func unmarshalJsonObject(ticketTypesResponseData []byte, obj interface{}) interface{} {
	if err := json.Unmarshal(ticketTypesResponseData, &obj); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return obj
}

func readResponseData(responseFileName string) []byte {
	responseData, e := ioutil.ReadFile("/tmp/" + responseFileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return responseData
}