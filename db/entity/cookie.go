package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type Cookie struct {
	id                 string
	Cookie             string `bson:"cookie"`
	Twitter_sess       string `bson:"twitter_sess"`
	Twid               string `bson:"twid"`
	Dnt                string `bson:"dnt"`
	Ct0                string `bson:"ct0"`
	Lang               string `bson:"lang"`
	Gt                 string `bson:"gt"`
	Guest_id           string `bson:"guest_id"`
	Kdt                string `bson:"kdt"`
	Auth_token         string `bson:"auth_token"`
	Guest_id_marketing string `bson:"guest_id_marketing"`
	Ga                 string `bson:"ga"`
	Gid                string `bson:"gid"`
	Guest_id_ads       string `bson:"guest_id_ads"`
}

func CookieFromMap(data map[string]string) Cookie {
	var cookie = Cookie{
		Cookie:             data["cookie"],
		Twitter_sess:       data["_twitter_sess"],
		Twid:               data["twid"],
		Dnt:                data["dnt"],
		Ct0:                data["ct0"],
		Lang:               data["lang"],
		Gt:                 data["gt"],
		Guest_id:           data["guest_id"],
		Kdt:                data["kdt"],
		Auth_token:         data["auth_token"],
		Guest_id_marketing: data["guest_id_marketing"],
		Ga:                 data["_ga"],
		Gid:                data["_gid"],
		Guest_id_ads:       data["guest_id_ads"],
	}
	return cookie
}

func (e *Cookie) ToBson() (data bson.M) {
	var tagValue string

	data = bson.M{}
	element := reflect.ValueOf(e).Elem()

	for i := 0; i < element.NumField(); i += 1 {
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		tagValue = tag.Get("bson")

		if tagValue == "-" {
			continue
		}

		switch element.Field(i).Kind() {
		case reflect.String:
			value := element.Field(i).String()
			data[tagValue] = value

		case reflect.Bool:
			value := element.Field(i).Bool()
			data[tagValue] = value

		case reflect.Int:
			value := element.Field(i).Int()
			data[tagValue] = value
		}
	}

	return
}
