package main

import (
"fmt"
"log"
"errors"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}
// docker command to build a mongo:
// docker run  -d -ti --name mongo -p 27017:27017 mongo
func main() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
                   &Person{"Cla", "+55 53 8402 8510"})
    
    if err != nil {
            log.Fatal(err)
    }

    //e := errors.New("error---")  // create common error handler object
    //fmt.Println(e)

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone":0}).One(&result) // "phone":0 -> do not fetch phone field.
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("Phone:", result)

}
