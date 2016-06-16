package main
import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"log"
	"strings"
	"bytes"
//	_ "github.com/mattn/go-sqlite3"
	//	"database/sql"
	"encoding/json"
	"net/http"
	
)
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


type Message struct {
	State string
	Guest string
	Uuid string
//	Uefi string
	CPU string
	Memory string
	MemoryResident string
	
}



func vmInfo(vmname string)(jsonout []byte){
	cmd := exec.Command("vm", "info", vmname)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	why := out.String()
	zee := strings.Split(why, "\n",)
	var State string
	var Guest string
	var Uuid string
	var CPU string
	var Memory string
	var MemoryResident string
		

	for _, zz := range zee {
		zzz := strings.Replace(zz, " ", "", -1)
		zuu := strings.SplitAfterN(zzz, ":", 2)
	//	fmt.Println(zuu[:1], zuu[1:2])
		first := strings.Join(zuu[:1], "")
		second := strings.Join(zuu[1:2], "")
		varname := strings.Replace(first, ":", "", -1)
		//	fmt.Println(varname, zuu[1:2])

		
		switch varname {
		case "state":
			State = second
			fmt.Println(second)
		case "guest":
			Guest = second
		case "uuid":
			Uuid = second
		case "cpu":
			CPU = second
		case "memory":
			Memory = second
		case "memory-resident":
			MemoryResident = second			
		}
	}
			m := Message{State, Guest, Uuid, CPU, Memory, MemoryResident}
		jsonout, _ = json.Marshal(m)

	return
	
}




func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
	
}



func mainHandler (w http.ResponseWriter, r *http.Request) {
	filess, _ := ioutil.ReadDir("/dev/vmm/")
		for _, ff := range filess {
		fmt.Println(ff.Name())
		what := ff.Name()
			js := vmInfo(what)
		w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	
	}


}
