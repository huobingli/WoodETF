import requests
import json
import urllib3
from urllib import parse

# ark网站升级后，pdf文件需要从如下接口获取地址，包含unix时间戳和另外一个数(未知)
url = "https://ark-funds.com/wp-admin/admin-ajax.php"

def GetArkETFFileNew():
    headers = {
        'Host': 'ark-funds.com',
        'Connection': 'keep-alive',
        # 'Content-Length': '42',
        # 'Accept-Encoding': 'gzip, deflate, br',
        'Accept-Language': 'en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7',
        'X-Requested-With': 'XMLHttpRequest',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        'Origin': 'https://ark-funds.com',
        'Referer': 'https://ark-funds.com/funds/arkk/'
    }

    # timestamp = time.strftime("%Y%m%d", time.localtime(time.time()))
    # timestamp = timestamp + "_ark.pdf"
    # action=generate_fund_holdings_pdf&pid=1287
    # generate_fund_holdings_pdf&pid=1565
    a = 1287
    data = {'action':'generate_fund_holdings_pdf','pid':a}
    # data = {'action':'generate_fund_holdings_csv','pid':'1287'}

    ret = requests.post(url, headers=headers, data=data)
    print(ret.status_code)
    print(ret.content)
    
    str1 = str(ret.content, "utf-8")
    print(str1)
    geturl = json.loads(str1)
    geturl = geturl["file"]
    print(geturl)

    GetArkETF(geturl, "arkk.pdf")


def GetArkETF(url, filename):
    headers = { 
        'Accept-Encoding': 'gzip, deflate, br',
        'Cache-Control': 'no-cache',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
    }

    ret = requests.get(url, headers=headers)

    with open(filename, 'wb') as f:
        f.write(ret.content) 


if __name__ == '__main__':
    GetArkETFFile()