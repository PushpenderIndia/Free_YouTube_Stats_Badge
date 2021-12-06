package main

import (
	"net/http"

	"github.com/PushpenderIndia/Free_YouTube_Stats_Badge/api"
)

func main() {
	http.HandleFunc("/subscribers", api.SubscriberCount)
	http.HandleFunc("/views", api.ViewCount)
	http.HandleFunc("/videos", api.VideoCount)

	http.ListenAndServe(":8090", nil)
}
