# Webchat API

* 获得所有用户

	`GET	http://ip:port/user/all`

  返回
  	
  	```
  	{
    "users": [
        {
            "username": "tom",
            "isonline": "false" //是否在线
        },
        {
            "username": "tomq",
            "isonline": "true"
        }
    ]
}
  	```
  	
* 用户注册

	```
	POST	http://ip:port/user/registered
	Data:
	{
		"password":"hello",
		"username":"tomq"
	}	
	```
  	返回
  	```
  	{
    "info": "success"
	}
   ```
   
* 用户登录

	```
	POST	http://ip:port/user/login
	Data:
	{
		"password":"hello",
		"username":"tomq"
	}	
	```
  	返回
  	```
  	{
    "info": "success"
	}
   ```
   
* 用户注销

	```
	DELETE	 http://ip:port/user/logout/:username
	```
  	返回
  	```
  	{
    "info": "success"
	}
   ```
* 获得所有消息

	```
	GET	http://ip:port/message/all
	
	```

  返回	
 ```
{
    "messages": [
        {
            "body": "",
            "username": "tom",
            "time": "2017-12-20 22:16:09"
        }
    ]
}
  	```
