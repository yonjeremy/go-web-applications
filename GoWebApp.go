package main

import (
	"io"
    "net/http"
    "text/template"
    "strconv"
    "time"
    "math/rand"
    

)

type Message struct {  
    S string
}

func templateHandler(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("template/guess.html")
    t.Execute(w,Message{S: "Guess a number between 1 and 20"})

     cookies, err := r.Cookie("thing")

     if err == http.ErrNoCookie {
        
        cookies = &http.Cookie{
            Name: "thing",
            Value: strconv.Itoa((rand.Intn(20)+1)),
            Expires: time.Now().Add(365 * 24 * time.Hour),

        }
        
    
    }

    http.SetCookie(w, cookies)

    io.WriteString(w,cookies.Value)

    


}


func main() {
        http.Handle("/", http.FileServer(http.Dir("./static")))
        http.HandleFunc("/guess", templateHandler)
        http.ListenAndServe(":8080", nil)
}