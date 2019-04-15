package search

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
	"math/rand"

	"github.com/lqlspace/algorithms/algorithms/sorting"
)



func CreateRandNums(filename string, nums uint) {
	rand.Seed(time.Now().UnixNano())
	var i uint 
	var str string
	for i= 0; i < nums; i++ {
		tmp := rand.Intn(10000)
		str += fmt.Sprintf("%d ", tmp)
	}
	strings.TrimRight(str, " ")
	ioutil.WriteFile(filename, []byte(str), 0666)
}


func GetRandNums(filename string) []int {
	byteInts, _ := ioutil.ReadFile(filename)
	strs := strings.Split(string(byteInts), " ")

	si := make([]int, len(strs))
	for i, v := range strs {
		e, _ := strconv.ParseInt(v, 10, 32)
		si[i] = int(e)
	}

	return si
}

func SortRandNums(originfile string, sortedfile string) error {
	byteInts, _ := ioutil.ReadFile(originfile)
	strs := strings.Split(string(byteInts), " ")
	var si = make([]int, len(strs))
	for i, v := range strs {
		e, _ := strconv.ParseInt(v, 10, 32)
		si[i] = int(e)
	}

	sorting.QuickSort(si, 0, len(si)-1)

	var str string
	for _, v := range si {
		str += fmt.Sprintf("%d ", v)
	}
	strings.TrimRight(str, " ")

	return ioutil.WriteFile(sortedfile, []byte(str), 0666)

}
