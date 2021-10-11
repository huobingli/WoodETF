import requests
import time
import os

import Middleware

# basic_url = "https://ark-funds.com/wp-content/fundsiteliterature/holdings/ARK_INNOVATION_ETF_ARKK_HOLDINGS.pdf"

download_dir = "/download/"

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
    current_dir = os.getcwd() + "/" + download_dir + "/" + timestamp + "/"
    print(current_dir)
    if not os.path.exists(current_dir):
        os.makedirs(current_dir)

    for file in filelist:
        # GetArkETFFile(file[1], current_dir + file[2])
        Middleware.redis_set(file[0], current_dir + file[0]+".pdf")
