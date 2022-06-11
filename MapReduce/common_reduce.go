package mapreduce

import (
	"encoding/json"
	"log"
	"os"
)

// 管理reduce任务
func doReduce(jobName string, reduceTaskNumber int, outFile string,
	nMap int, reducef func(key string, value []string) string) {
	// 打开每个中间文件，合并相同key的内容，生成最终文件
	var result map[string][]string = make(map[string][]string)
	for i := 0; i < nMap; i++ {
		interFile := reduceName(jobName, i, reduceTaskNumber)
		f, err := os.Open(interFile)
		if err != nil {
			log.Fatalf("read content from file [%s],err:%v", interFile, err)
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		var kv KeyValue
		for decoder.More() {
			err := decoder.Decode(&kv)
			if err != nil {
				log.Fatalf("Unable to write file\n")
			}

			//相同key合并内容
			result[kv.Key] = append(result[kv.Key], kv.Value)
		}
	}

	var keys []string
	for key, _ := range result {
		keys = append(keys, key)
	}
	// 新建输出文件
	out_file, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("create outfile failed.err:%v", outFile)
	}
	defer out_file.Close()
	encoder := json.NewEncoder(out_file)
	for _, key := range keys {
		encoder.Encode(KeyValue{key, reducef(key, result[key])})
	}
}
