from mysql_comm.mysql_comm import *

from Middleware import *
from GetArkETF import *

mysql_conf={}

# 加载配置
def load_conf():
    pass

ark_etf_array={}
def feach_spider_data():
    datalist = fecth_data("ARK_ETF_RELA")

    # for data in datalist:
    #     ark_etf_array
    pass

def test_feach_spider_data():
    feach_spider_data()

def main():
    load_conf()
    feach_spider_data()



if __name__ == '__main__':
    test_feach_spider_data()
