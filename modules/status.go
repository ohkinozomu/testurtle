package modules

import(
	"fmt"
	"net/http"
)

func StatusModule(r *http.Response, i Items, okNum int, ngNum int)(int, int){
	if i.Status != 0 {
		var err error
		fmt.Printf("[testurtle] %s : %d\n", i.URL, i.Status)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		b := r.StatusCode == i.Status
		okNum, ngNum = Judgement(b, okNum, ngNum)
	}
	return okNum, ngNum
}