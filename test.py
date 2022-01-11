def str_find(_name):
    return strpath.find("high")

def is_true(_bool):
    return _bool == "true"

if __name__ == '__main__':
    strpath ="small_high"
    print(strpath == "master")

    if str_find("high") != -1:
        print(str_find(strpath))
        print("no")
    else:
        print("dame")

    if is_true("true"):
        print("11")
    else:
        print("22")