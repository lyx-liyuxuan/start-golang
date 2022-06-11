package mapreduce

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	nNumber = 100
)

// 创建一个包含N个编号的输入文件
// 通过mapReduce进行处理
// 检查最终输出文件中是否包含了N个编号

// 自定义map分割处理函数
func Mapfunc(file string, value string) (res []KeyValue) {
	words := strings.Fields(value)
	for _, w := range words {
		kv := KeyValue{w, ""}
		res = append(res, kv)
	}
	return
}

// 自定义reduce汇总函数
func Reducefunc(key string, values []string) string {
	for _, element := range values {
		fmt.Printf("Reduce %s-%v\n", key, element)
	}
	return ""
}

// 顺序执行的mapreduce
func TestSquentialSingle(t *testing.T) {
	SequentialSingle("test", makeInputs(1), 1, Mapfunc, Reducefunc)
}

//func TestSquentialMulti(t *testing.T) {
//	Sequential("test", makeInputs(5), 3, Mapfunc, Reducefunc)
//}

// 创建输入文件
// 根据指定的数量创建输入文件，返回创建好的文件名列表写入相应的数据
// count: 创建的文件数量

func makeInputs(num int) []string {
	var names []string
	var i = 0
	for f := 0; f < num; f++ {
		// 文件命名方式
		names = append(names, fmt.Sprintf("824-mrinput-%d.txt", f))
		// 创建文件
		file, err := os.Create(names[f])
		if err != nil {
			log.Fatalf("create input file [%v] failed. error:%v", file, err)
		}
		w := bufio.NewWriter(file)
		for i < (f+1)*(nNumber/num) {
			// 写入i到w
			fmt.Fprintf(w, "%d\n", i)
			i++
		}
		// 把buffer中的内容写入文件
		w.Flush()
		file.Close()
	}
	return names
}
