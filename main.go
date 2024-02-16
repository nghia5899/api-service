package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func api(wg *sync.WaitGroup, dem *int, index int) int {
	//defer wg.Done()
	data := url.Values{}
	data.Set("debug", "true")
	data.Set("log", "[{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967964602,\"message\":\"normal/blue500/light\",\"event_info\":\"normal/blue500/light\",\"items\":[],\"event_namespace\":{\"page\":\"app\",\"component\":\"theme\",\"action\":\"launch\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":1,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967964808,\"tweet_id\":\"1757843982791643607\",\"items\":[],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"show\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":2,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967965289,\"tweet_id\":\"1757843982791643607\",\"event_initiator\":0,\"new_entries\":2,\"new_tweets\":2,\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"get_initial\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":3,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967965648,\"tweet_id\":\"1757843982791643607\",\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"action\":\"bottom\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":4,\"client_app_id\":\"3033300\"},{\"_category_\":\"client_event\",\"format_version\":2,\"triggered_on\":1707967966037,\"items\":[{\"item_type\":0,\"id\":\"1757843982791643607\",\"position\":1,\"sort_index\":\"7465528054063132200\",\"percent_screen_height_100k\":28444,\"author_id\":\"1736445965731811328\",\"in_reply_to_tweet_id\":\"1757801917298938350\",\"in_reply_to_author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":false,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"engagement_metrics\":{\"reply_count\":0,\"retweet_count\":0,\"favorite_count\":0,\"quote_count\":0}},{\"item_type\":0,\"id\":\"1757801917298938350\",\"position\":0,\"sort_index\":\"7465570119555837457\",\"percent_screen_height_100k\":69000,\"author_id\":\"1539308438030667777\",\"is_viewer_follows_tweet_author\":true,\"is_tweet_author_follows_viewer\":false,\"is_viewer_super_following_tweet_author\":false,\"is_viewer_super_followed_by_tweet_author\":false,\"is_tweet_author_super_followable\":false,\"media_details\":{\"photo_count\":1,\"content_id\":\"1757801911112400896\",\"publisher_id\":\"1539308438030667777\",\"media_type\":8,\"dynamic_ads\":false},\"media_details_v2\":[{\"content_id\":\"1757801911112400896\",\"publisher_id\":\"1539308438030667777\",\"media_type\":8,\"dynamic_ads\":false}],\"engagement_metrics\":{\"reply_count\":13,\"retweet_count\":7,\"favorite_count\":291,\"quote_count\":0}}],\"event_namespace\":{\"page\":\"tweet\",\"component\":\"stream\",\"action\":\"results\",\"client\":\"m5\"},\"client_event_sequence_start_timestamp\":1707967964529,\"client_event_sequence_number\":5,\"client_app_id\":\"3033300\"}]")
	request, _ := http.NewRequest(http.MethodPost, "https://twitter.com/i/api/1.1/jot/client_event.json", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "_ga=GA1.2.14795140.1700537829; kdt=LoS3Uj1VjCdKki4ZRIaYht53SQXzLifwL9yhHXIX; ads_prefs=\"HBERAAA=\"; guest_id_ads=v1%3A170565034985792655; guest_id_marketing=v1%3A170565034985792655; guest_id=v1%3A170565034985792655; g_state={\"i_l\":0}; auth_token=c33b0c3747352e5cc96f96d3b6904054849826eb; ct0=6da050330609d347000d0d53932e9a66ca49b8e0719cec0a7bd68a0d53f7ba7e27ff0b284115fbd7df8339324081faf5f21e01ca9c545d1b6e36ffb8157b24c6d0e520e3d6cc0f7253d256858fe675fb; twid=u%3D1750722368132308992; lang=en; _gid=GA1.2.1494649059.1707966529; personalization_id=\"v1_mYzLXec8rHpnE+9tfLEh0Q==\"; ct0=fd73f5d0898752892ca496d7b97508374f5cd00561d75891908bdbc97815466a9a56e4e770382b84ab1af31c203a35be2c896600bfc1bba0c086383a5d2cc0cc4564438342930576f18e81dc8321ba99; guest_id=v1%3A170537966547498546; guest_id_ads=v1%3A170537966547498546; guest_id_marketing=v1%3A170537966547498546; personalization_id=\"v1_AnMuDz7UZye9rg955pUVGQ==\"")
	request.Header.Add("X-Csrf-Token", "6da050330609d347000d0d53932e9a66ca49b8e0719cec0a7bd68a0d53f7ba7e27ff0b284115fbd7df8339324081faf5f21e01ca9c545d1b6e36ffb8157b24c6d0e520e3d6cc0f7253d256858fe675fb")
	request.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return 0
	}
	defer response.Body.Close()
	*dem += 1
	return response.StatusCode
}

func callManyApi(wgCall *sync.WaitGroup, wgResponse *sync.WaitGroup, indexNgoai int, dem *int) {
	num := 1000
	for i := 1; i < num; i++ {
		wgResponse.Add(1)
		go func(index int) {
			//fmt.Println("call ", index)
			status := api(wgResponse, dem, index)
			*dem++
			fmt.Println("dem ", *dem, " status", status)
			wgResponse.Done()
		}(num*indexNgoai + i)
	}
	fmt.Println("========================= call xong 1000 ============== = ", num*indexNgoai+num)
	//wg.Wait()
	//fmt.Println("========================= tra ve  xong 1000 ============== = ", 1000*indexNgoai+index)
	wgCall.Done()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	go callApiView()
}

func callApiView() {
	wgResponse := sync.WaitGroup{}
	wgCall := sync.WaitGroup{}
	index1 := 0
	dem := 0
	a := &dem
	fmt.Println("bat dau")
	for index1 < 500 {
		wgCall.Add(1)
		go callManyApi(&wgCall, &wgResponse, index1, a)
		fmt.Println("========================= het vong ngoai = ", index1)
		index1++
	}
	wgCall.Wait()
	wgResponse.Wait()
}

func main() {
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:80\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
