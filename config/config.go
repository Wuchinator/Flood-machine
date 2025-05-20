
package config

import (
    "flag"
    "time"
)

type Config struct {
    URL              string
    Threads          int
    RequestsPerThread int
    Timeout          time.Duration
    Delay            time.Duration
}

func ParseFlags() *Config {
    url := flag.String("url", "http://localhost:8080", "Target URL")
    threads := flag.Int("threads", 10, "Number of concurrent workers")
    requests := flag.Int("rpt", 100, "Requests per thread")
    timeout := flag.Duration("timeout", 5*time.Second, "HTTP request timeout")
    delay := flag.Duration("delay", 0, "Delay between requests (per worker)")

    flag.Parse()

    return &Config{
        URL:              *url,
        Threads:          *threads,
        RequestsPerThread: *requests,
        Timeout:          *timeout,
        Delay:            *delay,
    }
}
