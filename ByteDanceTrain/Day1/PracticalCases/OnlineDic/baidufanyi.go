package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type BaiduRequest struct {
	From  string
	To    string
	Query string
}

func (b BaiduRequest) String() string {
	return fmt.Sprintf("from=%s&to=%s&query=%s&transtype=translang&simple_means_flag=3&sign=232427.485594&token=cfd54f7d11b2f783f4441b6545ed958f&domain=common&ts=1690363122533", b.From, b.To, b.Query)
}

type BaiduResponse struct {
	TransResult struct {
		Data []struct {
			Dst string `json:"dst"`
			Src string `json:"src"`
		} `json:"data"`
		From   string `json:"from"`
		To     string `json:"to"`
		Status int    `json:"status"`
		Type   int    `json:"type"`
	} `json:"trans_result"`
	DictResult struct {
		Common struct {
			From string `json:"from"`
			Text string `json:"text"`
		} `json:"common"`
		From        string `json:"from"`
		SimpleMeans struct {
			From    string `json:"from"`
			Symbols []struct {
				Parts []struct {
					Means []struct {
						Means    []string `json:"means"`
						Part     string   `json:"part"`
						Text     string   `json:"text"`
						WordMean string   `json:"word_mean"`
					} `json:"means"`
					PartName string `json:"part_name"`
				} `json:"parts"`
				WordSymbol string `json:"word_symbol"`
			} `json:"symbols"`
			WordMeans []string `json:"word_means"`
			WordName  string   `json:"word_name"`
		} `json:"simple_means"`
		SynthesizeMeans struct {
			Symbols []struct {
				Cys []struct {
					Means []struct {
						CyID     string      `json:"cy_id"`
						MeanID   string      `json:"mean_id"`
						PartID   interface{} `json:"part_id"`
						WordMean string      `json:"word_mean"`
					} `json:"means"`
				} `json:"cys"`
				Parts      []interface{} `json:"parts"`
				SymbolID   string        `json:"symbol_id"`
				WordID     string        `json:"word_id"`
				WordSymbol string        `json:"word_symbol"`
				Xg         string        `json:"xg"`
			} `json:"symbols"`
			WordID   string `json:"word_id"`
			WordName string `json:"word_name"`
		} `json:"synthesize_means"`
		Zdict struct {
			Detail interface{} `json:"detail"`
			Simple struct {
				Chenyu interface{} `json:"chenyu"`
				Means  []struct {
					Exp []struct {
						Des []struct {
							Main string        `json:"main"`
							Sub  []interface{} `json:"sub"`
						} `json:"des"`
						Pos string `json:"pos"`
					} `json:"exp"`
					Pinyin string `json:"pinyin"`
				} `json:"means"`
			} `json:"simple"`
			Word string `json:"word"`
		} `json:"zdict"`
	} `json:"dict_result"`
	LijuResult struct {
		Double string   `json:"double"`
		Tag    []string `json:"tag"`
		Single string   `json:"single"`
	} `json:"liju_result"`
	Logid int `json:"logid"`
}

