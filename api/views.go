package api

import (
	"fmt"
	"net/http"

	"github.com/PushpenderIndia/Free_YouTube_Stats_Badge/internal"
)

func ViewCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Views", internal.ChannelStats.ViewsCount)
	fmt.Fprintf(w, s)
}
