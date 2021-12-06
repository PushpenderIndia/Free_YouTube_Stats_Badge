package internal

import (
	"encoding/json"
	"strconv"
	"math"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


func RoundPrec(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	if math.IsNaN(number) || math.IsInf(number, 0) {
		number = 0
	}

	var ret string
	var negative bool

	if number < 0 {
		number *= -1
		negative = true
	}

	d, fract := math.Modf(number)

	if decimals <= 0 {
		fract = 0
	} else {
		pow := math.Pow(10, float64(decimals))
		fract = RoundPrec(fract*pow, 0)
	}

	if thousandsSep == "" {
		ret = strconv.FormatFloat(d, 'f', 0, 64)
	} else if d >= 1 {
		var x float64
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = strconv.FormatFloat(x, 'f', 0, 64) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	} else {
		ret = "0"
	}

	fracts := strconv.FormatFloat(fract, 'f', 0, 64)

	// "0" pad left
	for i := len(fracts); i < decimals; i++ {
		fracts = "0" + fracts
	}

	ret += decPoint + fracts

	if negative {
		ret = "-" + ret
	}
	return ret
}

func RoundInt(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func FormatNumber(input float64) string {
	x := RoundInt(input)
	xFormatted := NumberFormat(float64(x), 2, ".", ",")
	return xFormatted
}

func NearestThousandFormat(num float64) string {
	if math.Abs(num) < 999.5 {
		xNum := FormatNumber(num)
		xNumStr := xNum[:len(xNum)-3]
		return string(xNumStr)
	}

	xNum := FormatNumber(num)
	// first, remove the .00 then convert to slice
	xNumStr := xNum[:len(xNum)-3]
	xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
	xNumSlice := strings.Fields(xNumCleaned)
	count := len(xNumSlice) - 2
	unit := [4]string{"K", "M", "B", "T"}
	xPart := unit[count]

	afterDecimal := ""
	if xNumSlice[1][0] != 0 {
		afterDecimal = "." + string(xNumSlice[1][0])
	}
	final := xNumSlice[0] + afterDecimal + xPart
	return final
}


type AllData struct {
	SubscriberCount     int `json:"est_sub"`
	TableList []Table  `json:"table"`
}

type Table struct {
	Count int `json:"count"`
}

type Statistics struct {
	SubscriberCount    string `json:"est_sub"`
	ViewsCount          string 
	VideoCount         string   
	Time            time.Time
}

var ChannelStats Statistics

func UpdateCounter() {
	reqString := fmt.Sprintf("https://socialcounts.org/api/counter/youtube-live-subscriber-count/%s", Info.ChannelID)
	// fmt.Println("[INFO] [URL: ", reqString, "]")
	// Response e.g. 
	// {"est_sub":10600,
	//  "API_sub":10600,
	//  "users":{"total":87,"you":1,"thisPage":1},
	//  "table":[
	// 	{"name":"Channel Views","count":385066},
	// 	{"name":"Videos","count":26}]
	// }
	
	now := time.Now()
	if now.After(ChannelStats.Time.Add(time.Second * time.Duration(Info.CacheTime))) {   // Cache Time to Prevent Rate Limit
		fmt.Println("[INFO] : Not Fresh")   // Printing that, Data is Not Fresh so requesting Fresh Data from API
		resp, err := http.Get(reqString)  // Getting Response & Error
		if err != nil {					  // If Error != None, then Throw Error
			panic(err)
		}
		defer resp.Body.Close()			         // Closing Response Body
		body, err := ioutil.ReadAll(resp.Body)   // Reading HTTP Response + HTTP Response Error


		var subscriber AllData
		json.Unmarshal(body, &subscriber)

		subscount_str := NearestThousandFormat(float64(subscriber.SubscriberCount))

		// int to string (Type Casting)
		ViewsCount    := strconv.Itoa(subscriber.TableList[0].Count)
		VideoCount    := strconv.Itoa(subscriber.TableList[1].Count)	
		
		var channel Statistics
		channel.SubscriberCount = subscount_str
		channel.ViewsCount      = ViewsCount
		channel.VideoCount      = VideoCount
		channel.Time            = time.Now()
		fmt.Println("[DATA] [SubscriberCount: ", subscount_str, "]")
		fmt.Println("[DATA] [ViewsCount     : ", ViewsCount, "]")
		fmt.Println("[DATA] [VideosCount    : ", VideoCount, "]")
		fmt.Println("[DATA] [Time           : ", time.Now(), "]")
		ChannelStats = channel

	} else {
		fmt.Println("[INFO] : Already fresh")
	}
}
