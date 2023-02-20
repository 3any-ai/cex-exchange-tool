package exchange

import (
	goexv2 "github.com/nntaoli-project/goex/v2"
	"github.com/nntaoli-project/goex/v2/logger"
	"github.com/nntaoli-project/goex/v2/model"
	"github.com/nntaoli-project/goex/v2/options"
	"log"
	"time"
	"strconv"
)

type ParamsData struct {
	Name string
	ApiKey string
	Secret string
	PassPhrase string
	QuotCoin string
	BaseCoin string
	OrderSpeed string
	OpenOrder string
	BuyPrice string
	BuyAmount string
	SellPrice string
    SellAmount string
    Proxy string
}

//初始化客户端和数据
func Init(p ParamsData) {
    log.Println(p)
    logger.SetLevel(logger.DEBUG)

   //socks代理
    goexv2.DefaultHttpCli.SetTimeout(5)                       // 5 second//设置日志输出级别

//     proxy,_ := strconv.ParseBool(p.Proxy)
//         if proxy  {
             goexv2.DefaultHttpCli.SetProxy(p.Proxy)
//         }
      // goexv2.DefaultHttpCli.SetProxy("socks5://127.0.0.1:7890")
//     if p.Name == "Okx" {
//
//     } else if p.Name == "Okx"  {
//
//     } else if p.Name == "Okx" {
//
//     } else {
//
//     }


    okxPrvApi := goexv2.OKx.Spot.NewPrvApi(
    		options.WithApiKey(p.ApiKey),
    		options.WithApiSecretKey(p.Secret),
    		options.WithPassphrase(p.PassPhrase))

    BuyPrice,_ := strconv.ParseFloat(p.BuyPrice, 32)
    BuyAmount,_ := strconv.ParseFloat(p.BuyAmount, 32)
    SellPrice,_ := strconv.ParseFloat(p.SellPrice, 32)
    //SellAmount,_ := strconv.ParseFloat(p.SellAmount, 32)
    OpenOrder, _ := strconv.ParseInt(p.OpenOrder, 10, 64)
    OrderSpeed, _ := strconv.ParseInt(p.OrderSpeed, 10, 64)


    for  {
        _, _, err := goexv2.OKx.Spot.GetExchangeInfo() //建议调用
        if err != nil {
            panic(err)
        }

      btcUSDTCurrencyPair,_ := goexv2.OKx.Spot.NewCurrencyPair(p.QuotCoin, p.BaseCoin)//建议这样构建CurrencyPair
      if btcUSDTCurrencyPair.Symbol != "" {
          order, _, err := okxPrvApi.CreateOrder(btcUSDTCurrencyPair, BuyAmount, BuyPrice, model.Spot_Buy, model.OrderType_Market)
          log.Println(err)
          log.Println(order)

          if err == nil {
            time.Sleep(time.Duration(OpenOrder)*time.Millisecond)
            for {
                acc, _, err := okxPrvApi.GetAccount(p.QuotCoin)
                if err == nil  {
                    order, _, err := okxPrvApi.CreateOrder(btcUSDTCurrencyPair, acc[p.QuotCoin].AvailableBalance * 0.97, SellPrice, model.Spot_Sell, model.OrderType_Market)
                    log.Println(err)
                    log.Println(order)
                   if err == nil {
                       panic(order)
                   }
                }
            }

          }
       }
       time.Sleep(time.Duration(OrderSpeed)*time.Millisecond)
    }
}