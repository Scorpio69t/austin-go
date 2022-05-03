# austin-go

#### 介绍

austin项目的golang版本


#### github地址:`https://github.com/rxrddd/austin-go`

#### 项目描述

1. 基于go-zero/grpc/ants/rabbitmq/mysql/redis 写的一个聚合消息推送平台
  
2. 你可以学习到:
  
  - go-zero的api和rpc层如何调用
  - go-zero项目如何使用gorm,以及封装gorm对应的gorm cache
  - go-zero项目中如何使用定时任务/mq消费

#### 项目启动
1. 需要搭建mysql/rabbitmq/redis服务
1. 导入数据库文件:`austin-go/sql/austin.sql`并修改`send_account`表中的`config`配置
2. 修改对应的mysql/rabbitmq/redis配置
 - rpc服务配置:`austin-go/app/austin-web/rpc/etc/austin.yaml`
 - consumer服务配置:`austin-go/app/austin-job/etc/austin-consumer.yaml`
3. 
 - 启动api服务:`austin-go/app/austin-web/api/austin.go`
 - 启动rpc服务:`austin-go/app/austin-web/rpc/austin.go`
 - 启动consumer服务:`austin-go/app/austin-job/rpc/austin-consumer.go`
4. 测试(目前只有邮箱服务)
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "test@qq.com",
        "variables": {
            "title": "测试操作",
            "content": "Hello <b>Bob</b> and <i>Cora</i>!"
        }
    },
    "messageTemplateId": 2
}'
```
5. 如需要测试去重服务则修改`message_template`表中的`deduplication_config`字段
```
{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
```


#### 目录说明

```
├── app  
│   ├── austin-admin        消息平台管理端  
│   ├── austin-common       项目公用的一些结构体,接口定义  
│   ├── austin-job          项目消费端 mq消费/定时任务  
│   ├── austin-support      项目独有的一些支持方法  
│   └── austin-web          项目对外提供的接口  
├── common                  项目使用的公共的一些方法  
├── gen.md                  生成api/rpc的脚本 参考goctl  
├── repo                    项目操作数据库逻辑  
└── sql                     项目sql文件  
```

#### 项目未完成功能

1. 对接管理平台
  
2. 实现对应的推送信息handler
  
3. 文件导入实时/定时推送
  
4. 客户端sdk



#### Thanks

go-zero : https://github.com/zeromicro/go-zero

austin : https://gitee.com/zhongfucheng/austin

gomail : https://gopkg.in/gomail.v2