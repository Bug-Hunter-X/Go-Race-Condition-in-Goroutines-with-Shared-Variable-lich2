```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        const numWorkers = 5

        for i := 0; i < numWorkers; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        count++
                        fmt.Printf("Worker %d: count = %d\n", i, count)
                }(i)
        }

        wg.Wait()
        fmt.Printf("Final count: %d\n", count)
}
```