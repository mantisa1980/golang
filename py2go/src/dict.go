package main

import "fmt"
// docs: https://blog.golang.org/go-maps-in-action

/*
python sample
x = {"A":1, "B":2 }
*/

/*
This variable m is a map of string keys to int values:
var m map[string]int

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:

m = make(map[string]int)


Maps are not safe for concurrent use

Iteration order is not guaranteed. Go 1 randomized iteration order 

*/


func main() {
    m := map[string]int {
        "ATK":250,
        "DEF":200,
    }
    m["DEX"] = 250

    fmt.Println("Basic access ------------- ")
    atk, ok := m["ATK"]
    fmt.Println(atk,ok)

    i, ok := m["SPD"]
    fmt.Println(i,ok, ok == false)

    //You cannot take the address of a map value because map values are not addressable. 
    //This restriction is in place because the map implementation may have to move values 
    // between hash map buckets, so callers must be prevented from taking the address of a map's value.
    
    //p_atk :=&m["ATK"]  // error錯誤
    
    fmt.Println("Interation ------------- ")
    for key, value := range m {
        fmt.Println("Key:", key, "Value:", value)
    }
    
}