func baidufanyi() {
	client := &http.Client{}
	//var data = strings.NewReader(`from=zh&to=en&query=%E4%BD%A0%E5%A5%BD&transtype=translang&simple_means_flag=3&sign=232427.485594&token=cfd54f7d11b2f783f4441b6545ed958f&domain=common&ts=1690363122533`)
	request := BaiduRequest{
		From:  "zh",
		To:    "en",
		Query: "你好",
	}
	var data = strings.NewReader(request.String())
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/v2transapi?from=zh&to=en", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Acs-Token", "1690363099852_1690363122547_hFxh+KkAbzS/rOrPMYObehlP7jfmjdP302XEag74snyLVChztnLWuZBNl40qqpsVL0r3/fTfRsytu9/RsO/pLi5dGAp0McBZlzS75wqOJ31A5GTChFT5j3ryYY2esBXqQobJDWLMqoYdeAdAAeKjwRCpvKsvrcxUJ9Gu5WsE/rKESYESJ4WQle0t5IJj9FCrnzSJiwrnqNxeholXo5qnbPtdjMsFe7n2aceAzYoQ8+QzHxSR2E+HpqmgvMHJZqLkcsVyeVPHCz+vWy7IsxzZ8+zFusJBdozcakLF7CGBwFnez6CoRwP2rJGPYW6JRWgmtHai5wpeApMlQvzdZWWlC9VM3w06zRLhYqVGwFBtNHrra9lltiApqj2HfAbQkcoj1evfxwQE341Uu8sBkQjAKDKzevWmHyeH2a5aYneKhXPAPmXgm8IiHQzPt13+l91oLfw020JqDqMBGqONkCAeoYGPPVkrFzRlHEG8/PwkvIk=")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", `BIDUPSID=E5E58E2B391AC9F9D8AD02CB185A7E9C; PSTM=1657018093; BAIDUID=E3BD110683316BA90E4C1532029BE6CD:FG=1; BAIDUID_BFESS=E3BD110683316BA90E4C1532029BE6CD:FG=1; H_WISE_SIDS=234020_131862_114552_216853_213350_214790_110085_243881_244727_257586_257015_253022_262380_236312_261869_259305_257289_256419_264205_256222_265007_265030_265056_265138_265649_265713_265695_265615_265853_265881_265989_261036_265277_234296_234207_265785_266367_266692_266827_264768_266758_266188_266743_267067_265887_267087_257442_267096_267256_266159_265776_266566_267112_267345_267298_267065_267374_266847_267419_267452_267461_265563_267405_253528_265986_267369_267787_267712_267899_267910_267926_197096_266420_266746_266713_265367_107313_268235_268416_260335_268321_268489_266558_267031_268637_268592_268686_264513_268030_268875_268878_268848_268927_269006_269027_8000053_8000101_8000111_8000130_8000135_8000149_8000158_8000168_8000176_8000185_8000190_8000203; PSINO=7; H_WISE_SIDS_BFESS=234020_131862_114552_216853_213350_214790_110085_243881_244727_257586_257015_253022_262380_236312_261869_259305_257289_256419_264205_256222_265007_265030_265056_265138_265649_265713_265695_265615_265853_265881_265989_261036_265277_234296_234207_265785_266367_266692_266827_264768_266758_266188_266743_267067_265887_267087_257442_267096_267256_266159_265776_266566_267112_267345_267298_267065_267374_266847_267419_267452_267461_265563_267405_253528_265986_267369_267787_267712_267899_267910_267926_197096_266420_266746_266713_265367_107313_268235_268416_260335_268321_268489_266558_267031_268637_268592_268686_264513_268030_268875_268878_268848_268927_269006_269027_8000053_8000101_8000111_8000130_8000135_8000149_8000158_8000168_8000176_8000185_8000190_8000203; ZFY=xUvTUdR:ABRVv:A1ZNTrvYoRCheD7TgG9XvVp26:A5jwfA:C; RT="z=1&dm=baidu.com&si=77aa6400-0e05-4056-8daa-a08371cc76a0&ss=lkj422r6&sl=3&tt=1oj&bcn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf&ld=jiq&ul=kaj&hd=kb6"; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1690363098; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1690363098; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; ab_sr=1.0.1_YjMyNmM1NTI1OTZjMzg1NmNmNWU1ZWYwNGYyNTM2NTFiZDMyZTZkOWQ0YWEwODdjNjkwNDlhMzk3YjViNDhjNDcyYzJlNDU2OWIwMTJlZWE5MmUwNGZlZGRjM2Q2ODk0NzQxMjBkYjg0NGRmOWJiMWI2ZTRhYWM4NWJjNWIyMTRhNGVkYThmYzg1MmU3NmZjOWE1YmEwNGY4MTNlZjVkYQ==`)
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response := BaiduResponse{}
	if err = json.Unmarshal(bodyText, &response); err != nil {
		log.Fatal(err)
	}
	for _, result := range response.TransResult.Data {
		fmt.Println(result.Src, result.Dst)
	}
}
