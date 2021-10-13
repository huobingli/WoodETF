import requests
import time
import os
import json

import Middleware

# basic_url = "https://ark-funds.com/wp-content/fundsiteliterature/holdings/ARK_INNOVATION_ETF_ARKK_HOLDINGS.pdf"

download_dir = "/download/"

def GetArkETFFile(url, filename):
    headers = { 
        'Accept-Encoding': 'gzip, deflate, br',
        'Cache-Control': 'no-cache',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
    }

    # timestamp = time.strftime("%Y%m%d", time.localtime(time.time()))
    # timestamp = timestamp + "_ark.pdf"

    ret = requests.get(url, headers=headers)

    with open(filename, 'wb') as f:
        f.write(ret.content)

# ark网站升级后，pdf文件需要从如下接口获取地址，包含unix时间戳和另外一个数(未知)


def GetArkETFFileNew(filename, pid):
    url = "https://ark-funds.com/wp-admin/admin-ajax.php"
    headers = {
        'Host': 'ark-funds.com',
        'Connection': 'keep-alive',
        'Accept-Language': 'en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7',
        'X-Requested-With': 'XMLHttpRequest',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        'Origin': 'https://ark-funds.com',
        'Referer': 'https://ark-funds.com/funds/arkk/'
    }

    data = {'action':'generate_fund_holdings_pdf','pid':pid}

    ret = requests.post(url, headers=headers, data=data)
    print(ret.status_code)
    print(ret.content)
    
    str1 = str(ret.content, "utf-8")
    print(str1)
    geturl = json.loads(str1)
    geturl = geturl["file"]
    print(geturl)

    GetArkETF(geturl, filename)

def GetArkETF(url, filename):
    headers = { 
        'Accept-Encoding': 'gzip, deflate, br',
        'Cache-Control': 'no-cache',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
    }

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
        GetArkETFFileNew(current_dir + file[2], file[3])
        Middleware.redis_set(file[0], current_dir + file[0]+".pdf")
