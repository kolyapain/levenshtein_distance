package main

import (
	"fmt"
	"time"
	"../levenschtein_distance"
)

func main() {
	start := time.Now()
	keyname, strlist := levenschtein_distance.Get_strings("input.txt")
	levenschtein_distance.Print_map(levenschtein_distance.Make_map(keyname, strlist))
	fmt.Println("finished in : ", time.Now().Sub(start) )
}