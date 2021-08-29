import yaml

from redis_comm.redis_comm import *

sum_file_path = "setting/setting.yaml"

# 加载配置
def load_conf():
    upstream = open(sum_file_path, mode='r', encoding='ANSI')
    conf = yaml.load(upstream, Loader=yaml.FullLoader)
    print(conf)

    # redis_set() 

def load_mysql_db():
    pass


if __name__ == '__main__':
    load_conf()