package turtle

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"
)

type Check struct {
    URL string `json:"URL"`
    Target string `json:"target"`
}

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(config string){
	var configFile string
	if config != ""{
		configFile = config
	} else {
		configFile = "turtleconfig.json"
	}
	bytes, err := ioutil.ReadFile(configFile)
    if err != nil {
        fmt.Println(err)
    }
    var checks []Check
    if err := json.Unmarshal(bytes, &checks); err != nil {
        fmt.Println(err)
    }
	Patrol(checks)
}

func Patrol(checks []Check){
	okNum := 0
	ngNum := 0
	for _, c := range checks {
		fmt.Printf("[testurtle] %s : %s\n", c.URL, c.Target)
		r, err := http.Get(c.URL)
		if err != nil{
			fmt.Println(err)
		}
		buf := new(bytes.Buffer)
    	buf.ReadFrom(r.Body)
    	newStr := buf.String()
		b := strings.Contains(newStr, c.Target)
		if b == true{
			fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>","ok")
			okNum++
		} else {
			fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>","ng")
			ngNum++
		}
		fmt.Printf("[testurtle] Test completed. OK: \x1b[32m%d\x1b[0m  NG: \x1b[31m%d\x1b[0m\n", okNum, ngNum)
    }
}
