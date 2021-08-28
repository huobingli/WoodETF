
class ArkETFConf():
    def __init__(self, *args):
        self.name = str(args[0])
        self.downloadUrl = str(args[1])
        self.downname = str(args[2])
    
    def __init__(self, name, downloadUrl, downname):
        self.name = name
        self.downloadUrl = downloadUrl
        self.downname = downname

    def setName(self, name):
        self.name = name

    def getName(self):
        return self.name

    def setDownloadUrl(self, downloadUrl):
        self.downloadUrl = downloadUrl
    
    def getDownloadUrl(self):
        return self.downloadUrl

    def setDownname(self, downname):
        self.downname = downname

    def getDownname(self):
        return self.downname

    def toString(self):
        return self.name  + ",   " + self.downloadUrl + ",   " + self.downname

    def toArray(self):
        return (self.name, self.downloadUrl, self.downname)

