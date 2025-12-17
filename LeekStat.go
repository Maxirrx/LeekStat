package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

type game struct{
    Id int 
    Type int
} 

type res struct{
    Fights []game 
}


type data struct{
    Data tab
    Leeks1 []leek1
    Winner int
}

type leek1 struct{
    Farmer int
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
    Summon bool
}



func main() {

    res := getallFight()

    for _, element := range res{
    res, _ := getFightByGame(element)
    fmt.Printf("%v\n", res)
    }


}

func getallFight()[]int{
    response, _ := http.Get("https://leekwars.com/api/history/get-farmer-history/107035")
    var resu res
    json.NewDecoder(response.Body).Decode(&resu)
    res := []int{}
    for _, element := range resu.Fights{
    if element.Type == 1{
        res = append(res,element.Id)
    }}
    return res
}


func getFightByGame(id int) ([]leek , bool){
    response, _ := http.Get("https://leekwars.com/api/fight/get/" + strconv.Itoa(id))
    var data data
    json.NewDecoder(response.Body).Decode(&data)
    res := []leek{}
    for _,element := range data.Data.Leeks {
        if !element.Summon && element.Farmer != 107035{
            res = append(res, element)
        }
        fmt.Println()
    }
    winner := false
    if data.Winner == 1 && data.Leeks1[0].Farmer == 107035 || data.Winner == 2 && data.Leeks1[0].Farmer != 107035 {
        winner = true
    }


    return res, winner
}
