package builder

// This file implements a job runner for the compiler, which runs jobs in
// parallel while taking care of dependencies.

import (
	"fmt"
	"runtime"
	"time"
)

// Set to true to enable logging in the job runner. This may help to debug
// concurrency or performance issues.
const jobRunnerDebug = false

type jobState uint8

const (
	jobStateQueued   jobState = iota // not yet running
	jobStateRunning                  // running
	jobStateFinished                 // finished running
)

// compileJob is a single compiler job, comparable to a single Makefile target.
// It is used to orchestrate various compiler tasks that can be run in parallel
// but that have dependencies and thus have limitations in how they can be run.
type compileJob struct {
	description  string // description, only used for logging
	dependencies []*compileJob
	run          func() error
	state        jobState
	err          error         // error if finished
	duration     time.Duration // how long it took to run this job (only set after finishing)
}

// readyToRun returns whether this job is ready to run: it is itself not yet
// started and all dependencies are finished.
func (job *compileJob) readyToRun() bool {
	if job.state != jobStateQueued {
		// Already running or finished, so shouldn't be run again.
		return false
	}

	// Check dependencies.
	for _, dep := range job.dependencies {
		if dep.state != jobStateFinished {
			// A dependency is not finished, so this job has to wait until it
			// is.
			return false
		}
	}

	// All conditions are satisfied.
	return true
}

// runJobs runs all the jobs indicated in the jobs slice and returns the error
// of the first job that fails to run.
// It runs all jobs in the order of the slice, as long as all dependencies have
// already run. Therefore, if some jobs are preferred to run before others, they
// should be ordered as such in this slice.
func runJobs(jobs []*compileJob) error {
	// Create channels to communicate with the workers.
	doneChan := make(chan *compileJob)
	workerChan := make(chan *compileJob)
	defer close(workerChan)

	// Start a number of workers.
	for i := 0; i < runtime.NumCPU(); i++ {
		if jobRunnerDebug {
			fmt.Println("## starting worker", i)
		}
		go jobWorker(workerChan, doneChan)
	}

	// Send each job in the jobs slice to a worker, taking care of job
	// dependencies.
	numRunningJobs := 0
	var totalTime time.Duration
	start := time.Now()
	for {
		// If there are free workers, try starting a new job (if one is
		// available). If it succeeds, try again to fill the entire worker pool.
		if numRunningJobs < runtime.NumCPU() {
			jobToRun := nextJob(jobs)
			if jobToRun != nil {
				// Start job.
				if jobRunnerDebug {
					fmt.Println("## start:   ", jobToRun.description)
				}
				jobToRun.state = jobStateRunning
				workerChan <- jobToRun
				numRunningJobs++
				continue
			}
		}

		// When there are no jobs running, all jobs in the jobs slice must have
		// been finished. Therefore, the work is done.
		if numRunningJobs == 0 {
			break
		}

		// Wait until a job is finished.
		job := <-doneChan
		job.state = jobStateFinished
		numRunningJobs--
		totalTime += job.duration
		if jobRunnerDebug {
			fmt.Println("## finished:", job.description, "(time "+job.duration.String()+")")
		}
		if job.err != nil {
			// Wait for running jobs to finish.
			for numRunningJobs != 0 {
				<-doneChan
				numRunningJobs--
			}
			// Return error of first failing job.
			return job.err
		}
	}

	// Some statistics, if debugging.
	if jobRunnerDebug {
		// Total duration of running all jobs.
		duration := time.Since(start)
		fmt.Println("## total:   ", duration)

		// The individual time of each job combined. On a multicore system, this
		// should be lower than the total above.
		fmt.Println("## job sum: ", totalTime)
	}

	return nil
}

// nextJob returns the first ready-to-run job.
// This is an implementation detail of runJobs.
func nextJob(jobs []*compileJob) *compileJob {
	for _, job := range jobs {
		if job.readyToRun() {
			return job
		}
	}
	return nil
}

// jobWorker is the goroutine that runs received jobs.
// This is an implementation detail of runJobs.
func jobWorker(workerChan, doneChan chan *compileJob) {
	for job := range workerChan {
		start := time.Now()
		err := job.run()
		if err != nil {
			job.err = err
		}
		job.duration = time.Since(start)
		doneChan <- job
	}
}
