package main

import (
    "fmt"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
)


type Result struct {
    Timestamp  int64 `json:"timestamp"`
    Buy  string `json:"buy"`
}

func perror(err error) {
    if err != nil {
        panic(err)
    }
}


func main() {
    url := "http://api.piyasa.com/json/?kaynak=metal_arsiv_ay_alti_GRM"

    res, err := http.Get(url)
    perror(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    perror(err)

    var jsonBlob = []byte(body)
    var results []Result

    json.Unmarshal(jsonBlob, &results)
    
    if err != nil {
        fmt.Printf("%T\n%s\n%#v\n",err, err, err)
        switch v := err.(type){
            case *json.SyntaxError:
                fmt.Println(string(body[v.Offset-40:v.Offset]))
        }
    }

    for i := range results {

        t := &results[i]
        t1 := time.Unix(t.Timestamp, 0).Format(time.RFC822)
            
        fmt.Printf("%s : %s\n",t1, t.Buy)

    }
}