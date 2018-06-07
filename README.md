### Install

1.dep ensure

2.go install



### 主要思路

在对现在数据库做分表时，从一个库查询数据，并发插入另外几十个分表中，大大提高数据迁移速度，已经测试这个是性能最好的办法

目前只是一个demo，后面可以考虑做成一个通用的程序

### go-sql-driver/mysql 参数说明，具体可查阅这个库的文档

maxAllowedPacket=0 设置为0会读取当前mysql配置

interpolateParams=true 因为mysql预处理参数太长mysql会报错，取消mysql预处理，而是使用go-sql-driver/mysql对参数做安全转译