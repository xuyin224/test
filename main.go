package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	buf := bytes.NewBufferString("sourceURL=aa&fop=bb")
	req, err := http.NewRequest("POST", "127.0.0.1:8601/dfop", buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "19")

	fmt.Println(req)
	mac := NewMac("PjFtQJWfvKrSLYkSlV-keCKWzmXzSK1Zp3R9S5MV", "Q48lAPnTPVxq20dUmfux9HVCrtC3h-p3MCTgMyXf")
	fmt.Println(mac.SignRequest(req, incBody(req)))
	client := &http.Client{}
	fmt.Println(client.Do(req))
}
