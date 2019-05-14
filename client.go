package openjd

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func NewJdClient(AppKey string, SecretKey string) *JdClient  {
	return &JdClient{
		AppKey:AppKey,
		SecretKey:SecretKey,
		GatewayUrl:"https://router.jd.com/api",
		ApiVersion:"1.0",
		Format:"json",
		SignMethod:"md5",
		ConnectTimeout: 5,

	}
}
type JdClient struct {
	AppKey         string
	SecretKey      string
	GatewayUrl     string
	Format         string
	SignMethod     string
	ApiVersion     string
	Session     string
	ConnectTimeout int
	ReadTimeout    int
}

func (client *JdClient) SetSession(session string)  {
	client.Session = session
}
func (client *JdClient) Execute(request Request, session string) (content []byte,err error) {
	requestParams := url.Values{}
	requestParams.Set("method", request.GetApiMethodName())
	if len(session) > 0 {
		requestParams.Set("access_token", session)
	}
	apiParams := request.GetApiParas()
 	requestParams.Set("param_json",apiParams)
	params := client.createQueryParams(&requestParams)

	req, err := http.NewRequest("GET", client.GatewayUrl + "?" + params.Encode(),nil)
	if err != nil {
		return
	}
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	proxy, _ := url.Parse("http://127.0.0.1:8888")
	netTransport := &http.Transport{Proxy:http.ProxyURL(proxy)}
	//netTransport := &http.Transport{}
	httpClient := &http.Client{Transport:netTransport}
	httpClient.Timeout = time.Duration(client.ConnectTimeout)*time.Second

	response, err := httpClient.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	res, err := simplejson.NewJson(body)
	if err != nil {
		return
	}
	if responseError, ok := res.CheckGet("error_response"); ok {
		errorBytes, _ := responseError.Encode()
		err = errors.New("执行错误:" + string(errorBytes))
	}
	return body, err
}
func (client *JdClient) createQueryParams(p *url.Values) (url.Values) {
	// 公共参数
	args := url.Values{}
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	args.Add("timestamp", loc.Format("2006-01-02 15:04:05"))
	args.Add("v", client.ApiVersion)
	args.Add("format", client.Format)
	args.Add("app_key", client.AppKey)
	args.Add("sign_method", client.SignMethod)
	// 请求参数
	for key, val := range *p {
		args.Set(key, val[0])
	}
	// 设置签名
	args.Add("sign", client.generateSign(args))
	return args
}

func (client *JdClient) generateSign(args url.Values) string {
	// 获取Key
	keys := []string{}
	for k := range args {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := client.SecretKey
	for _, k := range keys {
		query += k + args.Get(k)
	}
	query += client.SecretKey
	// 使用MD5加密
	signBytes := md5.Sum([]byte(query))
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(signBytes[:]))
}
