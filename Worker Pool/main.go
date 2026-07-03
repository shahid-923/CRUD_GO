package main

import (
	"fmt"
	"sync"
	"time"
)

// Job structure
type Job struct {
	ID int
}

// Result structure
type Result struct {
	JobID int
	Msg   string
}

// Worker function
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "started job", job.ID)

		time.Sleep(1 * time.Second) // simulate work

		results <- Result{
			JobID: job.ID,
			Msg:   fmt.Sprintf("done by worker %d", id),
		}
	}
}

func main() {
	numWorkers := 3
	numJobs := 6

	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j}
		}
		close(jobs)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results
	for res := range results {
	fmt.Printf("Job %d was %s\n", res.JobID, res.Msg)
}

}
