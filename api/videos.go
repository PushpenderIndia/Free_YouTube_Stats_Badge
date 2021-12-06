package api

import (
	"fmt"
	"net/http"

	"github.com/PushpenderIndia/Free_YouTube_Stats_Badge/internal"
)

func VideoCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Videos", internal.ChannelStats.VideoCount)
	fmt.Fprintf(w, s)
}
