package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

func main() {
	url := "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var list []string
	for k, v := range data {
		t := v.(map[interface{}]interface{})["type"]
		if t != "programming" {
			continue
		}

		list = append(list, k.(string))
	}

	sort.Strings(list)

	for _, name := range list {
		fmt.Println(name)
	}
}
