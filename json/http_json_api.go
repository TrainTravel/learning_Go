package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/url"
    "net/http"
)

type Numverify struct {
        Valid               bool   `json:"valid"`
        Number              string `json:"number"`
        LocalFormat         string `json:"local_format"`
        InternationalFormat string `json:"international_format"`
        CountryPrefix       string `json:"country_prefix"`
        CountryCode         string `json:"country_code"`
        CountryName         string `json:"country_name"`
        Location            string `json:"location"`
        Carrier             string `json:"carrier"`
        LineType            string `json:"line_type"`
}
