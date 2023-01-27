# kubernetes-spring-cassandra-cluster

```bash
  helm install my-release my-repo/cassandra --set global.storageClass=standard,replicaCount=3
  
  kubectl get secret --namespace "default" my-release-cassandra -o jsonpath="{.data.cassandra-password}" | base64 -d
  
  kubectl patch svc my-release-cassandra -p '{"spec": {"type": "NodePort"}}' 
```

```bash
  kubectl exec -it my-release-cassandra-0 bash

  cqlsh -u cassandra -p DQpR9uDXE2

  CREATE KEYSPACE csdb WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};

  USE csdb;

  CREATE TABLE students ( id int PRIMARY KEY , firstname text ,lastname text, age int, lastdate timestamp);

  consistency;
```
