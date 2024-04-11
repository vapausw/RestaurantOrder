package main

import (
	"RestaurantOrder/dao"
	"RestaurantOrder/log"
	"RestaurantOrder/models"
	"RestaurantOrder/router"
	"RestaurantOrder/setting"
	"fmt"
	"github.com/golang/groupcache"
	"go.uber.org/zap"
	"os"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/config.ini")
	}
	// 加载配置文件
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE restaurant;
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 模型绑定
	dao.DB.AutoMigrate(&models.Customer{}, &models.Merchant{}, &models.Order{},
		&models.OrderItem{}, &models.Product{}, &models.PaymentInfo{})
	defer dao.MySQLClose()
	dao.InitRedis(setting.Conf.RedisConfig)
	defer dao.RdbClose()
	/*
		mysql 数据库用于存放一些用户的基础信息
		redis 数据库主要存放需要进行排序的信息以及令牌信息
	*/
	// 初始化日志记录器
	if err = log.InitLogger(setting.Conf.LogConfig); err != nil {
		zap.L().Error("init logger failed", zap.Error(err))
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	//启动 HTTP 池, 用于节点间通信
	peers := groupcache.NewHTTPPool("http://myserver:8000")
	peers.Set("http://myserver:8000", "http://otherserver:8000")
	// 注册路由
	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
