# Amplitude HTTP API Integration [![GoDoc](https://godoc.org/github.com/replaygaming/gameanalytics?status.svg)](https://godoc.org/github.com/replaygaming/amplitude) [![Build Status](https://semaphoreci.com/api/v1/projects/695adf86-d24c-4a13-aa7f-1b4d3bf9e342/546999/badge.svg)](https://semaphoreci.com/luizbranco/amplitude)

## Usage example:

```go

import (
  "fmt"

  "github.com/replaygaming/amplitude"
)

func main() {
  apiKey := "abcdef"
  s := amplitude.NewServer(apiKey)
  e := amplitude.Event{EventType: "test", UserID: "1"}
  if _, err := s.SendEvent(e); err != nil {
    fmt.Println(err)
  }
}
```
