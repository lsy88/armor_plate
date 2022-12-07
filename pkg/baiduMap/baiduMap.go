package baiduMap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//获取经纬度
//百度赌徒API申请
//http://www.funboxpower.com/498.html

// 经纬度对照网址
//http://www.docin.com/p-655216087.html

// key
//pckg0S4gcS65cSZbRdlxyb4kTq3DIAsQ

// url
//http://api.map.baidu.com/geocoder?address=地址&output=输出格式类型&key=用户密钥&city=城市名

//http://api.map.baidu.com/geocoder?address=%E6%88%90%E9%83%BD&output=json&key=pckg0S4gcS65cSZbRdlxyb4kTq3DIAsQ&city=%E6%88%90%E9%83%BD

// {
//    "status":"OK",
//    "result":{
//        "location":{
//            "lng":104.047017,
//            "lat":30.645663
//        },
//        "precise":0,
//        "confidence":40,
//        "level":"\u57ce\u5e02"
//    }
//}

type JWData struct {
	Lng float64 // 经纬度
	Lat float64
}

type CtiyJWData struct {
	Location   JWData
	Precise    int
	Confidence int
	Level      string
}

type BodyData struct {
	Status string
	Result CtiyJWData
}

func Get_JWData_By_Ctiy(strCtiy string) (float64, float64) {
	resp, err := http.Get("http://api.map.baidu.com/geocoder?address=" + strCtiy + "&output=json&key=pckg0S4gcS65cSZbRdlxyb4kTq3DIAsQ&city=" + strCtiy)
	if err != nil {
		return 0.0, 0.0
	}
	body, errbody := ioutil.ReadAll(resp.Body)
	if errbody != nil {
		return 0.0, 0.0
	}
	// 解析数据
	st := &BodyData{}
	json.Unmarshal(body, &st)
	
	return st.Result.Location.Lng, st.Result.Location.Lat
}
