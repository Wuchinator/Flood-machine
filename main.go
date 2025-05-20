
package main

import (
    //"flag"
    "fmt"
    "http-flooder/config"
    "http-flooder/flood"
    "http-flooder/stats"
    "time"
)

func main() {
    cfg := config.ParseFlags()

    fmt.Println("Starting flood")
    s := stats.NewStats()

    start := time.Now()
    flood.Start(cfg, s)
    elapsed := time.Since(start)

    fmt.Println("Flood completed")
    s.Report(elapsed)
}
