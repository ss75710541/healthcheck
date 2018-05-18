Temporary used for marathon task health check monitoring and reported the unhealth state to Prometheus.

### Build
```
make
```

### Usage 
```
docker run -p 8899:8899 -e DELAY_SECONDS=60 -e listen=0.0.0.0:8899  -d healthcheck:latest
```

```
envs:
  delay=float
      wait seconds before health turned abnormal (default 60)
  listen=addr
      listening address (default "127.0.0.1:8899")
```

### HealthCheck endpoint
```
/health 

200 OK / 500 Error
```
### Metrics endpoint
```
/metrics
```

### Prometheus labels

```
unhealth_elapsed: unhealth state elapsed time.
```

使用方法：

1. 首先下载源码，然后运行命令make来编译镜像.
2. 用编译的镜像在marathon发布应用，可以设置环境变量DELAY_SECONDS来控制应用在多少秒之后变得不健康. 默认是60秒. 可以通过环境变量LISTEN来设置服务监听的端口,默认是8899. 健康检测的接口是/health, 返回200表示成功，返回500表示失败. 

	设置 healthcheck delay 参数 30

	marathon healthcheck服务的健康检查参考值：
	
	- 健康检查间隔时间 20s
	- 健康检查失败次数 2
	- 健康检查超时10

	正常情况marathon会在healthcheck 服务启动100s（正常60s+异常40s）后Kill异常实例再启动新healthcheck实例。
	
3. Prometheus可以通过访问/metrics接口来获取监控数据，其中键unhealth_elapsed 对应的值表示服务从不健康开始到当前时间所持续的秒数. 当服务健康时该值为0，服务不健康时该值开始累加. 
5. Zabbix 收集 unhealth_elapsed数据，可以设置编写收集脚本获取异常时间

	例如：
	
	```
	curl -s 127.0.0.1:8899/metrics|grep -v '^#'|awk '/unhealth_elapsed/{print $2}'
	```
	
	zabbix 可以设置阈值，unhealth_elapsed > 60 则认为 marathon 的 healthcheck 失效，没有正常Kill 异常实时，此时触发报警 
6. 模拟故障

	- 设置 marathon 健康检查间隔时间 60s, 健康检查超时 10s , 最大连续失败次数 3, 所以marathon 允许的 故障时间为 60s*3 = 180s
	- healthcheck delay 参数设置为 30s
	- zabbix 收集故障时间 unhealth_elapsed 的间隔为 30s
	- 在 healthcheck监测服务启动 60s后 , zabbix 收集到的 unhealth_elapsed 数据会大于0，在healthcheck监测服务启动 120s-150s 后，zabbix 收集到的 unhealth_elapsed > 60 ，报警

7. 上述的healthcheck delay 参数数据 和 unhealth_elapsed报警阈值只是参考值，可以根据实际情况进行调整


