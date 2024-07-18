package model

import "fmt"

/*
{
  "bandToFind": "Red",
  "colors": [
       {"band": "Yellow"}, {"band": "Red"}
   ]
}
*/

// maps the individual color objects within the colors array
type Color struct {
    Band string `json:"band"`
}

type RequestData struct {
    BandToFind string  `json:"bandToFind"`
    Colors     []Color `json:"colors"`
}

func (r RequestData) print() {
    fmt.Printf("Band to find: %s\n", r.BandToFind)
}

