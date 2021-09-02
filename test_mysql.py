from Middleware import *

def fecth_ark_data():
    print(fecth_data("ARKK_ETF", "where ark_stock_name = 'TSLA'"))

if __name__ == '__main__':
    fecth_ark_data()