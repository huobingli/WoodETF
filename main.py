from ArkETFConf import ArkETFConf

from mysql_comm.mysql_comm import *
from redis_comm.redis_comm import *

from Middleware import *
from GetArkETF import *
from ParsePdf import *


mysql_conf={}

ark_etf_url={}
ark_etf_filename={}
ark_etf_conf = []
# 获取下载链接 保存文件名
def feach_data_to_redis():
    datalist = fecth_data("ARK_ETF_RELA")

    # 下载名称和链接加入redis
    for data in datalist:
        ak = ArkETFConf(data["etf_name"], data["etf_download_url"], data["etf_download_name"])
        ark_etf_conf.append(ak.toArray())
        redis_set(data["etf_name"] + "_download_url", data["etf_download_url"])
        redis_set(data["etf_name"] + "_download_name", data["etf_download_name"])

def test_feach_data_to_redis():
    datalist = fecth_data("ARK_ETF_RELA")

    for data in datalist:
        ak = ArkETFConf(data["etf_name"], data["etf_download_url"], data["etf_download_name"])
        ark_etf_conf.append(ak.toArray())

    print(ark_etf_conf)

def test_redis():
    pass
    # redis_conn()

if __name__ == '__main__':
    # 载入配置到redis中
    # load_conf()

    # 更新redis 获取下载pdf数据
    # feach_data_to_redis()
    test_feach_data_to_redis()

    # 下载数据
    # DownFiles(ark_etf_conf)

    # 分析pdf数据
    AnalyseFile(ark_etf_conf)