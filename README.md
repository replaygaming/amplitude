# Amplitude HTTP API Integration [![GoDoc](https://godoc.org/github.com/replaygaming/gameanalytics?status.svg)](https://godoc.org/github.com/replaygaming/amplitude)

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
