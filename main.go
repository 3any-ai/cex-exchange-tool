package main

import (
    "os"
    _"log"
    "fmt"
    "encoding/json"
    "prex/exchange"
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_"fyne.io/fyne/v2/data/binding"
)

type Config struct {
    Name     string
    ApiKey    string
    Secret   string
    PassPhrase string
    QuotCoin  string
    BaseCoin string
    OrderSpeed string
    OpenOrder  string
    BuyPrice   string
    BuyAmount  string
    SellPrice  string
    SellAmount string
    Proxy string
}


func main() {
    os.Setenv("FYNE_FONT", "./font/mplus-1c-bold.ttf")
	myApp := app.New()
	myWindow := myApp.NewWindow("exchange")

	selectExchangeText := widget.NewLabel("交易所设置")
	title := container.New(layout.NewHBoxLayout(), selectExchangeText)
	var exchangeValue = "OKEX"
	combo := widget.NewSelect([]string{"Binance", "OKEX", "Huobi"}, func(value string){
	    exchangeValue = value
	})

    apiKey := widget.NewPasswordEntry()
    secret := widget.NewPasswordEntry()
    passPhrase := widget.NewPasswordEntry()

    baseCoin := widget.NewEntry()
    baseCoin.SetPlaceHolder("USDT")
    quotCoin := widget.NewEntry()
    quotCoin.SetPlaceHolder("CORE")

    orderSpeed := widget.NewEntry()
    openOrder := widget.NewEntry()
    buyPrice := widget.NewEntry()
    buyAmount := widget.NewEntry()
    sellPrice := widget.NewEntry()
    sellAmount := widget.NewEntry()
    proxy := widget.NewEntry()

    //获取配置数据
     filePtr, err := os.Open(".key.json")
        if err != nil {
            fmt.Println("文件打开失败 [Err:%s]", err.Error())
            return
        }
        defer filePtr.Close()
        var info Config
        // 创建json解码器
        decoder := json.NewDecoder(filePtr)
        err = decoder.Decode(&info)
        if err != nil {
            fmt.Println("解码失败", err.Error())
        } else {
            fmt.Println("解码成功")
            fmt.Println(info)
        }

    //数据双向绑定
    combo.SetSelected(info.Name)
    apiKey.SetText(info.ApiKey)
    secret.SetText(info.Secret)
    passPhrase.SetText(info.PassPhrase)
    baseCoin.SetText(info.BaseCoin)
    quotCoin.SetText(info.QuotCoin)
    orderSpeed.SetText(info.OrderSpeed)
    openOrder.SetText(info.OpenOrder)
    buyPrice.SetText(info.BuyPrice)
    buyAmount.SetText(info.BuyAmount)
    sellPrice.SetText(info.SellPrice)
    sellAmount.SetText(info.SellAmount)
    proxy.SetText(info.Proxy)


    form := &widget.Form{
        Items: []*widget.FormItem{ // we can specify items in the constructor
            {
                Text: "交易所",
                Widget: combo,
            },
            {
                Text: "apiKey",
                Widget: apiKey,
            },
            {
                Text: "密钥",
                Widget: secret,
            },
            {
                Text: "passPhrase",
                Widget: passPhrase,
            },
            {
                Text: "报价币种",
                Widget: quotCoin,
            },
            {
                Text: "基础币种",
                Widget: baseCoin,
            },
            {
                Text: "下单速度(毫秒)",
                Widget: orderSpeed,
            },
            {
                Text: "挂单N毫秒内",
                Widget: openOrder,
            },
            {
                Text: "购买价格",
                Widget: buyPrice,
            },
            {
                Text: "购买数量",
                Widget: buyAmount,
            },
            {
                Text: "售出价格",
                Widget: sellPrice,
            },
            {
                Text: "售出数量",
                Widget: sellAmount,
            },
            {
                Text: "代理",
                Widget: proxy,
            },
        },
        OnSubmit: func() { //optional, handle form submission
            data := exchange.ParamsData{
            		Name:      exchangeValue,
            		ApiKey:    apiKey.Text,
            		Secret:    secret.Text,
            		PassPhrase: passPhrase.Text,
            		QuotCoin:  quotCoin.Text,
            		BaseCoin:  baseCoin.Text,
            		OrderSpeed: orderSpeed.Text,
            		OpenOrder:  openOrder.Text,
            		BuyPrice:   buyPrice.Text,
            		BuyAmount:  buyAmount.Text,
            		SellPrice:  sellPrice.Text,
            		SellAmount: sellAmount.Text,
            		Proxy: proxy.Text,
            	}

             filePtr, err := os.Create(".key.json")
             if err != nil {
                 fmt.Println("文件创建失败", err.Error())
                 return
             }
             defer filePtr.Close()
             // 创建Json编码器
             encoder := json.NewEncoder(filePtr)
             err = encoder.Encode(data)
             if err != nil {
                 fmt.Println("编码错误", err.Error())
             } else {
                 fmt.Println("编码成功")
             }

            exchange.Init(data)
        },
    }

    form.SubmitText = "启动"

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), title, form))
    myWindow.Resize(fyne.NewSize(600, 800))
	myWindow.ShowAndRun()
}
