package main 

import ( 
        "fmt" 
        "net/http" 
        "io/ioutil"
	"time" 
) 

func main() { 
	start := time.Now()
        res, err := http.Get("http://www.vwbeatroute.com")
	elapsed := time.Since(start).Seconds() 
	fmt.Println(elapsed)
        if err != nil { 
                fmt.Println("http.Get", err) 
                return 
        } 
        defer res.Body.Close() 
//        fmt.Println(url) 
        body, err := ioutil.ReadAll(res.Body) 
        if err != nil { 
                fmt.Println("ioutil.ReadAll", err) 
                return 
        } 
        lenp := len(body) 
        if maxp := 60; lenp > maxp { 
                lenp = maxp 
        } 
        fmt.Println(len(body), string(body[:lenp])) 
	fmt.Println(string(body))
} 
