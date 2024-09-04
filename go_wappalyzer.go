package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: wappalyzergo-wrapper <url>")
    }

    url := os.Args[1]
    resp, err := http.DefaultClient.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    wappalyzerClient, err := wappalyzer.New()
    if err != nil {
        log.Fatal(err)
    }

    fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
    jsonOutput, err := json.Marshal(fingerprints)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonOutput))
}
