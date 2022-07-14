// https://github.com/gocolly/colly

// [크롤러 구현](https://m.blog.naver.com/PostView.naver?isHttpsRedirect=true&blogId=pjt3591oo&logNo=221346169732)

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "github.com/antchfx/htmlquery"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	// "with_go/datatype"
	// "with_go/cloud"
	"with_go/cloud"
	"with_go/datatype"
)

const (
	BASE_URL  = "https://docs.upbit.com/reference/"
	BASE_REST = "https://api.upbit.com/v1/"
)

func getHtml() (string, error) {
	url := "https://docs.upbit.com/reference/개별-주문-조회"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

// func checkSubstrings(str string, subs ...string) (bool, int) {

// 	matches := 0
// 	isCompleteMatch := true

// 	fmt.Printf("String: \"%s\", Substrings: %s\n", str, subs)

// 	for _, sub := range subs {
// 		if strings.Contains(str, sub) {
// 			matches += 1
// 		} else {
// 			isCompleteMatch = false
// 		}
// 	}

// 	return isCompleteMatch, matches
// }

func isInString(str string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(str, sub) {
			return true
		}
	}

	return false
}

func getStringsFromTable(table *html.Node, pres []string) [][]string {
	rows := [][]string{}
	for _, row := range htmlquery.Find(table, "//tr") {
		cells := pres // pres: 추가 셀 []string{}
		for _, cell := range htmlquery.Find(row, "//th | //td") {
			// cell = strings.TrimSuffix(cell, "\n")  // NOTE: 개행문자 제거
			// cell = strings.TrimRight(htmlquery.InnerText(cell), "\n") // NOTE: 개행문자 제거
			cells = append(cells, strings.TrimRight(htmlquery.InnerText(cell), "\n"))
			// cells = append(cells, `"`+htmlquery.InnerText(cell)+`"`)
		}
		// NOTE: 제목 셀인 경우 제외
		if !isInString(cells[len(pres)], "필드", "Name") {
			rows = append(rows, cells)
		}

	}
	return rows
}

// https://github.com/antchfx/htmlquery
// func GetApiList() (string, error) {
func GetApiList() [][]string {

	var apiInfo []string
	var apiList [][]string
	doc, _ := htmlquery.LoadURL("https://docs.upbit.com/reference/개별-주문-조회")
	// doc, _ := htmlquery.LoadDoc("./reference/01_accounts.html")

	nodes := htmlquery.Find(doc, `//article[@id="content"]//nav[@id="reference-sidebar"]/section/div`) // NOTE: 4개 (2개여야 함!!)

	for _, node := range nodes[:2] {
		category1 := htmlquery.Find(node, "./h3") // category1

		sect2s := htmlquery.Find(node, "./ul/li")
		for _, sect2 := range sect2s {
			category2 := htmlquery.Find(sect2, "./a/span/span") // category2
			sect3s := htmlquery.Find(sect2, "./ul/li")
			for _, sect3 := range sect3s {
				names := htmlquery.Find(sect3, "./a/span/span") // names
				hrefs := htmlquery.Find(sect3, "./a/@href")     // hrefs
				href := strings.Replace(htmlquery.SelectAttr(hrefs[0], "href"), "/reference/", "", 1)
				apiInfo = append(apiInfo, htmlquery.InnerText(category1[0]), htmlquery.InnerText(category2[0]), htmlquery.InnerText(names[0]), href)
				// fmt.Printf("%v\n", apiInfo)
				apiList = append(apiList, apiInfo)
				apiInfo = []string{}
			}
		}
	}
	return apiList
}

func GetApiDetail(apiInfo []string) ([]string, [][]string, [][]string) {
	url := BASE_URL + apiInfo[3]
	doc, _ := htmlquery.LoadURL(url)

	// doc, _ := htmlquery.LoadDoc("./reference/02_chance.html")
	root := htmlquery.Find(doc, `//article[@id="content"]`)[0] // NOTE:
	// fmt.Printf("%v\n", htmlquery.OutputHTML(root, true))

	headers := htmlquery.Find(root, "./header/div[2]/span")
	method := htmlquery.InnerText(headers[0])
	uri := htmlquery.InnerText(headers[1])
	nick := strings.Split(uri, "/{")[0] // NOTE: {unit} 제거
	nick = strings.Replace(nick, BASE_REST, "", 1)
	nick = strings.Replace(nick, "/", "_", -1)
	nick = method + "_" + nick
	desc := ""
	descNodes := htmlquery.Find(root, `//article[@id="content"]/header/div[@role="doc-subtitle"]/p`)
	if len(descNodes) > 0 {
		desc = htmlquery.InnerText(descNodes[0])
	}

	apiInfo = append(apiInfo, nick, uri, method, desc)

	// Request Parameters, Response
	var reqs, ress [][]string
	pres := []string{nick} // 추가 셀
	nodes := htmlquery.Find(root, `//div[@class="markdown-body"]`)

	if len(nodes) > 0 {
		titles := htmlquery.Find(nodes[0], `./*[contains(@class, "heading")]`)
		// titles := htmlquery.Find(node, `./h2`) // BUG!
		// fmt.Printf("titles 개수: %d\n", len(titles))
		for _, title := range titles {
			if strings.Contains(htmlquery.InnerText(title), "Request") {
				reqs = getStringsFromTable(htmlquery.Find(title, `./following-sibling::div//table`)[0], pres) // nick + ","
				// fmt.Println(reqs)
				// TODO: getStringsFromTable()
			}
			if strings.Contains(htmlquery.InnerText(title), "Response") {
				ress = getStringsFromTable(htmlquery.Find(title, `./following-sibling::div//table`)[0], pres)
				// fmt.Println(ress)
			}
		}
	} else {
		reqs = nil
		ress = nil
	}

	return apiInfo, reqs, ress
}

func saveSpecs() {
	apiInfoList := [][]string{{"category1", "category2", "name", "url_ref_suffix", "nick", "url_api", "method", "description"}}
	paramsList := [][]string{{"nick", "name", "description", "type"}}
	responseList := [][]string{{"nick", "name", "description", "type"}}

	apiList := GetApiList()

	// for _, api := range apiList[:3] {
	for _, api := range apiList {
		apiInfo, params, response := GetApiDetail(api)
		apiInfoList = append(apiInfoList, apiInfo)
		if params != nil {
			paramsList = append(paramsList, params...)
		}
		if response != nil {
			responseList = append(responseList, response...)
		}
	}

	// NOTE: Write to google sheets
	srv := cloud.SrvSheets("sheets", "moonsats", "")
	spreadsheetId := "1N7RoA0lSICAL3FzygOs0xKRalAsSJECg36KjbmLcQ5g"

	// apiInfo
	cloud.WriteSheet(datatype.Strings2Interfaces2D(apiInfoList), srv, spreadsheetId, "list")

	// params
	cloud.WriteSheet(datatype.Strings2Interfaces2D(paramsList), srv, spreadsheetId, "params")

	// response
	cloud.WriteSheet(datatype.Strings2Interfaces2D(responseList), srv, spreadsheetId, "response")
}

func main() {
	saveSpecs()
}
