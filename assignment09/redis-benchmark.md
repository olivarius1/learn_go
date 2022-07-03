# 第九周作业
-  使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

```
Redis server v=6.2.6
Hardware Overview:
     Model Name: MacBook Air
     Chip: Apple M1
     Total Number of Cores: 8 (4 performance and 4 efficiency)
     Memory: 16 GB
```

```
datasize 分别设置10 20 50 100 200 1000 5000
redis-benchmark -q -t set,get -r 20000 -n 100000 -d {datasize}
```
测试结果

| datasize(bytes) | qps(get) |qps(set) |
|--|--|--|
| 10 | 215517.25 | 206185.56 |
| 20 | 214592.28 | 202839.75 |
| 50 | 212765.95 | 200400.80 |
| 100 |212765.95  | 200400.80 |
| 200 |211416.50|198807.16|
| 1k |207468.88| 199203.20 |
| 5k | 196850.39 |186567.16  |


随着value增大，get/set吞吐降低，但仍有较好的性能.

- 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

测试之前`used_memory_dataset` ≈71000


|valuesize(bytes)    |used_memory_dataset|key number | bytes/key|
|--|--|--|--|
| 10  |4917688  | 86542  | 56|
|  20 |6294560  | 86496 |  71|
|  50  |9063768  | 86507 |  103|  
|  100  | 13236664 | 86641 |  151| 
|  200  | 21487928 | 86373 |  247| 
|  1000  | 90675600 | 86458 | 1047 |
|  5000  | 446232368 | 86466 |  5159|


见图表

随着 value 逐渐增大, 平均每隔 key 占用内存向 value 的大小收敛





## 图表:
【金山文档】 工作簿1
https://kdocs.cn/l/cfl9twfyL4wD



