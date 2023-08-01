package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"strings"
)

type Response struct {
	CityName                string `json:"CityName"`
	CountryCode             string `json:"CountryCode"`
	CountryName             string `json:"CountryName"`
	CityLName               string `json:"CityLName"`
	CountryLName            string `json:"CountryLName"`
	CountryAlpha2           string `json:"CountryAlpha2"`
	TimeZone                string `json:"TimeZone"`
	Imsaak                  string `json:"Imsaak"`
	Sunrise                 string `json:"Sunrise"`
	SunriseDT               string `json:"SunriseDT"`
	Noon                    string `json:"Noon"`
	Sunset                  string `json:"Sunset"`
	Maghreb                 string `json:"Maghreb"`
	Midnight                string `json:"Midnight"`
	Today                   string `json:"Today"`
	TodayQamari             string `json:"TodayQamari"`
	TodayGregorian          any    `json:"TodayGregorian"`
	DayLenght               any    `json:"DayLenght"`
	SimultaneityOfKaaba     string `json:"SimultaneityOfKaaba"`
	SimultaneityOfKaabaDesc string `json:"SimultaneityOfKaabaDesc"`
}

func main (){
	const url string = "https://prayer.aviny.com/api/prayertimes/11"
	
	resp, err := http.Get(url)
	CatchErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	//fmt.Printf("%+v\n", result)

	loc, _ := time.LoadLocation("Asia/Tehran")
	now := time.Now()
	year, month, day := now.Date()
	var sobh, zohr, magh time.Time

	sobh, _ = time.Parse("15:04:05", result.Imsaak)
	zohr, _ = time.Parse("15:04:05", result.Noon)
	magh, _ = time.Parse("15:04:05", result.Maghreb)

	sobh = time.Date(year, month, day, sobh.Hour(), sobh.Minute(), sobh.Second(), 0, loc)
	zohr = time.Date(year, month, day, zohr.Hour(), zohr.Minute(), zohr.Second(), 0, loc)
	magh = time.Date(year, month, day, magh.Hour(), magh.Minute(), magh.Second(), 0, loc)

	/*
	fmt.Println("now: ", now)
	fmt.Println("sobh: ", sobh)
	fmt.Println("zohr: ", zohr)
	fmt.Println("magh: ", magh)
	*/
	
	var duration time.Duration
	if now.Before(sobh) {
		duration = sobh.Sub(now)
	}else if now.Before(zohr) {
		duration = zohr.Sub(now)
	}else if now.Before(magh){
		duration = magh.Sub(now)
	}else {
		duration = sobh.AddDate(0,0,1).Sub(now)
	}

	fmt.Println(strings.TrimRight(duration.Truncate(time.Minute).String(), "0s") + "inutes")

}


func CatchErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
