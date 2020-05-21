智能识别收货地址省市区三级

返回数据
```
resp:=Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
type ParseResp struct {
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
```
