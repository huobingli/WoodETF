# 获取wood持仓数据

# 2021.10.7 
add 增加获取pdf文件post接口,post参数 "action=generate_fund_holdings_pdf&pid=1287"
https://ark-funds.com/wp-admin/admin-ajax.php

powershell 可以 go build /{目录}
cmd  不行 go build /{目录}

# 2021.11.6 
增加每次修改的新增、删除、股数变动的情况
python 存入新redis中，用zset，set为arkk ，key为stock，value为marketshares