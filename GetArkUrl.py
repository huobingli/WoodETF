import requests
import json
import urllib3
from urllib import parse

# ark网站升级后，pdf文件需要从如下接口获取地址，包含unix时间戳和另外一个数(未知)
url = "https://ark-funds.com/wp-admin/admin-ajax.php"

def GetArkETFFile():
    headers = {
        'Host': 'ark-funds.com',
        'Connection': 'keep-alive',
        'Content-Length': '42',
        'Accept-Encoding': 'gzip, deflate, br',
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
    data = {'action':'generate_fund_holdings_pdf','pid':'1287'}
    # data = {'action':'generate_fund_holdings_csv','pid':'1287'}

    ret = requests.post(url, headers=headers, data=data)
    print(ret.status_code)
    print(ret.content)
    # print(ret.content.decode('ascii','ignore'))

    # text = '€20'.encode('utf-8')
    # print(text)
    # print(text.decode('utf-8'))


    # test = '\x83B\x00\x00dqN\xe5\xab\x81G\xeez\xeez\xcci\x7f\xce\x81TB\xab,t\xf1E\xe0\xea\xba\x05@\x91\xbd\\\x8bQ7.-\x8a\xd0`\xed\x01Q\xc2\x12V^3\xe9\xdc\xd6\x00\xad\xf75a\xcf\x83,\xf0J\x98;\x96\xbdl a\xa6$\xad\xed\xc8\xd1t\t\xbb\xdd\xe3\x98\x0c\xfb\xd2\x8c[\x0fDxJy^\xa8\xb4&2\xf0\xa5\xe6\x9eVJs\xa1\xfc\xc0\xf7\xe8\xd1t\xe8o\x00\x03'
    # test1 = test.encode("gb2312",'ignore')
    # print(test1)
    # test11 = test1.decode('utf-8','ignore').replace('\\x', '%')
    # print(test11)
    # test111 = parse.unquote(test11)
    # print(test111)
    # test2 = test.decode('utf-8','ignore')
    # print(test2)
    # # test2 = json.loads(test)
    
    # print(len(test))


if __name__ == '__main__':
    GetArkETFFile()