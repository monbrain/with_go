// https://github.com/sangx2/upbit

package main

import (
	"fmt"

	"github.com/sangx2/upbit"
)

func main() {
	u := upbit.NewUpbit("", "")

	markets, remaining, e := u.GetMarkets()
	if e != nil {
		fmt.Println("GetMarkets error : %s", e.Error())
	} else {
		fmt.Printf("GetMarkets[remaining:%+v]\n", *remaining)
		for _, market := range markets {
			fmt.Printf("%+v, %v\n", *market, market.Market)
		}
	}
}
