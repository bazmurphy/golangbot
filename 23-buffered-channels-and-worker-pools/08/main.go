package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Each `Job` struct has a `id` and a `randomno` for which the sum of the individual digits has to be computed.
type Job struct {
	id       int
	randomno int
}

// The `Result` struct has a `job` field which is the job for which it holds the result (sum of individual digits) in the `sumofdigits` field.
type Result struct {
	job         Job
	sumofdigits int
}

// Worker Goroutines listen for new tasks on the `jobs` buffered channel.
var jobs = make(chan Job, 10)

// Once a task is complete, the result is written to the `results` buffered channel.
var results = make(chan Result, 10)

// The `digits` function does the actual job of finding the sum of the individual digits of an integer and returning it.
// We will add a sleep of 2 seconds to this function just to simulate the fact that it takes some time for this function to calculate the result.
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// This function creates a worker which reads from the `jobs` channel, creates a `Result` struct using the current `job` and the return value of the `digits` function and then writes the result to the `results` buffered channel.
// This function takes a WaitGroup `wg` as a parameter on which it will call the `Done()` method when all `jobs` have been completed.
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// This function takes the number of workers to be created as a parameter.
// It calls `wg.Add(1)` before creating the Goroutine to increment the WaitGroup counter.
// Then it creates the worker Goroutines by passing the pointer of the WaitGroup `wg` to the `worker` function.
// After creating the needed worker Goroutines, it waits for all the Goroutines to finish their execution by calling `wg.Wait()`.
// After all Goroutines finish executing, it closes the `results` channel since all Goroutines have finished their execution and no one else will further be writing to the `results` channel.
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// The `allocate` function above takes the number of jobs to be created as input parameter, generates pseudo random numbers with a maximum value of `998`, creates `Job` struct using the random number and the for loop counter `i` as the id and then writes them to the `jobs` channel.
// It closes the `jobs` channel after writing all jobs.
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// The `result` function reads the `results` channel and prints the job id, input random no, and the sum of digits of the random no.
// The result function also takes a `done` channel as a parameter to which it writes to once it has printed all the results.
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	// We first store the execution start time of the program in line no.2 of the main function and in the last line (line no. 12) we calculate the time difference between the endTime and startTime and display the total time it took for the program to run. This is needed because we will do some benchmarks by changing the number of Goroutines.
	startTime := time.Now()

	// The `noOfJobs` is set to 100 and then `allocate` is called to add jobs to the `jobs` channel.
	noOfJobs := 100
	go allocate(noOfJobs)

	// Then `done` channel is created and passed to the `result` Goroutine so that it can start printing the output and notify once everything has been printed.
	done := make(chan bool)
	go result(done)

	// Finally a pool of `10` worker Goroutines are created by the call to `createWorkerPool` function and then main waits on the `done` channel for all the results to be printed.
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
