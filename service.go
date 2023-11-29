package main

import (
	"fmt"
	L "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
)

const NumberOfValuesInResult = 5

type HttpService struct {
	Values []prime_value_type
	Port   uint16
	Id     string
}

type Response struct {
	Id     string
	Values [NumberOfValuesInResult]prime_value_type
}

func (service *HttpService) MakeResponse() Response {
	r := Response{
		Id: service.Id,
	}
	for i := range r.Values {
		r.Values[i] = service.Values[rand.Intn(len(service.Values))]
	}
	return r
}

func (response *Response) String() string {
	r := response.Id + " "
	for i, v := range response.Values {
		if i > 0 {
			r += " "
		}
		r += fmt.Sprint(v)
	}
	return r + "\n"
}

func (service *HttpService) Run() error {
	hPlainText := func(writer http.ResponseWriter, request *http.Request) {
		L.Tracef("%s %s %s", request.RemoteAddr, request.Method, request.URL.String())
		//writer.Header().Set("Content-Type", "text/plain")
		response := service.MakeResponse()
		_, err := fmt.Fprint(writer, response.String())
		if err != nil {
			L.Error(err)
		}
	}
	http.HandleFunc("/text", hPlainText)
	address := fmt.Sprintf(":%d", service.Port)
	L.Infof("start service over %s", address)
	return http.ListenAndServe(address, nil)
}
