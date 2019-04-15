package search

import (
	"os"
	"fmt"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
	"math/rand"
	"encoding/json"

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


func CreateRandNumsJSON(filename string, num uint) {
	//1、生成10000个随机数
	rand.Seed(time.Now().UnixNano())

	var i uint
	var ints []int
	for i = 0; i < num; i++ {
		ints = append(ints, rand.Intn(10000))
	}

	//2、以JSON格式存入filename
	jints,_ := json.Marshal(ints)

	//此处写入的格式为不带双引号的[]int，而如果使用encoder.Encode(string(.))写入的文件则包含双引号；
	ioutil.WriteFile(filename, jints, 0666)

}

func GetRandNumsJSON(filename string) []int {
	//1、从filename中读取json格式的数据
	b, _ := ioutil.ReadFile(filename)

	//2、Unmarshal json格式的数据
	var data []int
	json.Unmarshal(b, &data)

	return data
}

func SortRandNumsJSON(originfile string, sortedfile string) error {
	//1、从originfile文件中读取json格式的数据;
	b, _ := ioutil.ReadFile(originfile)

	//2、Unmarshal json格式的数据并排序;
	var data []int
	json.Unmarshal(b, &data)
	sorting.QuickSort(data, 0, len(data)-1)

	//3、将排好序的data Marshal为json格式，并写入文件；
	db, _ := json.Marshal(data)
	return ioutil.WriteFile(sortedfile, db, 0666)
}


func CreateRandNumsJSONStream(filename string, n uint) {
	var i uint
	//1、生成10000个数
	rand.Seed(time.Now().UnixNano())
	var datas []int
	for i = 0; i < n; i++ {
		datas = append(datas, rand.Intn(10000))
	}

	//Marshal为json格式
	jdatas, _ := json.Marshal(datas)
	//2、写入stream
	file, _ := os.Create(filename)

	enc := json.NewEncoder(file)
	//此处如果是jdata将是对[]byte进行写入，打开文件会是各种字符；而以string的形式写入，可以明文看到数值；
	enc.Encode(string(jdatas))
}



