package main

import (
    "net/http"
    "text/template"
    "strconv"
    "time"
    "math/rand"
    

)

type Message struct {  
    S string
    GuessedNumber int
    Status string
    GameFinished bool
    IsNewGame bool
}


func templateHandler(w http.ResponseWriter, r *http.Request){
    
    var isNewGame bool
    h2header := "Guess a number between 1 and 20"     

    cookies, err := r.Cookie("thing")

     if err == http.ErrNoCookie {
        
        cookies = &http.Cookie{
            Name: "thing",
            Value: strconv.Itoa((rand.Intn(20)+1)),
            Expires: time.Now().Add(1 * time.Hour),
        }
    http.SetCookie(w, cookies)
    }

    gameFinished := false 

    //r.ParseForm()
    currentNumber,_ := strconv.Atoi(r.FormValue("guess"))
    //Number{GuessedNumber:r.Form["guess"][0]}

    var status string
    randNum,_ := strconv.Atoi(cookies.Value)
    if currentNumber == randNum{
        status = "You have guessed the correct number. Click New Game for the next random number"
        cookies = &http.Cookie{
            Name: "thing",
            Value: strconv.Itoa((rand.Intn(20)+1)),
            Expires: time.Now().Add(1 * time.Hour),
        }
    http.SetCookie(w, cookies)
    gameFinished = true
    } else if (currentNumber < randNum) && (currentNumber > 0){
        status = "Number is too low"
    } else if currentNumber > randNum{
        status = "Number is too high"
    }


    msg  := &Message{S:h2header, GuessedNumber:currentNumber, Status:status, GameFinished: gameFinished, IsNewGame:isNewGame}

    t, _ := template.ParseFiles("template/guess.html")
    //t.Execute(w,Message{})  
    t.Execute(w,msg)

}


func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/guess", templateHandler)
    http.ListenAndServe(":8080", nil)
}