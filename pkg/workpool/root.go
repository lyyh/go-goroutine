package workpool

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func Workpool() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	createWorkPool(jobChan, resultChan, 64)
	printer(resultChan)
	id := 0
	// 循环创建job，输入到管道
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

func createWorkPool(jobChan chan *Job, resultChan chan *Result, num int) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				sum := 0
				r_num := job.RandNum
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				result := &Result{
					job: job,
					sum: sum,
				}

				resultChan <- result
			}
		}(jobChan, resultChan)
	}
}

func printer(resultChan chan *Result) {
	fmt.Println("aaa")
	go func() {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}()
}
