# video
**剪辑视频项目接口**
###### 接口功能
> 剪辑视频
 
###### URL
> [http://127.0.0.1:8080/api/Index/Up](127.0.0.1:8080/api/Index/Up)
 
###### 支持格式
> FROM
 
###### HTTP请求方式
> POST
 
###### 请求参数
| >     | 参数 | 必选   | 类型         | 说明 |
| :---- | :--- | :----- | ------------ |
| url   | ture | string | 视频地址     |
| begin | true | string | 剪辑起始时间 |
| end   | true | string | 剪辑结束时间 |
 
###### 返回字段
| >    | 返回字段 | 字段类型                         | 说明 |
| :--- | :------- | :------------------------------- |
| err  | int      | 返回结果状态。0：正常；1：错误。 |
| msg  | string   | 信息                             |
 
###### 接口示例
> 地址：[http://127.0.0.1:8080/api/Index/Up](http://127.0.0.1:8080/api/Index/Up)
