# import tabula

# df = tabula.read_pdf("D:\\gitProject\\goAlgorithm\\find\\wood\\20210827_160316_ark.pdf", encoding='gbk', pages='all')
# print(df)
# for indexs in df.index:
#     # 遍历打印企业名称
#     print(df.loc[indexs].values[1].strip())


import camelot

tables = camelot.read_pdf("D:\\gitProject\\goAlgorithm\\find\\wood\\20210827_160316_ark.pdf", flavor='stream')
print(tables[0].df)
# tables.export('foo.csv', f='csv', compress=True) 
# tables[0].to_csv('foo.csv')
# tables[0].parsing_report
for indexs in tables.index:
    print(indexs)