import os
import time

import camelot

from Middleware import *

from ArkETFStock import *
from Middleware import *

def parse_pdf(_pdf_name, database):
    print("----------------- parse_pdf analyse -------------------")
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    datetime = array[1:2][0]
    if datetime == "":
        print(_pdf_name + "error")
    else:
        time = datetime[0][6:]

        # if time != None:
        for item in array[3:]:
            aa = ArkETFStock(item)
            aa.setDateTime(time)
            # print(aa.toArray())
            insert_data(database, aa.toArray())
    # else:
    #     print("error format")
    #     date = "09/08/2021" 
    #     for item in array[1:]:
    #         aa = ArkETFStock(item)
    #     aa.setDateTime(date) 
    #     insert_data(database, aa.toArray())


# 当输出处理失败，往往是时间处理失败
def parse_pdf_Q(_pdf_name, database):

    # print("----------------- parse_pdf_Q analyse -------------------")
    # tables = camelot.read_pdf(_pdf_name, flavor='stream')
    # array = tables[0].data
    # date = "10/04/2021" #time.strftime("%m/%d/%Y", time.localtime(time.time()))

    # for item in array[1:]:
    #     aa = ArkETFStock(item)
    #     aa.setDateTime(date)
    #     # print(aa.toArray())   
    #     insert_data(database, aa.toArray())

    print("----------------- parse_pdf_Q analyse -------------------")
    tables = camelot.read_pdf(_pdf_name, pages='all', flavor='stream')

    date = "10/08/2021" #time.strftime("%m/%d/%Y", time.localtime(time.time()))

    for table in tables:
        array = table.data
        for item in array[1:]:
            if len(item) == 7 and len(item[0]) < 3 :  # disgusting
                aa = ArkETFStock(item)
                aa.setDateTime(date)
                insert_data(database, aa.toArray())

def pre_parse_pdf(_pdf_name):
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    datetime = array[1:2][0]
    return datetime[0].find('/') == -1
    # date = "09/15/2021"

def print_parse_pdf_Q(_pdf_name):
    
    print("----------------- parse_pdf_Q analyse -------------------")
    tables = camelot.read_pdf(_pdf_name, pages='all', flavor='stream')
    print("current pdf is -------- " + _pdf_name)
    date = "10/04/2021" #time.strftime("%m/%d/%Y", time.localtime(time.time()))

    for table in tables:
        array = table.data
        for item in array[1:]:
            if len(item) == 7 and len(item[0]) < 3 :
                aa = ArkETFStock(item)
                aa.setDateTime(date)
                print(aa.toArray())

def AnalyseParseFile(array):
    # PreAnalyseFile(array)
    AnalyseFile(array)

# 判断是否可以正确处理pdf文件，需要特殊处理的表格，标志位s（special）,否则位n(normal)
def PreAnalyseFile(array):
    for data in array:
        file_path = redis_get(data[0])
        print(data[0])
        print(file_path)
        if pre_parse_pdf(file_path):
            redis_set(data[0] + "_parse", "S")
        else:
            redis_set(data[0] + "_parse", "N")

def AnalyseFile(array):
    for data in array:
        file_path = redis_get(data[0])
        
        print("----------------- begin analyse -------------------")
        print(data[0])
        print(file_path)

        # TODO 经常出现pdf表格处理获取失败的情况
        # 有资料说是    
        # if data[0] == "ARKQ" or data[0] == "ARKX":
        if redis_get(data[0] + "_parse") == 'S':
            parse_pdf_Q(file_path, data[0] + "_ETF")
            # print_parse_pdf_Q(file_path)
        else:
            print(data[0] + "_ETF")
            print("break!!!!!!!!!!")
            # parse_pdf(file_path, data[0] + "_ETF")

        print("----------------- end analyse -------------------")

if __name__ == '__main__':
    file_name = "D:\\gitProject\\WoodETF\\download\\20211007\\20211007_arkq.pdf"
    print_parse_pdf_Q(file_name)
    # tables = camelot.read_pdf(file_name)
    # camelot.plot(tables[0], kind='grid').show()