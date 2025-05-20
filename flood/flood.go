
package flood

import (
    "net/http"
    "sync"
    "time"
    "http-flooder/config"
    "http-flooder/stats"
)

func worker(cfg *config.Config, s *stats.Stats, wg *sync.WaitGroup) {
    defer wg.Done()
    client := &http.Client{
        Timeout: cfg.Timeout,
    }

    for i := 0; i < cfg.RequestsPerThread; i++ {
        start := time.Now()
        resp, err := client.Get(cfg.URL)
        duration := time.Since(start)

        if err != nil {
            s.IncrementErrors()
            continue
        }

        s.Record(duration)
        resp.Body.Close()

        time.Sleep(cfg.Delay)
    }
}

func Start(cfg *config.Config, s *stats.Stats) {
    var wg sync.WaitGroup
    for i := 0; i < cfg.Threads; i++ {
        wg.Add(1)
        go worker(cfg, s, &wg)
    }
    wg.Wait()
}
