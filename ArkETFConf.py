
class ArkETFConf():
    def __init__(self, *args):
        self.name = str(args[0][1])
        self.downloadUrl = str(args[0][4])
        self.downname = str(args[0][5])

    def toString(self):
        return self.name  + ",   " + self.downloadUrl + ",   " + self.downname

    def toArray(self):
        return (self.name, self.downloadUrl, self.downname)

    