package utils

//import "fmt"

type EventType string

const (
	Play    EventType = "play"
	Playing EventType = "playing"
	Pause   EventType = "pause"
	Resume   EventType = "resume"
	Start   EventType = "start"
	End   EventType = "end"
	Stop   EventType = "stop"
	Cancel   EventType = "cancel"
)

type Topic struct {
    ViewId string `json:"veiwId"`
    PlaybackId string `json:"playbackId"`
    EventType EventType `json:"EventType"`
    ViewStartTime string `json:"viewStartTime"`
    VideoStartTime int64 `json:"videoStartTime"`
    CountryId int `json:"countryId"`
    UserId int `json:"userId"`
}

// type EventSchema struct{
//     Events []Topic `json:"events"`
// }

// func main() {
// 	// Create a slice of events
// 	eventList := EventSchema{
// 		Events: []Topic{
// 			{ViewId: "view1", PlaybackId: "playback1", EventType: "Play"},
// 			{ViewId: "view2", PlaybackId: "playback2", EventType: "Playing"},
// 			{ViewId: "view3", PlaybackId: "playback3", EventType: "Pause"},
// 			{ViewId: "view3", PlaybackId: "playback3", EventType: "anything"},
// 		},
// 	}

// 	// Print the event list
// 	fmt.Println(eventList)
// }