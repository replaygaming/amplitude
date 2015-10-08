// Package amplitude is a simple client implementation of the Amplitude HTTP
// API.
//
// Basic usage example:
//
//		import (
//			"fmt"
//
//			"github.com/replaygaming/amplitude"
//		)
//
//		func main() {
//			apiKey := "abcdef"
//			c := amplitude.NewClient(apiKey)
//			e := amplitude.Event{EventType: "test", UserID: "1"}
//			if _, err := c.Send(e); err != nil {
//				fmt.Println(err)
//			}
//		}
//
package amplitude
