package address_parse

import (
	"encoding/json"
	"regexp"
	"strings"
)

// 数据
type AddressList struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	ZipCode  string        `json:"zipcode"`
	Children []AddressList `json:"child"`
}

func init() {
	shipingAddress = make([]AddressList, 0)
	loadData()
}

func loadData() {
	err := json.Unmarshal([]byte(data), &shipingAddress)
	if err != nil {
		panic(err)
	}
	// 记录简写
	for _, item := range shipingAddress {
		name := item.Name
		for _, s := range provinceKey {
			name = strings.ReplaceAll(name, s, "")
		}
		provinces[item.Name] = name
		for _, subItem := range item.Children {
			name = subItem.Name
			for _, s := range cityKey {
				name = strings.ReplaceAll(name, s, "")
			}
			citries[subItem.Name] = name
		}

	}
}

type Address struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Phone  string `json:"phone"`
	//省
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Addr     string `json:"addr"`
	ZipCode  string `json:"zip_code"`
}

var (
	shipingAddress []AddressList
	search         = []string{"地址", "收货地址", "收货人", "收件人", "收货", "邮编", "电话", "：", ":", "；", ";", "，", ",", "。", " "}
	provinceKey    = []string{"特别行政区", "古自治区", "维吾尔自治区", "壮族自治区", "回族自治区", "自治区", "省省直辖", "省", "市"}
	// 存储省份简写
	provinces = make(map[string]string)
	cityKey   = []string{"布依族苗族自治州", "苗族侗族自治州", "自治州", "州", "市", "县"}
	// 存储区简写
	citries = make(map[string]string)
)

func Parse(address string) Address {
	p := Address{}
	for _, s := range search {
		address = strings.ReplaceAll(address, s, " ")
	}

	//整理电话格式
	reg := regexp.MustCompile(`(\d{3})-(\d{4})-(\d{4})`)
	address = reg.ReplaceAllString(address, `${1}${2}${3}`)
	reg = regexp.MustCompile(`(\d{3}) (\d{4}) (\d{4})`)
	address = reg.ReplaceAllString(address, `${1}${2}${3}`)
	// 获取手机号
	reg = regexp.MustCompile(`(86-[1][0-9]{10})|(86[1][0-9]{10})|([1][0-9]{10})`)
	as := reg.FindAllString(address, 1)
	if len(as) == 1 {
		p.Mobile = as[0]
		address = strings.ReplaceAll(address, as[0], " ")
	}
	// 获取电话号码
	reg = regexp.MustCompile(`(([0-9]{3,4}-)[0-9]{7,8})|([0-9]{12})|([0-9]{11})|([0-9]{10})|([0-9]{9})|([0-9]{8})|([0-9]{7})`)
	as = reg.FindAllString(address, 1)
	if len(as) == 1 {
		p.Phone = as[0]
		address = strings.ReplaceAll(address, as[0], " ")
	}
	//多个空格replace为一个
	reg = regexp.MustCompile(` {2,}`)
	address = reg.ReplaceAllString(address, " ")
	DetailParseForward(address, &p)
	return p
}

// 正向解析
func DetailParseForward(address string, p *Address) {
	AddressRead(shipingAddress, p, address, 0, make([]string, 0))
}

func AddressRead(al []AddressList, p *Address, address string, index int, repStr []string) {
	for _, item := range al {
		name := item.Name
		if index == 0 {
			name = provinces[item.Name]
		} else if index == 1 {
			name = citries[item.Name]
		} else {
			for _, k := range cityKey {
				name = strings.ReplaceAll(name, k, "")
			}
		}
		if pindex := strings.Index(address, name); pindex > -1 {
			switch index {
			// 省
			case 0:
				p.Province = item.Name
				if pindex == 0 {
					as := strings.Split(address, " ")
					p.Name = as[len(as)-1]
				} else {
					p.Name = strings.Split(address, " ")[0]
				}
				address = strings.ReplaceAll(address, p.Name, "")
				repStr = append(repStr, item.Name, name)
			//市
			case 1:
				p.City = item.Name
				for _, k := range repStr {
					address = strings.ReplaceAll(address, k, "")
				}
			//区
			case 2:
				p.County = item.Name
				p.ZipCode = item.ZipCode
				address = strings.ReplaceAll(address, item.Name, "")
				address = strings.ReplaceAll(address, name, "")
				p.Addr = address
				return
			}
			AddressRead(item.Children, p, address, index+1, repStr)
			break
		}
	}
}
