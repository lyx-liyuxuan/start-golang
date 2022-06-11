package mapreduce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// 实现一个map任务管理函数,从input files中读取内容
// 将输出分成指定数量的中间文件
// 自定义分割标准
func doMap(jobName string, mapTaskNumber int, inFile string,
	nReduce int, mapf func(file string, contents string) []KeyValue) {
	// 打开文件
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("open file %s failed. err:%v\n", inFile, err)
	}
	defer f.Close()
	// 从输入文件inFile中读取内容
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("read the content of file failed! %v\n", err)
	}
	//通过调用mapF对内容进行处理，分割map任务输出将指定文件的内容解析为key-value
	kvs := mapf(inFile, string(content))
	// 生成一个编码对象
	// 每个map任务生成nReduce个中间文件对象
	encoders := make([]*json.Encoder, nReduce)
	for i := 0; i < nReduce; i++ {
		file_name := reduceName(jobName, mapTaskNumber, i)
		f, err := os.Create(file_name)
		if err != nil {
			log.Fatalf("Unable to create file [%s]: %v\n", file_name, err)
		}
		defer f.Close()
		encoders[i] = json.NewEncoder(f)
	}

	// 将kvs生成的内容存入前面生成的中间文件中去
	for _, v := range kvs {
		// 对key值进行分类， 用编号的哈希值对nReduce取余进行分类
		index := ihash(v.Key) % nReduce
		err := encoders[index].Encode(&v)
		if err != nil {
			log.Fatalf("Unable to write file\n")
		}
	}
}
