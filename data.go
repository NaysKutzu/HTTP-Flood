package main

import(
  "time"
  "os"
  "math/rand"
  "bufio"
  "github.com/valyala/fasthttp"
  "strings"
)

type spoof struct{
  ip string
  count int
}

type header struct{
  advanced string
  a_headers []string
}

type bilder interface{
  spoofS(req *fasthttp.Request)
  headers(req *fasthttp.Request)
}

func (s spoof) spoofS(req *fasthttp.Request) {
  rangee := []string{}
  file, err := os.Open("src/rate_headers.txt")
  if err != nil{
    panic(err)
    os.Exit(1)
  }
  scan := bufio.NewScanner(file)
  for scan.Scan(){
    //fmt.Println(scan.Text() + ":" + s.ip)
    rangee = append(rangee, scan.Text())
  }
  f := rand.NewSource(time.Now().Unix())
  r := rand.New(f)
  for i:=0; i<s.count; i++{
    //fmt.Println(rangee[r.Intn(len(rangee))] + ":" + s.ip)
    req.Header.Add(string(rangee[r.Intn(len(rangee))]), string(s.ip))
  }
}

func (b header) headers(req *fasthttp.Request){
  uas := []string{}
  file, err := os.Open("src/user-agents.txt")
  if err != nil {
    panic(err)
    os.Exit(1)
  }
  scan := bufio.NewScanner(file)
  for scan.Scan(){
    uas = append(uas, scan.Text())
  }
  f := rand.NewSource(time.Now().Unix())
  r := rand.New(f)
  req.Header.Add("User-Agent", uas[r.Intn(len(uas))])
  if b.advanced == "true"{
    for i:=0; i < len(b.a_headers); i++{
      hed := strings.Split(b.a_headers[i], ":")
      req.Header.Add(hed[0], hed[1])
    }
  }
}
