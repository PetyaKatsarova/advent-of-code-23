package main

import (
    "fmt"
    "sync"
	"time"
)

var (
    counter int
    mutex   sync.Mutex
)

// func increment() {
//     mutex.Lock()   // Lock the Mutex before modifying shared data
//     counter++
//     mutex.Unlock() // Unlock the Mutex after modifying shared data
// }

// func main() {
//     var wg sync.WaitGroup
//     for i := 0; i < 1000; i++ {
//         wg.Add(1)
//         go func() {
//             defer wg.Done()
//             increment()
//         }()
//     }
//     wg.Wait()
//     fmt.Println("Counter:", counter)
// }
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 9; a++ {
        <-results
    }
}


