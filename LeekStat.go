package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {

    type game struct{
        id [16]int
    } 

    type reponseApi struct{
        fights []game
    } 

    response, err := http.Get("https://leekwars.com/api/history/get-farmer-history/107035")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	dec := json.NewDecoder(strings.NewReader(string(responseData)))
	for {
    var fights game
		if err := dec.Decode(&fights); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println( fights)
	}
    



}