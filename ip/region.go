package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

type IpRegionResponse struct {
	Code int                  `json:"code"`
	Data IpRegionResponseData `json:"data"`
}

type IpRegionResponseData struct {
	Status      string `json:"status"`
	CountryCode string `json:"countryCode"`
}

func GetIpRegion(ip string) (string, error) {
	var url = "http://ip-api.com/json/" + ip + "?fields=status,countryCode"

	//logrus.Info("GetIpRegion, url:", url)

	status, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}

	if status != fasthttp.StatusOK {
		return "", errors.New(fmt.Sprintf("status is:%d", status))
	}

	/*
	{
	  "status": "success",
	  "countryCode": "CN"
	}
	*/

	var data = IpRegionResponseData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Status != "success" {
		return "", errors.New("requst failed")
	}

	return data.CountryCode, nil
}
