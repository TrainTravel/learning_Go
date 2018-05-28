package main

import (
//        "encoding/json"
        "fmt"
        "net/http"
        "io/ioutil"
        "regexp"
        "log"
       )

type Store struct {
        ID                    int           `json:"id"`
        Name                  string        `json:"name"`
        Slug                  string        `json:"slug"`
        Description           string        `json:"description"`
        About                 string        `json:"about"`
        ImageURL              string        `json:"imageUrl"`
        ProductsImageURL      string        `json:"productsImageUrl"`
        BrandColor            string        `json:"brandColor"`
        Currency              string        `json:"currency"`
        SameStorePrice        interface{}   `json:"sameStorePrice"`
        BrandType             string        `json:"brandType"`
        PromotionText         string        `json:"promotionText"`
        ParentBrandID         interface{}   `json:"parentBrandId"`
        ProductsCount         int           `json:"productsCount"`
        StoreID               int           `json:"storeId"`
        ServiceType           string        `json:"serviceType"`
        FreeDeliveryEligible  bool          `json:"freeDeliveryEligible"`
        EstimatedDeLiveryTime interface{}   `json:"estimatedDe    liveryTime"`
        Closed                interface{}   `json:"closed"`
        OpensAt               interface{}   `json:"opensAt"`
        Tags                  []interface{} `json:"tags"`
        MinimumSpend          int           `json:"minimumSpend"`
        MinimumSpendExtraFee  int           `json:"minimumSpendExtraFee"`
        DeliveryTypes         []interface{} `json:"deliveryTypes"`
        ShippingModes         []string      `json:"shippingModes"`
        CashbackAmount        int           `json:"cashbackAmount"`
        ReservedTags          []struct {
                                ID       int    `json:"id"`
                                Title    string `json:"title"`
                                Key      string `json:"key"`
                                ImageURL string `json:"imageUrl"`
                                } `json:"reservedTags"`
    MinimumOrderFreeDelivery string `json:"minimumOrderFreeDelivery"`
        DefaultDeliveryFee       int    `json:"defaultDeliveryFee"`
        DefaultConciergeFee      int    `json:"defaultConciergeFee"`
        PricingType              string `json:"pricingType"`
        BrandTraits              []struct {
                Slug             string `json:"slug"`
                Name             string `json:"name"`
                Header           string `json:"header"`
                ImageURLBasename string `json:"imageUrlBasename"`
                } `json:"brandTraits"`
    FreeDeliveryAmount   int      `json:"freeDeliveryAmount"`
        IsImmediateDelivery  bool     `json:"isImmediateDelivery"`
        IsAllowNotifyMe      bool     `json:"isAllowNotifyMe"`
        SubstituteOptionList []string `json:"substituteOptionList"`
        ShippingMode         string   `json:"shippingMode"`
}

var HOST string = "https://www.honestbee.tw/"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
    stores_url := HOST + "/zh-tw/groceries/stores"
    req, err := http.NewRequest("GET", stores_url, nil)
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
    page, err := ioutil.ReadAll(resp.Body)
    if err != nil{
          log.Fatal("ReadAll: ", err)
          return
    }


    text := string(page)
    //re, _ := regexp.Compile("form.*groceries.*banners.*(ById.*shippingMode.*?regular.{3})")
    re, _ := regexp.Compile(`({"id".*Name.*shippingMode.*regular"})`)
    data := re.FindStringSubmatch(text)
    fmt.Println(data)
    fmt.Println(len(data))
    //bd := []byte(data)
    //err = ioutil.WriteFile("dat1", bd, 0644)
    //check(err)
    //re, _ := regexp.Compile("window.__data=(.*); window.__i18n")
    //stores := re.FindAllString(text, -1)[0]
    //var stores []Store
    //json.Unmarshal(bd, &stores)

    //fmt.Println(stores[0].Name)
          //fmt.Printf(stores)
          //fmt.Printf("%T \n", stores)
   // b, err := json.Marshal(stores)
          //fmt.Printf("%T \n", b)
          //fmt.Println(string(b))
}
