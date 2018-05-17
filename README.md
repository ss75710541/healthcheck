Temporary used for marathon task health check monitoring and reported the unhealth state to Prometheus.

### Build
```
make
```

### Usage 
```
docker run -d healthcheck:bc2b19b --delay=120 --listen=0.0.0.0:8899
```

```
-delay float
    wait seconds before health turned abnormal (default 60)
-listen string
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

