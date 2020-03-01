# 序言
这里将会用 [golang](https://golang.org/) 中的 [gin](https://www.yoytang.com/go-gin-doc.html) 框架 吐出漫画模块相关数据  

[点此查看接口文档](http://api_puppeteer.doc.hlzblog.top/)  

## 基于框架
该框架使用请看这里 [go-gin-example](https://email_server/blob/master/README_ZH.md)   

消息队列采用驱动模式  
目前融入 `kafka` 与 `rabbitMQ`  

###### 目前包管理已调为 `gomod`  

~~~bash
cp conf/app.ini.example conf/app.ini  
~~~

## 工具

[json转go结构体](https://www.sojson.com/json/json2go.html)  
请注意json中的数据类型  

## 常用功能

> 生成应用

~~~bash
go build -o 应用的英文名 main.go
~~~

或者  

~~~bash
make build
~~~

运行应用  

~~~bash
刚刚生成的应用地址 配置文件在系统中的绝对路径
~~~

> 格式化代码

~~~bash
make tool
~~~

##### 示例运行

~~~bash
project_path="/data/www/site/mail.ops.hlzblog.top"
app_name="mail.ops.hlzblog.top"

cd ${project_path}
rm -rf ${app_name} 
go build -o ${app_name} -v .
./${app_name} ${project_path}/conf/app.ini 
~~~

> 单元测试

~~~bash
文件必须以 ...test.go 结尾
测试函数必须以 TestX... 开头, X 可以是 _ 或者大写字母，不可以是小写字母或数字
参数：*testing.T
样本测试必须以 Example... 开头，输入使用注释的形式
TestMain 每个包只有一个，参数为 *testing.M
t.Error 为打印错误信息,并当前test case会被跳过
~~~