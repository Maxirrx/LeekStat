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
    Agility float64
    Farmer int
    Level float64
    Life float64
    Magic float64
    Mp float64
    Resistance float64
    Science float64
    Strength float64
    Tp float64
    Wisdom float64
    Summon bool
}



func main() {

    var idfarmer string
    fmt.Print("Entrez votre farmerID : ")
    fmt.Scanln(&idfarmer)

    getMoyenStatSimple(idfarmer)


}

func getallFight(farmer string)[]int{
    response, _ := http.Get("https://leekwars.com/api/history/get-farmer-history/" + farmer)
    var resu res
    json.NewDecoder(response.Body).Decode(&resu)
    res := []int{}
    for _, element := range resu.Fights{
    if element.Type == 1{
        res = append(res,element.Id)
    }}
    return res
}


func getFightByGame(id int, farmer int) ([]leek , bool){
    response, _ := http.Get("https://leekwars.com/api/fight/get/" + strconv.Itoa(id))
    var data data
    json.NewDecoder(response.Body).Decode(&data)
    res := []leek{}
    for _,element := range data.Data.Leeks {
        if !element.Summon && element.Farmer != farmer{
            res = append(res, element)
        }
    
    }
    winner := false
    if data.Winner == 1 && data.Leeks1[0].Farmer == farmer || data.Winner == 2 && data.Leeks1[0].Farmer != farmer {
        winner = true
    }


    return res, winner
}


func getMoyenStatSimple(farmer string){
    res := getallFight(farmer)
    total := leek{}
    win := leek{}
    loose := leek{}

    countertotal,counterwin ,counterloose := 0.0,0.0,0.0
    

    tab := []*leek {&total, &win, &loose}
    status := []string {"global","win","loose"}
    tabcounter := []*float64 {&countertotal, &counterwin, &counterloose}

    farmerstr, _ := strconv.Atoi(farmer)
    for _, element := range res{
        res, aswin := getFightByGame(element, farmerstr)
        for _, v := range res{
            total.Agility += v.Agility
            total.Level += v.Level
            total.Life += v.Life
            total.Magic += v.Magic 
            total.Mp += v.Mp
            total.Resistance += v.Resistance
            total.Science += v.Science
            total.Strength += v.Strength
            total.Tp += v.Tp
            total.Wisdom += v.Wisdom
            
            countertotal += 1
            if(aswin){
                win.Agility += v.Agility
                win.Level += v.Level
                win.Life += v.Life
                win.Magic += v.Magic 
                win.Mp += v.Mp
                win.Resistance += v.Resistance
                win.Science += v.Science
                win.Strength += v.Strength
                win.Tp += v.Tp
                win.Wisdom += v.Wisdom
            
                counterwin += 1
            }else{
                loose.Agility += v.Agility
                loose.Level += v.Level
                loose.Life += v.Life
                loose.Magic += v.Magic 
                loose.Mp += v.Mp
                loose.Resistance += v.Resistance
                loose.Science += v.Science
                loose.Strength += v.Strength
                loose.Tp += v.Tp
                loose.Wisdom += v.Wisdom
            
                counterloose += 1
            }
    
        }

    }

    for index,resu := range tab{
        

        fmt.Printf("Moyen %s des enemis\n\n",status[index])
        fmt.Printf("Niveau :%.2f\n" , resu.Level/ *tabcounter[index])
        fmt.Printf("PV :%.2f\n" , resu.Life/ *tabcounter[index])
        fmt.Printf("Point de tour :%.2f\n" , resu.Tp/ *tabcounter[index])
        fmt.Printf("Point de mouvement :%.2f\n" , resu.Mp/ *tabcounter[index])
        fmt.Printf("Résistance :%.2f\n" , resu.Resistance/ *tabcounter[index])
        fmt.Printf("Force :%.2f\n" , resu.Strength/ *tabcounter[index])
        fmt.Printf("Magie :%.2f\n" , resu.Magic/ *tabcounter[index])
        fmt.Printf("Sagesse :%.2f\n" , resu.Wisdom/ *tabcounter[index])
        fmt.Printf("Aglitié :%.2f\n" , resu.Agility/ *tabcounter[index])
        fmt.Printf("Science :%.2f\n\n\n\n\n" , resu.Science/ *tabcounter[index])

    }
    


}
