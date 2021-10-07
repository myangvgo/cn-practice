package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleAllRequests(w http.ResponseWriter, req *http.Request) {
	// Task 1: 接收客户端 request，并将 request 中带的 header 写入 response header
	for rKey, rVal := range req.Header {
		w.Header().Set(rKey, strings.Join(rVal, ""))
	}

	// Task 2: 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))

	// Construct response body and status code
	payload := make(map[string]string)
	payload["message"] = "hello, golang"
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	code := 200
	if strings.ToLower(req.URL.Path) == "/notfound" {
		code = 404
	}

	// Task 3: Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Println("Client IP: ", req.RemoteAddr)
	fmt.Println("Status Code: ", code)

	// Task 4: 当访问 localhost/healthz 时，应返回 200
	if strings.ToLower(req.URL.Path) == "/healthz" {
		w.WriteHeader(200)
	} else if strings.ToLower(req.URL.Path) == "/notfound" {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(200)
		w.Write(data)
		fmt.Println("Stdout: ", data)
	}
}

func main() {
	server := &http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/", handleAllRequests)
	log.Println("Server is running on http://localhost")
	log.Fatal(server.ListenAndServe(), nil)
}
