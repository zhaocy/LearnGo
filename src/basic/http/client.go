package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request,err:=http.NewRequest(http.MethodGet,
		"http://www.imooc.com",nil)
	request.Header.Add("User-Agent","iPhone")
	//resp, err:= http.DefaultClient.Do(request)

	//resp, err:= http.Get("http://www.imooc.com")

	client:=http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err:= client.Do(request)
	if err!=nil {
		panic(err)
	}

	defer resp.Body.Close()

	s,err:= httputil.DumpResponse(resp,true)

	if err!= nil {
		panic(err)
	}

	 fmt.Printf("%s \n" ,s)
}
