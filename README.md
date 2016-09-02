##asyoume postgresql工具包？
* asyoume是极限的传承项目，定义了一系列saas云组件
* postgresql是一个将json格式的数据库文件转换成golang形式的机构体的库
* 同时提供了一套用于操作转换后的表结构的操作接口 实现了类似orm的功能


##需要实现的功能
* sql表结构增删改查
* 实现nosql相关的操作
* 实现带缓存的查询（未实现）

##使用方法
* 获取源码  go get github.com/asyou-me/postgresql
* 安装表结构映射命令  go install github.com/asyou-me/postgresql/pgsql_map
* import "github.com/asyou-me/postgres"
