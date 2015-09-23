// Package amplitude is a proof-of-concept integration with
// [Amplitude](http://amplitude.com) HTTP API
//
// Usage example:
//
//		import (
//			"fmt"
//
//			"github.com/replaygaming/amplitude"
//		)
//
//		func main() {
//			apiKey := "abcdef"
//			s := amplitude.NewClient(apiKey)
//			e := amplitude.Event{EventType: "test", UserID: "1"}
//			if _, err := s.SendEvent(e); err != nil {
//				fmt.Println(err)
//			}
//		}
//
package amplitude
