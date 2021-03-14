package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type IpInfo struct {
	Ip       string `json:"ip"`       // ip
	Pro      string `json:"pro"`      // 省会
	ProCode  string `json:"proCode"`  // 邮政编码
	City     string `json:"city"`     // 城市
	CityCode string `json:"cityCode"` // 城市邮政编码
	Addr     string `json:"addr"`     // 运营商
}

// 返回：IP地址的信息（格式：字符串的json）
func GetCityByIp(ipAddr string) string {
	if ipAddr == "[::1]" || ipAddr == "127.0.0.1" {
		return "内网IP"
	}
	ip := IpInfo{}
	client := http.Client{Timeout: 5 * time.Second}
	r, _ := client.Get(`http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=` + ipAddr)

	body, _ := io.ReadAll(r.Body)
	res, _ := GbkToUtf8(body) // 转换编码格式 (return:[]bytes,error)
	_ = json.Unmarshal(res, &ip)
	if ip.City == "" {
		ip.City = ip.Addr
		return ip.City
	}
	return ip.City
}

// 返回：IP地址的信息(结构体)
func IPInfo(ipAddr string) (ip IpInfo) {
	if ipAddr == "[::1]" || ipAddr == "127.0.0.1" {
		ip.City = "内网IP"
		return ip
	}
	client := http.Client{Timeout: 5 * time.Second}
	r, _ := client.Get(`http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=` + ipAddr)

	body, _ := io.ReadAll(r.Body)
	res, _ := GbkToUtf8(body) // 转换编码格式 (return:[]bytes,error)
	if err := json.Unmarshal(res, &ip); err != nil {
		return ip
	}
	if ip.City == "" {
		ip.City = ip.Addr
		return ip
	}
	return ip
}
