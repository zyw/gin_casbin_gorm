# 1、casbin jwt测试
## 1.1、登录
请求
```shell
curl -X POST -d "username=admin&password=123456" http://localhost:8099/api/v1/login
```
响应
```json
{
    "code": 200,
    "expire": "2022-07-08T16:24:52+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTcyNjg2OTIsImlkZW50aXR5S2V5IjoiYWRtaW4iLCJsYXN0TmFtZSI6IkJvLVlpIiwib3JpZ19pYXQiOjE2NTcyNjUwOTIsInJvbGVJZEtleSI6Ild1In0.zel9VeYbrkbaWFXV2Yej-KnIIjGBTga98ncP96CQG-c"
}
```
## 1.2、接口验证登录
请求
```shell
 curl -X GET http://localhost:8099/api/v1/ping?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTcyNjg2OTIsImlkZW50aXR5S2V5IjoiYWRtaW4iLCJsYXN0Tm FtZSI6IkJvLVlpIiwib3JpZ19pYXQiOjE2NTcyNjUwOTIsInJvbGVJZEtleSI6Ild1In0.zel9VeYbrkbaWFXV2Yej-KnIIjGBTga98ncP96CQG-c
```
响应
```json
{"code":0,"message":"success","data":"pong"}
```
## 1.3、接口验证登录和访问权限
请求
```shell
curl -i -H "userName: tom" http://localhost:8099/api/v2/ping?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTcyNjg2OTIsImlkZW50aXR5S2V5IjoiYWR taW4iLCJsYXN0TmFtZSI6IkJvLVlpIiwib3JpZ19pYXQiOjE2NTcyNjUwOTIsInJvbGVJZEtleSI6Ild1In0.zel9VeYbrkbaWFXV2Yej-KnIIjGBTga98ncP96CQG-c
```
响应
```json
{"code":0,"message":"success","data":"pong"}
```
# 2、gorm简单测试
运行如下命令：
```shell
go run main.go gorm
```