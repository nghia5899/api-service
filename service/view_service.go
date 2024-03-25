package service

import (
	"apiService/db"
	"apiService/db/entity"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func View(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	var requestBody map[string]string
	json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)

	cookies := getCookie()
	if cookies == nil {
		return
	}
	go callApiView(cookies, requestBody["tweetId"], requestBody["publisherId"], requestBody["solan"])
}

func getCookie() []entity.Cookie {
	var cookies []entity.Cookie
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := db.OpenDB()
	if db == nil {
		return nil
	}
	cursor, err := db.Collection("cookie").Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var cookie entity.Cookie
		cursor.Decode(&cookie)
		cookies = append(cookies, cookie)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}

	return cookies
}

func api(cookie entity.Cookie, tweetId string, publisherId string) int {
	data := url.Values{}
	data.Set("debug", "true")
	data.Set("log", "[{  \"_category_\": \"client_event\",\n  \"format_version\": 2,\n  \"triggered_on\": 1707967966037,\n  \"items\": [{\n    \"item_type\": 0,\n    \"id\": \""+tweetId+"\",\n    \"position\": 0,\n    \"sort_index\": \"7465570119555837457\",\n    \"percent_screen_height_100k\": 69000,\n    \"author_id\": \""+publisherId+"\",\n    \"is_viewer_follows_tweet_author\": true,\n    \"is_tweet_author_follows_viewer\": false,\n    \"is_viewer_super_following_tweet_author\": false,\n    \"is_viewer_super_followed_by_tweet_author\": false,\n    \"is_tweet_author_super_followable\": false,\n    \"media_details\": {\n      \"photo_count\": 1,\n      \"content_id\": \""+tweetId+"\",\n      \"publisher_id\": \""+publisherId+"\",\n      \"media_type\": 8,\n      \"dynamic_ads\": false\n    },\n    \"media_details_v2\": [{\n      \"content_id\": \"1768118708667621422\",\n      \"publisher_id\": \"1750722368132308992\",\n      \"media_type\": 8,\n      \"dynamic_ads\": false\n    }],\n    \"engagement_metrics\": {\n      \"reply_count\": 13,\n      \"retweet_count\": 7,\n      \"favorite_count\": 291,\n      \"quote_count\": 0\n    }\n  }],\n  \"event_namespace\": {\n    \"page\": \"tweet\",\n    \"component\": \"stream\",\n    \"action\": \"results\",\n    \"client\": \"m5\"\n  },\n  \"client_event_sequence_start_timestamp\": 1707967964529,\n  \"client_event_sequence_number\": 5,\n  \"client_app_id\": \"3033300\"\n}]\n'")
	//data.Set("log", "[{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967964602,\"message\":\"normal/blue500/light\",\"event_info\":\"normal/blue500/light\",\"items\":[],\"event_namespace\":{\"page\":\"app\",\"component\":\"theme\",\"action\":\"launch\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":1,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967964808,\"tweet_id\":\""+tweetId+"\",\"items\":[],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"show\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":2,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967965289,\"tweet_id\":\"1757843982791643607\",\"event_initiator\":0,\"new_entries\":2,\"new_tweets\":2,\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"get_initial\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":3,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967965648,\"tweet_id\":\"1757843982791643607\",\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"bottom\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":4,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967966037,\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"position\":1,\"sort_index\":\"7465528054063132200\",\"percent_screen_height_100k\":28444,\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}},{\"item_type\":0,\"id\":\"1757801917298938350\",\"position\":0,\"sort_index\":\"7465570119555837457\",\"percent_screen_height_100k\":69000,\"author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":true,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"media_details\":{\"photo_count\":1,\"content_id\":\"1757801911112400896\",\"publisher_id\":\""+publisherId+"\",\"media_type\":8,\"dynamic_ads\":false},\"media_details_v2\":[{\"content_id\":\"1757801911112400896\",\"publisher_id\":\"1539308438030667777\",\"media_type\":8,\"dynamic_ads\":false}],\"engagement_metrics\":{\"reply_count\":13,\"retweet_count\":7,\"favorite_count\":291,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"component\":\"stream\",\"action\":\"results\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":5,\"client_app_id\":\"3033300\"}]")
	request, _ := http.NewRequest(http.MethodPost, "https://twitter.com/i/api/1.1/jot/client_event.json", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", cookie.Cookie)
	request.Header.Add("X-Csrf-Token", cookie.Ct0)
	request.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	//transport := http.Transport{
	//	Proxy: http.ProxyURL(&url.URL{
	//		Scheme: "http",
	//		Host:   "04.239.52.219:7381",
	//		User:   url.UserPassword("bnns2305", "bnns2305"),
	//	}),
	//}
	client := http.Client{Timeout: 50 * time.Second}
	//client.Transport = &transport
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer response.Body.Close()
	return response.StatusCode
}

func callManyApi(tweetId string, publisherId string, wgCall *sync.WaitGroup, wgResponse *sync.WaitGroup, indexNgoai int, cookie entity.Cookie, chanelSuccess chan bool, dem *int, tong int, solan int) {
	//so lan moi nick call
	num := solan
	for i := 1; i <= num; i++ {
		wgResponse.Add(1)
		go func(index int, cookie2 entity.Cookie, cSuccess chan bool, dem *int, tong int) {
			status := api(cookie2, tweetId, publisherId)
			fmt.Println(status)
			if status == 200 {
				cSuccess <- true
			}
			*dem++
			fmt.Println("dem  ", *dem)
			if *dem == tong {
				close(cSuccess)
			}
			wgResponse.Done()
		}(num*indexNgoai+i, cookie, chanelSuccess, dem, tong)
	}
	fmt.Println("==================== call xong: ", num*(indexNgoai+1), "lan ============")
	wgCall.Done()
}

func callApiView(cookies []entity.Cookie, tweetId string, publisherId string, solan string) {
	wgResponse := sync.WaitGroup{}
	wgCall := sync.WaitGroup{}
	soLanCalLMoiNick, _ := strconv.Atoi(solan)
	tong := len(cookies) * soLanCalLMoiNick
	demThanhCong := 0
	demTong := 0
	var dem = &demTong
	chanelSuccess := make(chan bool)
	for index, cookie := range cookies {
		wgCall.Add(1)
		go callManyApi(tweetId, publisherId, &wgCall, &wgResponse, index, cookie, chanelSuccess, dem, tong, soLanCalLMoiNick)
	}
	for range chanelSuccess {
		fmt.Println("chanelSuccess")
		demThanhCong++
	}
	wgCall.Wait()
	wgResponse.Wait()
	fmt.Println("====== Xong ======")
	fmt.Println("Dem Tong: ", demTong)
	fmt.Println("Dem Thanh Cong: ", demThanhCong)
}
