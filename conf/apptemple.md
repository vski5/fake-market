# 对于app.ini的备份，记得加上密码
app_name   = 你好 gin
# possible values: DEBUG, INFO, WARNING, ERROR, FATAL
log_level  = DEBUG
admin_path = /admin
excludeAuthPath="/,/welcome,/loginOut"
[app]
port = :8080
[mysql]
ip       = 127.0.0.1
port     = 3306
user     = root
password = 记得加上密码
database = storeDB

[redis]
ip   = 127.0.0.1
port = 6379
ip_Port= 127.0.0.1:6379
password = 记得加上密码

[zap]
infoDir = ./logger/test.log

[webcookie]
cookieinfoDir = .20011111.xyz