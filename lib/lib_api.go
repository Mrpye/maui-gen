package lib

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//Headers struct for the CallApi function
type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//General function for making api requests
func CallApi(url string, method string, headers []Header, payload io.Reader, ignore_ssl bool) ([]byte, bool, error) {
	//****************************
	//Create the connection string
	//****************************
	req, err := http.NewRequest(method, url, payload) //New Request
	if err != nil {
		return nil, false, err
	}
	//***************
	//Set the headers
	//***************
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: ignore_ssl}
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	//*************
	//Read the data
	//*************
	responseData, err := ioutil.ReadAll(resp.Body)
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if err != nil {
		return responseData, false, err
	} else {
		if !statusOK {
			fmt.Println("Non-OK HTTP status:", resp.StatusCode)
			return responseData, false, err
		}
	}
	return responseData, true, nil
}
