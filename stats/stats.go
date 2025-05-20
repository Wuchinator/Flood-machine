
package stats

import (
    "fmt"
    "sync"
    "time"
    "log"
)

type Stats struct {
    mu          sync.Mutex
    total       int
    totalTime   time.Duration
    errors      int
}

func NewStats() *Stats {
    return &Stats{}
}

func (s *Stats) Record(duration time.Duration) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.total++
    s.totalTime += duration
}

func (s *Stats) IncrementErrors() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.errors++
}

func (s *Stats) Report(elapsed time.Duration) {
    s.mu.Lock()
    defer s.mu.Unlock()

    avg := time.Duration(0)
    if s.total > 0 {
        avg = s.totalTime / time.Duration(s.total)
    }
    if avg <= 0 {
        log.Fatal("\nНе удалось рассчитать среднее время ответа или подключиться к серверу\nПроверьте подключение и повторите попытку\n")
    }

    fmt.Printf("Total Requests: %d\n", s.total)
    fmt.Printf("Errors: %d\n", s.errors)
    fmt.Printf("Average Response Time: %s\n", avg)
    fmt.Printf("Test Duration: %s\n", elapsed)
}
