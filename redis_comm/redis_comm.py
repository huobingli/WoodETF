import redis   

from setting.setting import *


def redis_conn():
    r = redis.Redis(host=host, port=redis_port, password='111111', decode_responses=True)
    print(r)
    r.set('name', 'zhangsan') 
    print(r['name'])
    print(r.get('name'))  # 取出键name对应的值
    print(type(r.get('name')))
