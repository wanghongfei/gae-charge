# GAE-Charge

扣费和报表计算服务, Go语言实现。

主要功能为: 监听kafka扣费日志流，利用Redis完成扣费，并操作MySQL将欠费广告下线; 每小时出一套分纬度报表文件。



## 编译

Go 1.9.2编译通过。

依赖：

```
go get github.com/mediocregopher/radix.v2/...
go get github.com/Shopify/sarama
go get github.com/bsm/sarama-cluster
go get github.com/go-sql-driver/mysql
```

或者使用godep工具：

```
godep restore
```



仓库根目录位于Go标准项目结构的`src/`目录下，因此将代码`clone`至`GOPATH/src`下即可，执行

```shell
go build gaecharge
```

进行编译。