package main

import (
	"fmt"
	"time"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	var j int
	for {
		j++
		jobs <- j
	}
	close(jobs)

	for {
		<-results
	}

	// for j := 1; j <= 90; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// for a := 1; a <= 90; a++ {
	// 	<-results
	// }
	// c := cron.New()
	// c.AddFunc("0 * * * * *", func() { fmt.Println("Every hour on the half hour") })

	// go c.Start()
	// sig := make(chan os.Signal)
	// signal.Notify(sig, os.Interrupt, os.Kill)
	// <-sig
	// Do jobs with params
	// gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	// gocron.Every(1).Second().Do(task)
	// gocron.Every(5).Seconds().Do(task)
	// gocron.Every(1).Minute().Do(task)
	// gocron.Every(1).Minutes().Do(task)
	// gocron.Every(1).Hour().Do(task)
	// gocron.Every(2).Hours().Do(task)
	// gocron.Every(1).Day().Do(task)
	// gocron.Every(2).Days().Do(task)

	// Do jobs on specific weekday
	// gocron.Every(1).Monday().Do(task)
	// gocron.Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	// gocron.Every(1).Day().At("10:30").Do(task)
	// gocron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	// _, time := gocron.NextRun()
	// fmt.Println(time)

	// gocron.Remove(task)
	// gocron.Clear()

	// function Start start all the pending jobs
	// <-gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(task)
	// <-s.Start()
}
