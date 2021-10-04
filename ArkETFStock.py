
class ArkETFStock():
    def __init__(self, *args):
        self.id = int(args[0][0])
        self.company = str(args[0][1])
        self.stock = str(args[0][2])
        self.shares = str(args[0][4])
        self.marketValue = str(args[0][5].replace('$', ''))
        self.weight = float(args[0][6].replace('%', ""))
        self.datetime = ""

    def toString(self):
        return self.id  + ",   " + self.company + ",   " + self.stock + ",  " + self.shares + ",  " + self.marketValue + ",  " + self.weight

    def setDateTime(self, _datetime):
        self.datetime = _datetime

    def toArray(self):
        return (self.id, self.company, self.stock, self.shares, self.marketValue, self.weight, self.datetime)