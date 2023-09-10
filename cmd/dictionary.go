package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const API_KEY = "91e5181d-789d-4f6b-9c58-72be96d0bbf7"

type dictionary struct {
	Meta     Meta       `json:"meta"`
	Hom      int64      `json:"hom,omitempty"`
	Hwi      Hwi        `json:"hwi"`
	FL       string     `json:"fl"`
	Def      []Def      `json:"def"`
	Et       [][]string `json:"et,omitempty"`
	Date     string     `json:"date,omitempty"`
	Shortdef []string   `json:"shortdef"`
}

type Def struct {
	Sseq [][][]interface{} `json:"sseq"`
	Vd   string            `json:"vd,omitempty"`
}

type Hwi struct {
	Hw  string `json:"hw"`
	Prs []PR   `json:"prs,omitempty"`
}

type PR struct {
	Mw    string `json:"mw"`
	Sound Sound  `json:"sound"`
}

type Sound struct {
	Audio string `json:"audio"`
	Ref   string `json:"ref"`
	Stat  string `json:"stat"`
}

type Meta struct {
	ID        string   `json:"id"`
	UUID      string   `json:"uuid"`
	Sort      string   `json:"sort"`
	Src       string   `json:"src"`
	Section   string   `json:"section"`
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}

func filterResult(searchKey string) string {
	var meaningOfWord []string
	var output string

	InfoLogger.Println("searchKey", searchKey)

	results := getWordsMeaning(searchKey)

	for i, result := range results {
		metaId := strings.ToLower(result.Meta.ID)

		if strings.Contains(metaId, ":1") {
			searchKey = searchKey + ":1"
		}

		if strings.Contains(metaId, searchKey) && result.FL == "noun" {
			InfoLogger.Println("result.Hwi.Prs[i].Mw", result.Hwi.Prs[i].Mw)
			InfoLogger.Println("result.FL", result.FL)
			InfoLogger.Println("fmt.Sprint(result.Shortdef)", strings.Join(result.Shortdef, " "))
			meaningOfWord = append(meaningOfWord, result.Hwi.Prs[i].Mw, " ("+result.FL+"): ", strings.Join(result.Shortdef, " "))
			output = result.Hwi.Prs[i].Mw + " " + " (" + result.FL + "): " + " " + strings.Join(result.Shortdef, " ")
			InfoLogger.Println("results", output)
		}
	}
	//output = strings.Join(meaningOfWord, " ")
	return output
}

func getWordsMeaning(searchKey string) []dictionary {
	var dicts []dictionary
	resp, err := http.Get("https://dictionaryapi.com/api/v3/references/collegiate/json/" + searchKey + "?key=" + API_KEY)

	if err != nil {
		ErrorLogger.Println("API-Access", err)
		return dicts
	}
	defer resp.Body.Close()
	body, jsonErr := io.ReadAll(resp.Body) // response body is []byte

	if jsonErr != nil {
		ErrorLogger.Println("API-Response-Body", jsonErr)
		return dicts
	}
	jsonParserErr := json.Unmarshal(body, &dicts)

	if jsonParserErr != nil {
		ErrorLogger.Println("JSON Parser", jsonParserErr)
		return dicts
	}
	return dicts
}
