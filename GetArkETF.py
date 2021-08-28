import requests
import time

# basic_url = "https://ark-funds.com/wp-content/fundsiteliterature/holdings/ARK_INNOVATION_ETF_ARKK_HOLDINGS.pdf"

def GetArkETFFile(url, filename):
    headers = { 
        'Accept-Encoding': 'gzip, deflate, br',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
    }

    # timestamp = time.strftime("%Y%m%d", time.localtime(time.time()))
    # timestamp = timestamp + "_ark.pdf"

    ret = requests.get(url, headers=headers)

    with open(filename, 'wb') as f:
        f.write(ret.content)

# 下载文件，更新数据库
def DownFiles(filelist):

    timestamp = time.strftime("%Y%m%d", time.localtime(time.time()))
    for file in filelist:
        file
