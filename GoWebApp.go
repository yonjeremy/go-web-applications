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
}


func templateHandler(w http.ResponseWriter, r *http.Request){
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

    

    //r.ParseForm()
    currentNumber,_ := strconv.Atoi(r.FormValue("guess"))
    //Number{GuessedNumber:r.Form["guess"][0]}

    var status string
    randNum,_ := strconv.Atoi(cookies.Value)
    if currentNumber == randNum{
        status = "You have guessed the correct number. Click New Game for the next random number"
    } else if currentNumber < randNum{
        status = "Number is too low"
    } else if currentNumber > randNum{
        status = "Number is too high"
    }

    msg  := &Message{S:h2header, GuessedNumber:currentNumber, Status:status}

    t, _ := template.ParseFiles("template/guess.html")
    //t.Execute(w,Message{})  
    t.Execute(w,msg)

}


func main() {
        http.Handle("/", http.FileServer(http.Dir("./static")))
        http.HandleFunc("/guess", templateHandler)
        http.ListenAndServe(":8080", nil)
}