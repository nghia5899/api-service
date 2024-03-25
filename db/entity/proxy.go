package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type Proxy struct {
	id       string
	Host     string `bson:"host"`
	UserName string `bson:"user_name"`
	PassWord string `bson:"pass_word"`
}

func ProxyFromMap(data map[string]string) Proxy {
	var proxy = Proxy{
		Host:     data["host"],
		UserName: data["user_name"],
		PassWord: data["pass_word"],
	}
	return proxy
}

func (e *Proxy) ToBson() (data bson.M) {
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
