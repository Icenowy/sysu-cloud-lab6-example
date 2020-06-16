package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"math/rand"
	"fmt"
	"example/metrics"
)

func main(){
	http.HandleFunc("/rand", Rand)
	http.Handle("/metrics", promhttp.Handler())
	metrics.Register()
	err := http.ListenAndServe(":5565", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Rand(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()
	metrics.RequestIncrease()
	num:=os.Getenv("UPPER_LIMIT") // 在 Dockerfile 里指定
	str:=""
	name,nerr:=os.Hostname() // Hostname 看上去似乎获取到的是容器的 hostname
	lower_half:=false
	if nerr!=nil{
		name="<unknown>"
		log.Println("err:"+nerr.Error()+" Yes\n")
	}
	if num==""{
		rint:=rand.Intn(114514)
		str=fmt.Sprintf("%d/Default from %s\n", rint, name)
		if (rint < 114514/2) {
			lower_half=true
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		rint:=rand.Intn(numInt)
		str=fmt.Sprintf("%d/%d from %s\n", rint, numInt, name)
		if (rint < numInt/2) {
			lower_half=true
		}
	}
	_,err:=w.Write([]byte(str))
	if err!=nil{
		log.Println("err:"+err.Error()+" Yes\n")
	}
	if lower_half {
		metrics.RequestIncreaseLowerHalf()
	}
	timer.Observe()
}
