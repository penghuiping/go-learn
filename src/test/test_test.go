package test

import (
	"testing"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"html"
)

func TestFunc(t *testing.T) {
	t.Fail()
}

func TestFunc1(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("数字为:", i)
	}

	m := make(map[string]interface{})

	m["age"] = 12
	m["name"] = "jack"

	fmt.Println(m)

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println("数组:", arr[i])
	}

	res := make([]string, 0, 10)
	res = append(res, "1", "2", "3", "4")
	fmt.Println(res)
	res = res[1:3]
	fmt.Println(res)

	file1, _ := os.OpenFile("/Users/penghuiping/Desktop/1.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer file1.Close()

	//file1.Write([]byte("hello world!!!"))

	file2, _ := os.OpenFile("/Users/penghuiping/Desktop/2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer file2.Close()

	io.Copy(file2, file1)

	resp, _ := http.DefaultClient.Get("http://www.baidu.com")
	bytes,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func TestFunc2(in *testing.T)  {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}
