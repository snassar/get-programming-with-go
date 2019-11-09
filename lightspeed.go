package main

import "fmt"

func main() {
    const hoursPerday = 24
    var (
        distance = 96300000     // km
        speed = 100800          // km/h
    )

    fmt.Println(distance / speed / hoursPerday, "days") // shortest distance
}
