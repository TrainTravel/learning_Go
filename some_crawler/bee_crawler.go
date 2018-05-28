package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        //"io/ioutil"
        "log"
       )

type BeeStore struct {
        ID                       int           `json:"id"`
        Name                     string        `json:"name"`
        PickupPoint              string        `json:"pickupPoint"`
        Slug                     string        `json:"slug"`
        BrandID                  int           `json:"brandId"`
        AddressID                int           `json:"addressId"`
        CatalogID                int           `json:"catalogId"`
        Priority                 interface{}   `json:"priority"`
        Notes                    string        `json:"notes"`
        Description              string        `json:"description"`
        ImageURL                 interface{}   `json:"imageUrl"`
        Closed                   interface{}   `json:"closed"`
        TemporarilyClosed        bool          `json:"temporarilyClosed"`
        OpensAt                  interface{}   `json:"opensAt"`
        EstimatedDeliveryTime    interface{}   `json:"estimatedDeliveryTime"`
        BufferTime               int           `json:"bufferTime"`
        DeliveryTypes            []interface{} `json:"deliveryTypes"`
        ShippingModes            []string      `json:"shippingModes"`
        StoreType                string        `json:"storeType"`
        MinimumOrderFreeDelivery interface{}   `json:"minimumOrderFreeDelivery"`
        DefaultDeliveryFee       interface{}   `json:"defaultDeliveryFee"`
        FreeDeliveryEligible     bool          `json:"freeDeliveryEligible"`
        MinimumSpend             interface{}   `json:"minimumSpend"`
        MinimumSpendExtraFee     string        `json:"minimumSpendExtraFee"`
}

var url string = "https://www.honestbee.tw/api/api/stores/8674?sort=ranking&page=1&fields%5B%5D=departments"

func main () {
    //res, _ := http.Get( HOST + "/zh-tw/groceries/stores" )
    //res, _ := http.Get( url )
    //res, _ := http.Get( url )
    req, err := http.NewRequest("GET", url, nil)
    if err != nil{
        log.Fatal("NewRequest: ", err) 
        return
    }

    // For control over HTTP client headers,
    // redirect policy, and other settings,
    // create a Client
    // A Client is an HTTP client
    client := &http.Client{}

    // For control over HTTP client headers,
    // redirect policy, and other settings,
    // create a Client
    // A Client is an HTTP client
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal("Do: ", err) 
        return
    }
    
    defer resp.Body.Close()

    // Fill the store with the JSON data 
    var store BeeStore
    fmt.Println( store.Name )
    if err := json.NewDecoder(resp.Body).Decode(&store); err != nil {
        log.Println(err)
    }
    fmt.Println( store.Name )
}
