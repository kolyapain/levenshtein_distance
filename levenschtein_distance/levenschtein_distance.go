package levenschtein_distance

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Mmap struct {
	Val int
	Key string
}

func Sort(m []Mmap) []Mmap {
	for i := 0; i < len(m); i++ {
		for j := i; j < len(m); j++ {
			if m[i].Val > m[j].Val {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
	return m
}

func open_file(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

func Get_strings(fname string) (string, []string) {
	scanner := open_file(fname)
	scanner.Scan()
	var strlist []string
	var keyname string = scanner.Text()
	for scanner.Scan() {
		strlist = append(strlist, scanner.Text())
	}
	return keyname, strlist
}

func print(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func min3(val1, val2, val3 int) int {
	if val1 < val2 {
		if val1 < val3 {
			return val1
		}
		return val3
	}
	if val2 < val3 {
		return val2
	}
	return val3
}

func Dp_init(s1, s2 string) [][]int {
	dp := [][]int{}
	for i := 0; i <= len(s1); i++ {
		arr := make([]int, len(s2)+1)
		dp = append(dp, arr)
	}
	for i := 0; i <= len(s1); i++ {
		dp[i][0] = i
	}
	for i := 0; i <= len(s2); i++ {
		dp[0][i] = i
	}
	return dp
}

func Levenshtein_distance(s1, s2 string, dp [][]int) int {
	var c int
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				c = 0
			} else {
				c = 1
			}
			dp[i][j] = min3(dp[i-1][j-1]+c, dp[i-1][j]+1, dp[i][j-1]+1)
		}
	}
	return dp[len(s1)][len(s2)]
}

func Make_map(keyword string, strlist []string) []Mmap {
	m := []Mmap{}
	for i := range strlist {
		dp := Dp_init(keyword, strlist[i])
		res := Levenshtein_distance(keyword, strlist[i], dp)
		m = append(m, Mmap{res, strlist[i]})
	}
	return Sort(m)
}

func Print_map(tbl []Mmap) {
	file, err := os.Create("lev.log")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for i := range tbl {
		//fmt.Println(tbl[i].key, " : ", tbl[i].val)
		msg := tbl[i].Key + " : " + strconv.Itoa(tbl[i].Val) + "\n"
		file.WriteString(msg)
	}
}
