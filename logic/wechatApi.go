package logic

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func initClient() *core.Client {
	ctx := context.Background()
	mchID := "your_merchant_id"                       // 替换为你的商户号
	mchAPIv3Key := "your_mch_api_v3_key"              // 替换为你的APIv3密钥
	mchSerialNo := "your_mch_serial_number"           // 替换为你的商户证书序列号
	privateKeyPath := "/path/to/your/private/key.pem" // 替换为你的私钥文件路径

	privateKey, err := utils.LoadPrivateKeyWithPath(privateKeyPath)
	if err != nil {
		log.Fatalf("加载商户私钥失败：%v", err)
	}

	// 使用 WithWechatPayAutoAuthCipher 初始化客户端
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchSerialNo, privateKey, mchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("初始化微信支付客户端失败：%v", err)
	}

	return client
}

func createOrder(client *core.Client) {
	ctx := context.Background()
	svc := jsapi.JsapiApiService{Client: client}
	resp, _, err := svc.PrepayWithRequestPayment(ctx, jsapi.PrepayRequest{
		Appid:       core.String("your_app_id"),      // 替换为你的AppID
		Mchid:       core.String("your_merchant_id"), // 替换为你的商户号
		Description: core.String("商品描述"),
		OutTradeNo:  core.String("订单号"),
		NotifyUrl:   core.String("https://你的回调地址"),
		Amount: &jsapi.Amount{
			Total: core.Int64(100), // 订单总金额，单位为分
		},
		Payer: &jsapi.Payer{
			Openid: core.String("用户的openid"),
		},
	})

	if err != nil {
		log.Fatalf("创建订单失败: %v", err)
	} else {
		log.Printf("Prepay ID: %s", *resp.PrepayId)
	}
}

func handleNotifications(w http.ResponseWriter, req *http.Request) {
	// APIv3 密钥
	apiV3Key := "your_api_v3_key"

	// 创建一个 NotifyHandler
	handler, _ := notify.NewRSANotifyHandler(apiV3Key, nil)

	// 传递HTTP请求和你期望解析的结构体
	var transaction payments.Transaction
	_, err := handler.ParseNotifyRequest(context.Background(), req, &transaction)
	if err != nil {
		zap.L().Error("Failed to parse notification", zap.Error(err))
		http.Error(w, "failed to parse notification", http.StatusBadRequest)
		return
	}

	// 此处已通过验签，可以信任 transaction 中的数据进行业务处理
	log.Printf("Transaction ID: %s", transaction.TransactionId)
	// 正常响应微信服务器
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"code":"SUCCESS","message":"Received"}`))
}
