package sms

import (
	"net/url"
	"crypto/md5"
	"io"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"bytes"
)

type Option struct {
	SmsUser string
	MsgType string
	ApiKey  string
}

var option Option

//func Send(templateID string, phone string, vars map[string]string) error {
//
//	u, err := url.Parse("http://sendcloud.sohu.com/smsapi/send")
//	if err != nil {
//		log.Fatal(err)
//	}
//	query := u.Query()
//	query.Set("smsUser", option.SmsUser)
//	query.Set("msgType", option.MsgType)
//	query.Set("templateId", templateID)
//	query.Set("phone", phone)
//	jsonVars, err := json.Marshal(vars)
//	if err != nil {
//		return err
//	}
//	query.Set("vars", string(jsonVars))
//
//	queryString := query.Encode()
//
//	signature := md5.New()
//	io.WriteString(signature, option.ApiKey+"&"+queryString+"&"+option.ApiKey)
//	signatureResult := signature.Sum(nil)
//	query.Set("signature", hex.EncodeToString(signatureResult))
//
//	queryString = query.Encode()
//	fmt.Println("queryString", queryString)
//
//	resp, err := http.Get(u.String() + "?" + queryString)
//	if err != nil {
//		log.Println("err", err)
//		return err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("err", err)
//		return err
//	}
//	fmt.Println(string(body))
//	return nil
//}

func LoadOption(sendCloudOption Option) {
	option.ApiKey = sendCloudOption.ApiKey
	option.MsgType = sendCloudOption.MsgType
	option.SmsUser = sendCloudOption.SmsUser
}

func Send(templateID string, phone string, vars map[string]string) error {
	sms_user := option.SmsUser
	sms_key := option.ApiKey
	sms_url := "http://sendcloud.sohu.com/smsapi/send"

	jsonVars, _ := json.Marshal(&vars)
	PostParams := url.Values{
		`smsUser`:    {sms_user},
		`templateId`: {templateID},
		`msgType`:    {`0`},
		`phone`:      {phone},
		`vars`:       {string(jsonVars)},
	}
	paramsKeyS := make([]string, 0, len(PostParams))
	for k, _ := range PostParams {
		paramsKeyS = append(paramsKeyS, k)
	}
	sort.Strings(paramsKeyS)
	sb := sms_key + "&"
	for _, v := range paramsKeyS {
		sb += fmt.Sprintf("%s=%s&", v, PostParams.Get(v))
	}
	sb += sms_key
	hashMd5 := md5.New()
	io.WriteString(hashMd5, sb)
	sign := fmt.Sprintf("%x", hashMd5.Sum(nil))
	PostParams.Add("signature", sign)

	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(sms_url, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		return err
	}
	defer ResponseHandler.Body.Close()
	_, err = ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		return err
	}
	return nil
}
