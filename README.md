# [Amplitude](http://amplitude.com) HTTP API Client 
[![GoDoc](https://godoc.org/github.com/replaygaming/gameanalytics?status.svg)](https://godoc.org/github.com/replaygaming/amplitude) [![Build Status](https://semaphoreci.com/api/v1/projects/695adf86-d24c-4a13-aa7f-1b4d3bf9e342/546999/shields_badge.svg)](https://semaphoreci.com/luizbranco/amplitude)[![Coverage Status](https://coveralls.io/repos/replaygaming/amplitude/badge.svg?branch=master&service=github)](https://coveralls.io/github/replaygaming/amplitude?branch=master)

## APIs supported:

  - [x] Events API (httpapi)
  - [x] Identify API

## Basic Event Example:

```go

import (
  "fmt"

  "github.com/replaygaming/amplitude"
)

func main() {
  apiKey := "abcdef"
  c := amplitude.NewClient(apiKey)
  e := amplitude.Event{EventType: "test", UserID: "1"}
  if _, err := c.Send(e); err != nil {
    fmt.Println(err)
  }
}
```

## Basic Identify Example:

```go

import (
  "fmt"

  "github.com/replaygaming/amplitude"
)

func main() {
  apiKey := "abcdef"
  c := amplitude.NewClient(apiKey)
  i := amplitude.Identification{UserID: "1", Platform: "Browser"}
  if _, err := c.Send(i); err != nil {
    fmt.Println(err)
  }
}
```
