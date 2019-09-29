package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	graph   map[string]*Job
	workers [5]*Worker
	//workers [2]*Worker // for testinput
)

type Job struct {
	Label     string
	Deps      []*Job
	Worker    int
	StartedAt int
	Done      bool
}

func (j *Job) Duration() int {
	//return int([]byte(j.Label)[0]) - 65 // for testinput
	return 60 + int([]byte(j.Label)[0]) - 65
}

func (j *Job) Ready() bool {
	return len(j.Deps) == 0
}

func (j *Job) DeleteDependent() {
	for k, c := range graph {
		del := -1
		for i := 0; i < len(c.Deps); i++ {
			if c.Deps[i].Label == j.Label {
				del = i
				break
			}
		}
		if del != -1 {
			graph[k].Deps = append(graph[k].Deps[:del], graph[k].Deps[del+1:]...)
		}
	}
}

func (j *Job) Solve() {
	delete(graph, j.Label)
	j.DeleteDependent()
}

type Worker struct {
	ID           int
	CurrentJob   *Job
	TimeToFinish int
}

func (w *Worker) Assign(t int, j *Job) {
	w.CurrentJob = j
	w.TimeToFinish = t + j.Duration()
	j.StartedAt = t
	j.Worker = w.ID
}

func (w *Worker) Available(t int) bool {
	return w.CurrentJob == nil
}

func (w *Worker) Complete(t int) {
	if w.TimeToFinish == t {
		w.CurrentJob.Solve()
		w.CurrentJob = nil
		w.TimeToFinish = -1
	}
}

func main() {
	graph = make(map[string]*Job)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			var x, y string
			n, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &x, &y)
			if err != nil || n != 2 {
				log.Fatal(err)
			}

			var dep *Job

			if graph[x] == nil {
				dep = &Job{
					Label:     x,
					Deps:      make([]*Job, 0),
					Done:      false,
					Worker:    -1,
					StartedAt: -1,
				}
				graph[x] = dep
			} else {
				dep = graph[x]
			}

			if graph[y] == nil {
				graph[y] = &Job{
					Label:     y,
					Deps:      make([]*Job, 0),
					Done:      false,
					Worker:    -1,
					StartedAt: -1,
				}
			}

			graph[y].Deps = append(graph[y].Deps, dep)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	workers = [...]*Worker{
		{0, nil, -1},
		{1, nil, -1},
		{2, nil, -1},
		{3, nil, -1},
		{4, nil, -1},
	}

	for t := 0; ; t++ {
		if len(graph) == 0 {
			fmt.Println(t)
			os.Exit(0)
		}

		jobs := getReadyJobs(t)

		for _, j := range jobs {
			if j.Worker == -1 {
				w := getAvailableWorker(t)
				if w != nil {
					w.Assign(t, j)
				}
			}
		}

		//dumpTime(t)

		for _, w := range workers {
			w.Complete(t)
		}

	}
}

func getAvailableWorker(t int) *Worker {
	for i := 0; i < len(workers); i++ {
		if workers[i].Available(t) {
			return workers[i]
		}
	}
	return nil
}

func getReadyJobs(t int) []*Job {
	ready := make([]*Job, 0)
	for _, c := range graph {
		if c.Ready() && c.StartedAt < t {
			ready = append(ready, c)
		}
	}
	sort.Slice(ready, func(i, j int) bool {
		return ready[i].Label < ready[j].Label
	})
	return ready
}

func dumpTime(t int) {
	fmt.Printf("%d\t", t)
	for _, w := range workers {
		if w.CurrentJob == nil {
			fmt.Printf(".\t")
		} else {
			fmt.Printf("%s\t", w.CurrentJob.Label)
		}
	}
	fmt.Println()
}
