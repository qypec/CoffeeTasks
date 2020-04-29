package main

func ExecutePipeline(jobs ...job) {
	input := make(chan interface{}, 100)
	output := make(chan interface{}, 100)

	job1 := jobs[0]
	job2 := jobs[1]

	go job1(input, output)
	go job2(output, input)
}