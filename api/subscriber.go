package api

import (
	"fmt"
	"net/http"

	"github.com/PushpenderIndia/Free_YouTube_Stats_Badge/internal"
)

func SubscriberCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Subscribe", internal.ChannelStats.SubscriberCount)
	fmt.Fprintf(w, s)
}
