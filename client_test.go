package openjd_test

import (
	"fmt"
	"github.com/go-pack/openjd"
	"github.com/go-pack/openjd/requests"
	"github.com/go-pack/openjd/responses"
	"github.com/spf13/viper"
	"testing"
)

func init()  {
	viper.SetConfigName(".config") //  设置配置文件名 (不带后缀)
	viper.AddConfigPath("/etc/test/")   // 第一个搜索路径
	viper.AddConfigPath("$HOME/.test")  // 可以多次调用添加路径
	viper.AddConfigPath(".")               // 比如添加当前目录
	viper.SetConfigType("yml")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
func TestClient(t *testing.T) {
	session := ""
	topClient := openjd.NewJdClient(viper.GetString("APP_JD_KEY") , viper.GetString("APP_JD_SECRET") )
	request := requests.NewJdUnionOpenOrderRequest()
	request.SetTime("201905141335")
	request.SetPageNo(1);
	request.SetPageSize(30);
	request.SetType(1);
	request.SetKey("");

	content, err := topClient.Execute(request, session)
	if err != nil {
		println(err.Error())
		return
	}
	result := responses.NewUnionOpenOrderQueryResponse(content)
	for _, value := range result.Resp[0].SkuList {
		println(value.SkuName)
	}
}
