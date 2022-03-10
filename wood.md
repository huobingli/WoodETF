# 获取wood持仓数据

# 2021.10.7 
add 增加获取pdf文件post接口,post参数 "action=generate_fund_holdings_pdf&pid=1287"
https://ark-funds.com/wp-admin/admin-ajax.php

powershell 可以 go build /{目录}
cmd  不行 go build /{目录}

# 2021.11.6 
增加每次修改的新增、删除、股数变动的情况
python 存入新redis中，用zset，set为arkk ，key为stock，value为marketshares

# BACK
```
docker run -d --name docker-redis-2 -p 32336:6379 redis --requirepass "111111" --appendonly yes
docker run --name docker-mysql -p 32333:3306 -e MYSQL_ROOT_PASSWORD=111111 -d mysql
export PATH=/usr/bin:/usr/sbin:/bin:/sbin:/usr/X11R6/bin
```

## 部分数据缺少
20220309之后数据需要重新生成