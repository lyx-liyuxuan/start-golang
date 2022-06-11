package mapreduce

type Master struct {
}

// 任务调度函数
func SequentialSingle(jobName string, files []string, nReduce int,
	mapf func(string, string) []KeyValue, reducef func(string, []string) string) {

	// 执行分配的任务，map任务和reduce任务
	mr := NewMaster()
	mr.run(jobName, files, nReduce, func(phase jobPhase) {
		switch phase {
		case mapPhase:
			// 执行map任务，调用次数由输入文件的个数决定
			for i, filename := range files {
				doMap(jobName, i, filename, nReduce, mapf)
			}
		case reducePhase:
			// 执行reduce任务，reduce调用次数由nReduce决定
			for i := 0; i < nReduce; i++ {
				doReduce(jobName, i, mergeName(jobName, i), len(files), reducef)
			}
		}
	})
}

// 初始化Master实例
func NewMaster() *Master {

	return nil
}

// 执行给定的任务
func (mr *Master) run(jobName string, files []string, nreduce int, schedule func(phase jobPhase)) {
	// 顺序执行map任务
	schedule(mapPhase)
	// 顺序执行reduce任务
	schedule(reducePhase)
}
