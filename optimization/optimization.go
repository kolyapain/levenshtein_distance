package optimization

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"

	ld "../levenschtein_distance"
)

type Results struct {
	Dists []ld.Mmap
	mtx   sync.Mutex
}

func (r *Results) Sort() {
	sort.Slice(r.Dists, func(i, j int) bool {
		if r.Dists[i].Val < r.Dists[j].Val {
			return true
		}
		return false
	})
}

func (r *Results) calcDst(keyword string, word2 string, wg *sync.WaitGroup) {
	defer wg.Done()
	dp := ld.Dp_init(keyword, word2)
	res := ld.Levenshtein_distance(keyword, word2, dp)
	r.mtx.Lock()
	r.Dists = append(r.Dists, ld.Mmap{res, word2})
	r.mtx.Unlock()
}

func Make_map(keyword string, strlist []string) []ld.Mmap {
	var dts Results
	var wg sync.WaitGroup
	dts.Dists = make([]ld.Mmap, 0, len(strlist)+1)
	for i := range strlist {
		wg.Add(1)
		dts.calcDst(keyword, strlist[i], &wg)
	}
	wg.Wait()
	dts.Sort()
	return dts.Dists
}

func Print_map(tbl []ld.Mmap) {
	file, err := os.Create("opt.log")
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
