import yaml
from yaml import loader

sum_file_path = "mysql_comm/setting.yaml"

if __name__ == '__main__':
    upstream = open(sum_file_path, mode='r', encoding='ANSI')
    conf = yaml.load(upstream, Loader=yaml.FullLoader)

    print(conf)