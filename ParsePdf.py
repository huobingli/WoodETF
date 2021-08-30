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

    for item in array[3:]:
        aa = ArkETFStock(item)
        aa.setDateTime(time)
        # print(aa.toArray())
        insert_data(database, aa.toArray())

def parse_pdf_Q(_pdf_name, database):
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    date = "" #time.strftime("%d/%m/%Y", time.localtime(time.time()))

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

        # if data[0] == "ARKQ":
        #     parse_pdf_Q(file_path, data[0] + "_ETF")
        # else:
        parse_pdf(file_path, data[0] + "_ETF")

        print("----------------- end analyse -------------------")


if __name__ == '__main__':
    file_name = "D:\GitProject\woodETF\download/20210827_arkf.pdf"
    parse_pdf(file_name, "ARKF_ETF")