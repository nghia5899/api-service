package service

import (
	"apiService/db"
	"apiService/db/entity"
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const MAX_UPLOAD_SIZE = 1024 * 1024

func Upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	io.WriteString(w, "Upload files\n")
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err) //dont do this
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create("./uploads/cookie.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookies, err := ReadFile()
	if err != nil {
		return
	}

	db := db.OpenDB()
	if db == nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db.Collection("cookie").DeleteMany(ctx, bson.M{})

	newValue := make([]interface{}, len(cookies))
	for i := range cookies {
		newValue[i] = cookies[i].ToBson()
	}
	db.Collection("cookie").InsertMany(ctx, newValue)

	fmt.Fprintf(w, "Upload successful")
}

func UploadProxy(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	io.WriteString(w, "Upload files\n")
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err) //dont do this
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create("./uploads/proxy.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxys, err := ReadFileProxy()
	if err != nil {
		return
	}

	db := db.OpenDB()
	if db == nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db.Collection("proxy").DeleteMany(ctx, bson.M{})

	newValue := make([]interface{}, len(proxys))
	for i := range proxys {
		newValue[i] = proxys[i].ToBson()
	}
	db.Collection("proxy").InsertMany(ctx, newValue)

	fmt.Fprintf(w, "Upload successful")
}

func ReadFile() ([]entity.Cookie, error) {
	f, err := os.Open("./uploads/cookie.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var myMapSlice []entity.Cookie
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	//tach truong trong coookie
	for scanner.Scan() {
		//fmt.Printf("line: %s\n", scanner.Text())
		var cookie = scanner.Text()
		if cookie == "" {
			continue
		}
		res1 := strings.Split(cookie, ";")
		dataCookie := make(map[string]string)
		dataCookie["cookie"] = cookie
		for _, v := range res1 {
			res2 := strings.Split(v, "=")
			if len(res2) <= 1 {
				continue
			}
			var key = res2[0]
			if key != "" {
				dataCookie[key] = res2[1]
			}
		}
		fmt.Println(dataCookie)
		var ck = entity.CookieFromMap(dataCookie)
		myMapSlice = append(myMapSlice, ck)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return myMapSlice, nil
}

func ReadFileProxy() ([]entity.Proxy, error) {
	f, err := os.Open("./uploads/proxy.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var myMapSlice []entity.Proxy
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	//tach truong trong coookie
	for scanner.Scan() {
		//fmt.Printf("line: %s\n", scanner.Text())
		var proxy = scanner.Text()
		if proxy == "" {
			continue
		}
		fmt.Println(proxy)
		res1 := strings.Split(proxy, ":")
		fmt.Println(res1)
		dataCookie := make(map[string]string)
		dataCookie["host"] = res1[0] + ":" + res1[1]
		if len(res1) >= 4 {
			dataCookie["user_name"] = res1[2]
			dataCookie["pass_word"] = res1[3]
		}
		fmt.Println(dataCookie)
		var ck = entity.ProxyFromMap(dataCookie)
		myMapSlice = append(myMapSlice, ck)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return myMapSlice, nil
}
