from ArkETFConf import ArkETFConf
from mysql_comm.mysql_comm import *

from Middleware import *
from GetArkETF import *

mysql_conf={}

# 加载配置
def load_conf():
    pass

ark_etf_url={}
ark_etf_filename={}
ark_etf_conf = []
# 获取下载链接 保存文件名
def feach_spider_data():
    datalist = fecth_data("ARK_ETF_RELA")

    for data in datalist:
        # ark_etf_url[data["etf_name"]] = data["etf_download_url"]
        # ark_etf_filename[data["etf_name"]] = data["etf_download_name"]
        ak = ArkETFConf(data["etf_name"], data["etf_download_url"], data["etf_download_name"])
        ark_etf_conf.append(ak.toArray())
    
    # print(ark_etf_url)
    # print(ark_etf_filename)
    print(ark_etf_conf)

def test_feach_spider_data():
    feach_spider_data()

def main():
    # 载入配置到redis中
    load_conf()

    # 获取下载pdf数据
    feach_spider_data()

    
    DownFiles(ark_etf_conf, dir)

if __name__ == '__main__':
    test_feach_spider_data()
