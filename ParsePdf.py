import camelot

from ArkETFStock import *
from middleware import *

def parse_pdf(_pdf_name):
    tables = camelot.read_pdf(_pdf_name, flavor='stream')
    array = tables[0].data
    datetime = array[1:2][0]
    time = datetime[0][6:]

    for item in array[3:]:
        aa = ArkETFStock(item)
        aa.setDateTime(time)
        # print(aa.toArray())
        insert_data(aa.toArray())


if __name__ == '__main__':
    file_name = "D:\\gitProject\\goAlgorithm\\find\\wood\\20210828_145754_ark.pdf"
    parse_pdf(file_name)