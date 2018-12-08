package modules

import(
	"fmt"
	"os/exec"
)

func ShModule(n Notifications){
	if n.Sh != "" {
		script := n.Sh
		out, err := exec.Command("sh", script).Output()
		fmt.Println(string(out))
		if err != nil{
			fmt.Printf("[testurtle] error! %s\n", err)
		} else {
			fmt.Printf("[testurtle] %s is done.\n", script)
		}
	}
}