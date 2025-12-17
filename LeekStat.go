package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

type game struct{
    Id int 
} 

type res struct{
    Fights []game 
}

type data struct{
    Data []tab
}

type tab struct{
    Leeks []leek
}

type leek struct{
    Agility int
    Farmer int
    Level int
    Life int
    Magic int
    Mp int
    Resistance int
    Science int
    Strength int
    Tp int
    Wisdom int
}



func main() {
    all := getFightByGame(50465960)
    fmt.Println(all)
}

func getallFight()[]game{
    response, _ := http.Get("https://leekwars.com/api/history/get-farmer-history/107035")
    var resu res
    json.NewDecoder(response.Body).Decode(&resu)
    return resu.Fights
}


func getFightByGame(id int) []tab{
    response, _ := http.Get("https://leekwars.com/api/fight/get/" + strconv.Itoa(id))
    var data data
    json.NewDecoder(response.Body).Decode(&data)
    return data.Data
}
