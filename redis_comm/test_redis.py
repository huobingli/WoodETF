import redis


def test_redis():
    r = redis.Redis(host="47.114.171.118", port=32336, password='111111', decode_responses=True)   # host是redis主机，需要redis服务端和客户端都启动 redis默认端口是6379
    # print(r)
    # r.set('name', 'zhangsan') 
    # print(r['name'])
    # print(r.get('name'))  # 取出键name对应的值
    # print(type(r.get('name')))
    print(r.zrank("ARKW_ETF", "SNOW"))

if __name__ == '__main__':
    test_redis()
