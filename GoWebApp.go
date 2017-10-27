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
}


func templateHandler(w http.ResponseWriter, r *http.Request){
    
    // declare game status message variable
    var status string

    // h2 header for game
    h2header := "Guess a number between 1 and 20"     

    // request cookie from the page
    cookies, err := r.Cookie("target")
    
    // if no cookie found, create a new cookie
    if err == http.ErrNoCookie {   
        cookies = &http.Cookie{
            Name: "target",
            Value: strconv.Itoa((rand.Intn(20)+1)), // generates random number between 1-20
            Expires: time.Now().Add(1 * time.Hour), // creates cookie that expires in one hour
        }
        // set the cookie
        http.SetCookie(w, cookies)
    }

    gameFinished := false 

    // get the guessed Number from the form
    guessedNumber,_ := strconv.Atoi(r.FormValue("guess"))

    // get the random number from the cookie
    randNum,_ := strconv.Atoi(cookies.Value)

    // check if the users guessed number and the random number matches
    if guessedNumber == randNum{
        status = "You have guessed the correct number. Click New Game for the next random number"
        //generate a new cookie with new random number
        cookies = &http.Cookie{
            Name: "target",
            Value: strconv.Itoa((rand.Intn(20)+1)),
            Expires: time.Now().Add(1 * time.Hour),
        }
    http.SetCookie(w, cookies)
    // set the game to finish
    gameFinished = true
    } else if (guessedNumber < randNum) && (guessedNumber > 0){
        status = "Number is too low"
    } else if guessedNumber > randNum{
        status = "Number is too high"
    }

    // setup the template to be executed
    msg  := &Message{S:h2header, GuessedNumber:guessedNumber, Status:status, GameFinished: gameFinished}

    // tells the page the location of the template files and executes them
    t, _ := template.ParseFiles("template/guess.tmpl")
    t.Execute(w,msg)

}


func main() {
    // generate seed for random number
    rand.Seed(time.Now().UTC().UnixNano())

    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/guess", templateHandler)
    http.ListenAndServe(":8080", nil)
}