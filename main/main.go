package main

import (
	"fmt"
	"time"

	"../levenschtein_distance"
	op "../optimization"
)

func main() {
	keyname, strlist := levenschtein_distance.Get_strings("input.txt")
	start := time.Now()
	//levenschtein_distance.Make_map(keyname, strlist)
	levenschtein_distance.Print_map(levenschtein_distance.Make_map(keyname, strlist))
	time1 := time.Now().Sub(start)
	fmt.Println("not optimized finished in : ", time1)
	start = time.Now()
	//op.Make_map(keyname, strlist)
	op.Print_map(op.Make_map(keyname, strlist))
	time2 := time.Now().Sub(start)
	fmt.Println("optimized finished in : ", time2)
	fmt.Println("optimizated : ", (float64(time1)-float64(time2))/float64(time2)*100, "%")
}
