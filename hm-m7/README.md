How to run
ubuntu@master:~/chchang/homeworks/hm-m7$ sudo kubectl apply -f httpserver.yaml
deployment.apps/httpserver created
ubuntu@master:~/chchang/homeworks/hm-m7$ sudo kubectl describe pods httpserver -n default
Name:         httpserver-5f74646554-bvfcd
Namespace:    default
Priority:     0
Node:         node2/172.21.0.43
Start Time:   Mon, 21 Nov 2022 03:14:18 +0800
Labels:       app=httpserver
              pod-template-hash=5f74646554
Annotations:  cni.projectcalico.org/containerID: 68a21000601d8789da04c179445c5b1c611813b862c0efdf379143ab387fc982
              cni.projectcalico.org/podIP: 192.168.104.13/32
              cni.projectcalico.org/podIPs: 192.168.104.13/32
Status:       Running
IP:           192.168.104.13
IPs:
  IP:           192.168.104.13
Controlled By:  ReplicaSet/httpserver-5f74646554
Containers:
  httpserver:
    Container ID:   docker://f108ba240f9900f9db6ba74d1ad176f2aba2731516ea4343e43473879530a79e
    Image:          registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker:0.0.1
    Image ID:       docker-pullable://registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker@sha256:dabaf2596d6421e49f10bcbe4423d8b5e3428ca85dad89ce5cc1785fa34a76d2
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Mon, 21 Nov 2022 03:14:18 +0800
    Ready:          False
    Restart Count:  0
    Limits:
      cpu:     1
      memory:  1Gi
    Requests:
      cpu:        1
      memory:     1Gi
    Liveness:     http-get http://:8090/healthz delay=30s timeout=1s period=5s #success=1 #failure=3
    Readiness:    http-get http://:8090/healthz delay=30s timeout=1s period=5s #success=1 #failure=3
    Environment:  <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-ql7sx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  kube-api-access-ql7sx:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Guaranteed
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  12s   default-scheduler  Successfully assigned default/httpserver-5f74646554-bvfcd to node2
  Normal  Pulled     12s   kubelet            Container image "registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker:0.0.1" already present on machine
  Normal  Created    12s   kubelet            Created container httpserver
  Normal  Started    12s   kubelet            Started container httpserver

Name:         httpserver-5f74646554-r5mpl
Namespace:    default
Priority:     0
Node:         node1/172.21.0.114
Start Time:   Mon, 21 Nov 2022 03:14:18 +0800
Labels:       app=httpserver
              pod-template-hash=5f74646554
Annotations:  cni.projectcalico.org/containerID: e2f65501bd26750c70cd87fd44333d61390d6cb6f029a14dcbf36d77d2357dc7
              cni.projectcalico.org/podIP: 192.168.166.143/32
              cni.projectcalico.org/podIPs: 192.168.166.143/32
Status:       Running
IP:           192.168.166.143
IPs:
  IP:           192.168.166.143
Controlled By:  ReplicaSet/httpserver-5f74646554
Containers:
  httpserver:
    Container ID:   docker://fdcf0ae66c891088f250afca1e439d93aaca19b7897f1fcf152f9c14e344e3e0
    Image:          registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker:0.0.1
    Image ID:       docker-pullable://registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker@sha256:dabaf2596d6421e49f10bcbe4423d8b5e3428ca85dad89ce5cc1785fa34a76d2
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Mon, 21 Nov 2022 03:14:18 +0800
    Ready:          False
    Restart Count:  0
    Limits:
      cpu:     1
      memory:  1Gi
    Requests:
      cpu:        1
      memory:     1Gi
    Liveness:     http-get http://:8090/healthz delay=30s timeout=1s period=5s #success=1 #failure=3
    Readiness:    http-get http://:8090/healthz delay=30s timeout=1s period=5s #success=1 #failure=3
    Environment:  <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-htmhv (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  kube-api-access-htmhv:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Guaranteed
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  12s   default-scheduler  Successfully assigned default/httpserver-5f74646554-r5mpl to node1
  Normal  Pulled     12s   kubelet            Container image "registry.cn-beijing.aliyuncs.com/chchang_docker/ccdocker:0.0.1" already present on machine
  Normal  Created    12s   kubelet            Created container httpserver
  Normal  Started    12s   kubelet            Started container httpserver

ubuntu@master:~/chchang/homeworks/hm-m7$ sudo kubectl delete -f httpserver.yaml
deployment.apps "httpserver" deleted
ubuntu@master:~/chchang/homeworks/hm-m7$ sudo kubectl describe pods httpserver -n default
Error from server (NotFound): pods "httpserver" not found
ubuntu@master:~/chchang/homeworks/hm-m7$
