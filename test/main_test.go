package test

import (
	"../levenschtein_distance"
	"testing"
)

func BenchmarkGet_strings(b *testing.B)  {
	for i:= 0; i < b.N; i++ {
		levenschtein_distance.Get_strings("../main/input.txt")
	}
}

func BenchmarkMake_map(b *testing.B)  {
	keyname, strlist := levenschtein_distance.Get_strings("../main/input.txt")
	for i:= 0; i < b.N; i++ {
		levenschtein_distance.Make_map(keyname, strlist)
	}
}