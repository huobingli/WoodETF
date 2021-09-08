import os
import time

import camelot

from Middleware import *

from ArkETFStock import *
from Middleware import *

def parse_pdf(_pdf_name, database):
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    datetime = array[1:2][0]
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
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    date = "09/08/2021" #time.strftime("%m/%d/%Y", time.localtime(time.time()))

    for item in array[1:]:
        aa = ArkETFStock(item)
        aa.setDateTime(date)
        # print(aa.toArray())   
        insert_data(database, aa.toArray())

def AnalyseFile(array):
    for data in array:
        file_path = redis_get(data[0])
        
        print("----------------- begin analyse -------------------")
        print(data[0])
        print(file_path)

        # TODO 经常出现pdf表格处理获取失败的情况
        # if data[0] == "ARKG" or data[0] == "ARKX":
        #     parse_pdf_Q(file_path, data[0] + "_ETF")
        # else
        parse_pdf(file_path, data[0] + "_ETF")

        print("----------------- end analyse -------------------")

if __name__ == '__main__':
    file_name = "D:\\gitProject\\WoodETF\\download\\20210908\\20210908_arkg.pdf"
    parse_pdf_Q(file_name, "ARKG_ETF")