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
  delay float
    wait seconds before health turned abnormal (default 60)
  listen string
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
3. Prometheus可以通过访问/metrics接口来获取监控数据，其中键unhealth_elapsed 对应的值表示服务从不健康开始到当前时间所持续的秒数. 当服务健康时该值为0，服务不健康时该值开始累加. 

