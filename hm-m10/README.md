
1. 为 HTTPServer 添加 0-2 秒的随机延时；
2. 为 HTTPServer 项目添加延时 Metric；

http://127.0.0.1:8080/metrics

# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
...
# HELP http_server_func_latency_seconds function latency analysis for seconds
# TYPE http_server_func_latency_seconds histogram
http_server_func_latency_seconds_bucket{step="total",le="0.001"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.002"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.004"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.008"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.016"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.032"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.064"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.128"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.256"} 0
http_server_func_latency_seconds_bucket{step="total",le="0.512"} 0
http_server_func_latency_seconds_bucket{step="total",le="1.024"} 2
http_server_func_latency_seconds_bucket{step="total",le="2.048"} 6
http_server_func_latency_seconds_bucket{step="total",le="4.096"} 6
http_server_func_latency_seconds_bucket{step="total",le="8.192"} 6
http_server_func_latency_seconds_bucket{step="total",le="16.384"} 6
http_server_func_latency_seconds_bucket{step="total",le="+Inf"} 6
http_server_func_latency_seconds_sum{step="total"} 7.5345718779999995
http_server_func_latency_seconds_count{step="total"} 6
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0


3. 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
httpserver
chchang@master-1:~/cnhomeworks/hm-m10$ kubectl get svc
NAME                            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
httpserver                      ClusterIP   10.99.94.164     <none>        80/TCP         6s
chchang@master-1:~/cnhomeworks/hm-m10$ kubectl get ep
NAME                            ENDPOINTS                               AGE
httpserver                      10.244.1.24:8080                        40s
chchang@master-1:~/cnhomeworks/hm-m10$ kubectl get pods
NAME                                           READY   STATUS             RESTARTS         AGE
httpserver-675dd98448-kjgmb                    0/1     Running            0                27s
httpserver-675dd98448-s45fb                    0/1     Running            0                27s

chchang@node2:~$ curl -k -v http://10.244.1.24:8080
*   Trying 10.244.1.24:8080...
* Connected to 10.244.1.24 (10.244.1.24) port 8080 (#0)
> GET / HTTP/1.1
> Host: 10.244.1.24:8080
> User-Agent: curl/7.81.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 05 Jan 2023 22:02:29 GMT
< Content-Length: 181
< Content-Type: text/plain; charset=utf-8
<
hello [stranger]
===================Details of the http request header:============
User-Agent=[curl/7.81.0]
Accept=[*/*]
===================Server go version:============
Version=
* Connection #0 to host 10.244.1.24 left intact
chchang@node2:~$


Prometheus 
chchang@master-1:~/cnhomeworks/hm-m10$ sudo snap install helm --classic
helm 3.7.0 from Snapcrafters installed

chchang@master-1:~/cnhomeworks/hm-m10$ helm repo add grafana https://grafana.github.io/helm-charts

chchang@master-1:~/cnhomeworks/hm-m10$ helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false

chchang@master-1:~/cnhomeworks/hm-m10$ kubectl get pods
NAME                                           READY   STATUS    RESTARTS       AGE
loki-0                                         0/1     Running   1 (15s ago)    2m17s
loki-grafana-7cc8d5fbdc-48jbm                  2/2     Running   2 (45s ago)    2m17s
loki-kube-state-metrics-769569844d-slhlk       1/1     Running   2 (110s ago)   2m17s
loki-prometheus-alertmanager-d8d759b88-7hrbv   2/2     Running   2 (74s ago)    2m17s
loki-prometheus-node-exporter-hmkdn            1/1     Running   1 (2m5s ago)   2m17s
loki-prometheus-node-exporter-mlv9k            1/1     Running   0              2m17s
loki-prometheus-pushgateway-6848bb6f5c-wrrqk   1/1     Running   2 (65s ago)    2m17s
loki-prometheus-server-65c886cc5c-tbzms        2/2     Running   2 (44s ago)    2m17s
loki-promtail-4psd2                            1/1     Running   0              2m17s
loki-promtail-h4lc2                            0/1     Running   2 (35s ago)    2m17s
loki-promtail-k7txq                            1/1     Running   0              2m17s
chchang@master-1:~/cnhomeworks/hm-m10$



