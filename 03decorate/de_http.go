package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
)

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		next.ServeHTTP(w,r)
		log.Printf("时间:%s",time.Now())
	})
}
func adderRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		next.ServeHTTP(w,r)
		log.Printf("地址:%s",r.RemoteAddr)
	})
}
func urlRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		next.ServeHTTP(w,r)
		log.Printf("方法和url:%s,%s",r.Method,r.URL)
	})
}

func hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello")
}

func main()  {
	http.Handle("/",timeRecording(adderRecording(urlRecording(http.HandlerFunc(hello)))))
	http.ListenAndServe(":8080",nil)
}