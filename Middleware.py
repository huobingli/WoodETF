from mysql_comm.mysql_comm import *
from redis_comm.redis_comm import *

from ArkETFConf import *

def insert_datas(datas):     
    with UsingMysql(log_time=True) as um:
        for data in datas:
            sql = "insert into ARK_INNOVATION_ETF(ark_date, ark_id, ark_stock_name, ark_company, ark_shares, ark_market_value, ark_weight)  \
            values(%s, %s, %s, %s, %s, %s, %s)"
            params = ('%s' % data[6], '%d' % data[0], "%s" % data[2], "%s" % data[1], "%s" % data[3], "%s" % data[4], "%f" % data[5])
            um.cursor.execute(sql, params)

def insert_data(database, data):
    with UsingMysql(log_time=True) as um:
        sql = "insert into " + database + "(ark_date, ark_id, ark_stock_name, ark_company, ark_shares, ark_market_value, ark_weight)  \
        values(%s, %s, %s, %s, %s, %s, %s)" 
        params = ('%s' % data[6], '%d' % data[0], "%s" % data[2], "%s" % data[1], "%s" % data[3], "%s" % data[4], "%f" % data[5])
        um.cursor.execute(sql, params)

def test_insert_data():
    pass

def fecth_data(database, condition=""):
    with UsingMysql(log_time=True) as um:
        sql = 'select * from %s %s' % (database, condition)
        print(sql)
        um.cursor.execute(sql)
        data_list = um.cursor.fetchall()
        print('-- 总数: %d' % len(data_list))
        return data_list

def test_feach_data():
    database = "ARKK_ETF"
    condition = "where ark_id = 1"
    datalist = fecth_data(database, condition)
    print(datalist)

def update_data(database, setdata, condition=""):
    with UsingMysql(log_time=True) as um:
        sql = "update %s %s %s" % (database, setdata, condition)
        um.cursor.execute(sql)

def test_update_data():
    pass

def delete_data(database, condition=""):
    with UsingMysql(log_time=True) as um:
        sql = 'delete from %s %s' % (database, condition)
        um.cursor.execute(sql)

def redis_set(key, value):
    with UsingRedis(log_time=True) as ur:
        ur.set_key_value(key, value) 

def redis_get(key):
    with UsingRedis(log_time=True) as ur:
        return ur.get_key_value(key)

def update_ARK_ETF_RELA_data(tradedate):
    datalist = fecth_data("ARK_ETF_RELA")

    ark_etf_conf = []
    for data in datalist:
        filename = data["etf_download_name"]
        filename = tradedate + filename[9:]

        ak = ArkETFConf(data["etf_name"], "", filename)
        ark_etf_conf.append(ak.toArray())

    print(ark_etf_conf)
    

if __name__ == '__main__':
    pass