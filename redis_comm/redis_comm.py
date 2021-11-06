import redis   
from timeit import default_timer

from setting.setting import *


def get_redis_conn():
    r = redis.Redis(host=host, port=redis_port, password='111111', decode_responses=True)
    return r
    # print(r)
    # r.set('name', 'zhangsan') 
    # print(r['name'])
    # print(r.get('name'))  # 取出键name对应的值
    # print(type(r.get('name')))

class UsingRedis(object):

    def __init__(self, commit=True, log_time=True, log_label='总用时'):
        """
        :param commit: 是否在最后提交事务(设置为False的时候方便单元测试)
        :param log_time:  是否打印程序运行总时间
        :param log_label:  自定义log的文字
        """
        self._log_time = log_time
        self._commit = commit
        self._log_label = log_label

    def get_redis_conn(self):
        try:
            self.r = redis.StrictRedis(host=host,port=redis_port,decode_responses=True)
        except Exception as e:
            print("redis连接失败,错误信息为%s" %e)
        # r = redis.Redis(host=host, port=redis_port, password='111111', decode_responses=True)
        # return r

    def __enter__(self):

        # 如果需要记录时间
        if self._log_time is True:
            self._start = default_timer()

        # 在进入的时候自动获取连接和cursor
        conn = get_redis_conn()
        try:
            self.r = redis.StrictRedis(host=host,port=port,decode_responses=True)
        except Exception as e:
            print("redis连接失败,错误信息为%s" %e)

        self._conn = conn
        return self

    def __exit__(self, *exc_info):
        self._conn.close()

        if self._log_time is True:
            diff = default_timer() - self._start
            print('-- %s: %.6f 秒' % (self._log_label, diff))

    def get_ttl(self,key):
        '''
        获取key的过期时间
        :param key:
        :return:
        '''
        return self._conn.ttl(key)

    def set_key_value(self,key,value):
        '''
        往redis中设值
        :param key:
        :param value:
        :return:
        '''
        self._conn.set(key,value)

    def get_key_value(self,key):
        '''
        往redis中获取值
        :param key:
        :return value
        '''
        return self._conn.get(key)

    def zset_set_key_value(self, set, key, value):
        mapping={key:value}
        self._conn.zadd(set, mapping)

    def zset_get_key_value(self, set, key):
        return self._conn.zscore(set, key)

    def zset_is_in_set(self, set, key):
        return self._conn.zrank(set, key) >= 0