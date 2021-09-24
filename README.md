# micro-mall-logistics

#### 介绍
微商城-物流系统

#### 软件架构
gRPC

#### 框架，库依赖
kelvins框架支持（gRPC，cron，queue，web支持）：https://gitee.com/kelvins-io/kelvins   
g2cache缓存库支持（两级缓存）：https://gitee.com/kelvins-io/g2cache   

#### 安装教程

1.仅构建  sh build.sh   
2 运行  sh build-run.sh   
3 停止 sh stop.sh

#### 使用说明

```toml
[kelvins-server]
Environment = "dev"

[kelvins-logger]
RootPath = "./logs"
Level = "debug"

[kelvins-auth]
Token = "c9VW6ForlmzdeDkZE2i8"
TransportSecurity = false
ExpireSecond = 100

[kelvins-mysql]
Host = "127.0.0.1:3306"
UserName = "root"
Password = "xxx"
DBName = "micro_mall_logistics"
Charset = "utf8mb4"
PoolNum =  10
MaxIdleConns = 5
ConnMaxLifeSecond = 3600
MultiStatements = true
ParseTime = true

[kelvins-redis]
Host = "127.0.0.1:6379"
Password = "xx"
DB = 6
PoolNum = 10

[kelvins-queue-amqp]
Broker = "amqp://micro-mall:szJ9aePR@localhost:5672/micro-mall"
DefaultQueue = "trade_logistics_notice"
ResultBackend = "redis://xxx@127.0.0.1:6379/6"
ResultsExpireIn = 36000
Exchange = "trade_logistics_notice"
ExchangeType = "direct"
BindingKey = "trade_logistics_notice"
PrefetchCount = 5
TaskRetryCount = 3
TaskRetryTimeout = 36000


[email-config]
Enable = false
User = "xxx@qq.com"
Password = "xxx"
Host = "smtp.qq.com"
Port = "465"

```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
