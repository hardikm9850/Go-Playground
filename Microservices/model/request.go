// maps the individual color objects within the colors array
package model

import (
	"fmt"
)

type Color struct {
	Band string `json:"band"`
}

type RequestData2 struct {
	BandToFind string  `json:"bandToFind"`
	Colors     []Color `json:"colors"`
}

func (r RequestData2) PrintMe() {
	fmt.Printf("Band to find: %s\n", r.BandToFind)
}

/*
Request data - 
{
  "bandToFind": "Red",
  "colors": [
       {"band": "Yellow"}, {"band": "Red"}
   ]
}
*/