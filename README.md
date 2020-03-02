# 序言
这里将会用 [golang](https://golang.org/) 中的 [gin](https://www.yoytang.com/go-gin-doc.html) 框架 吐出漫画模块相关数据  

[点此查看接口文档](http://api_puppeteer.doc.hlzblog.top/)  

## 基于框架
该框架使用请看这里 [go-gin-example](https://email_server/blob/master/README_ZH.md)   

消息队列采用驱动模式  
已完成驱动 `rabbitMQ`  

###### 目前包管理已调为 `gomod`  

~~~bash
cp conf/app.ini.example conf/app.ini  
~~~

## 工具

[json转go结构体](https://www.sojson.com/json/json2go.html)  
请注意json中的数据类型  

## 常用功能
集成命令都已集成到 `Makefile`

> 生成应用

~~~bash
make build
~~~

> 生成配置文件

~~~bash
cp conf/app.ini.example conf/app.ini  
~~~

设置好配置文件后,生成配置文件到默认目录下

~~~bash
make ini
~~~

rabbitMQ 请配置 `exchange 名为 `email_sender` 
rabbitMQ 请配置 `queue 名为 `email_sender` 

数据库、表设置  
~~~sql
CREATE DATABASE common_service charset = utf8mb4;

DROP TABLE IF EXISTS `email`;
CREATE TABLE `email`  (
  `id` int(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮件标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件内容',
  `sender_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发件者姓名.发起方自定义',
  `receiver` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '接收者邮箱.多个以逗号隔开',
  `receiver_name` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '接收者姓名.多个以逗号隔开',
  `attachment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件信息',
  `is_ok` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '枚举值 0:发送成功,1:发送失败',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '枚举值 0:正常,1:删除',
  `updated_at` datetime(0) NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '更新时间',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx-created_at`(`created_at`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '邮件服务' ROW_FORMAT = Compact;
~~~

> 更多

###### 格式化代码

~~~bash
make tool
~~~

###### 单元测试

~~~bash
make test
~~~

关于单元测试的书写  

~~~bash
文件必须以 ...test.go 结尾
测试函数必须以 TestX... 开头, X 可以是 _ 或者大写字母，不可以是小写字母或数字
参数：*testing.T
样本测试必须以 Example... 开头，输入使用注释的形式
TestMain 每个包只有一个，参数为 *testing.M
t.Error 为打印错误信息,并当前test case会被跳过
~~~

##### 示例运行

~~~bash
make build
./email_server
~~~

## 使用

#### 发送邮件

- API `127.0.0.1:8100/api/email/send`  
- 方式 `POST`
- 入参
    - `title` 邮件名
    - `content` 待发送正文,支持html
    - `sender_name` string 发件人昵称
    - `receiver` string 接收者邮箱.多个以逗号隔开
    - `receiver_name` string 接收者邮箱昵称,可以不填,多个以逗号隔开

~~~bash
HTTP/1.1 POST `127.0.0.1:8100/api/email/send`  
{
    "code": 200,
    "message": "success",
    "data": null
}
~~~

> TODO

支持多附件上传  