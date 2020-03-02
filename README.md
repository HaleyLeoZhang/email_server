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
    - `title` 长度 1到


~~~bash
POST `127.0.0.1:8100/api/email/send`  

HTTP/1.1 200 OK
{
    "code": 200,
    "message": "",
    "data":
    {
        "quert_string": "token_name=End-Token&token_value=9abc3156eb404f72b8a7d9286d01307b&expire_at=1542867765" // 身份令牌信息
    }
}