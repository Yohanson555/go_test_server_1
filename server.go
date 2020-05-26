package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DefaultPort = 8080

func main() {
	args := mapArgs(os.Args)

	var port int64 = 8080

	if p, e := getServicePort(args); e == nil {
		port = p
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am net/http server: \n")
		fmt.Fprintf(w, "Host: %v \n", html.EscapeString(r.Host))
		fmt.Fprintf(w, "Remove addres: %v\n", html.EscapeString(r.RemoteAddr))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func mapArgs(args []string) map[string]string {
	var p interface{} = nil
	mappedArgs := map[string]string{}

	for _, s := range args {
		if p == nil {
			p = s
		} else {
			mappedArgs[p.(string)] = s
			p = nil
		}
	}

	return mappedArgs
}

func getServicePort(args map[string]string) (int64, error) {
	var port int64

	if val, ok := args["-p"]; ok {
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			if i > 0 && i < 65535 {
				port = i
			} else {
				return 0, fmt.Errorf("port value should be greater than 0 and less than 65535")
			}
		} else {
			return 0, err
		}
	} else {
		port = DefaultPort
	}

	return port, nil
}
