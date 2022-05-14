# 软件介绍
daemon是一个守护进程,可使用配置文件配置想要开启的程序,并监控程序运行状态，在程序运行失败后将自动重启程序。如果有更新包，daemon将自动升级指定的程序。

# 配置信息

WaitAddress: 需要damon等待地址可用才继续运行，例如可等待数据库准备完毕

Directory: 需要监控的程序运行目录  
Exe: 需要监控程序启动文件名称  
Environment: 需要执行的环境变量  ["PYTHONPATH=/lib/python","CUDA=/lib/cuda"]  
FaildSecond: 当程序运行失败或者挂掉后重启时间间隔  
Args: 程序运行参数 ["arg1=1","arg2=2"]  
UpdateFileName: 程序更新压缩包，支持zip、tar格式，daemon检测到目录中存在此文件后，将自动解压，并重启程序  
需要运行多个监控程序，可配置多个Servers节点  


``` toml

WaitAddress = ["localhost:3306","localhost:3306"]

[[Servers]]
Directory = "/home/server"
Exe = "server"
Environment = ["PYTHONPATH=/lib/python","CUDA=/lib/cuda"]
FaildSecond = 5
Args = ["arg1=1","arg2=2"]
UpdateFileName = "update.zip"

```
