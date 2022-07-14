// go get -u github.com/sangx2/upbit
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/dataframe"

	"github.com/sangx2/upbit"
	"gopkg.in/yaml.v3"
)

// func getServerConfig(fileName string) (*Person, error) {
func GetApiKeys(apiName string) (map[string]interface{}, error) {
	path := "/home/ubuntu/dev/inGo/with_go/settings/coin_apis.yml"
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configs map[string]interface{}

	if err := yaml.Unmarshal(buf, &configs); err != nil {
		panic(err)
	}

	return configs[apiName].(map[string]interface{}), nil
}

func main() {

	keys, _ := GetApiKeys("upbit")
	// fmt.Println(keys["Accesskey"].(string))

	u := upbit.NewUpbit(keys["AccessKey"].(string), keys["SecretKey"].(string))

	accounts, remaining, e := u.GetAccounts()
	// account.Account
	// var accounts2 []struct
	if e != nil {
		fmt.Println("GetAccounts error : %s", e.Error())
	} else {
		fmt.Printf("GetAccounts[remaining:%+v]\n", *remaining)
		// df := dataframe.LoadStructs(accounts) // BUG: DataFrame error: load: type  (ptr slice) is not supported, must be []struct
		// // https://stackoverflow.com/questions/70781996/go-gota-dataframe-not-loading-struct-requires-struct
		// fmt.Println(df)
		meJson, err := json.Marshal(accounts)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(string(meJson))
		df := dataframe.ReadJSON(strings.NewReader(string(meJson)))
		fmt.Println(df)
		// for _, account := range accounts {
		// 	// accounts2.append(*account)
		// 	fmt.Printf("%T\n", *account)
		// 	// fmt.Printf("%+v\n", account.Balance)
		// }
		// fmt.Printf("%+v\n", accounts2)
	}
}
