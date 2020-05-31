import os
import pandas as pd
import seaborn as sns

def GetFilelist(subname):
    result = {}
    for root, _, files in os.walk('.', topdown=False):
        for name in files:
            if name.find(subname) != -1:
                result[os.path.join(root, name)] = int(name[5:7])
    return result


def MergeCsv(filepaths):
    totalData = pd.DataFrame()
    for key in filepaths:
        csvdata = pd.read_csv(key)
        print("{} {} lines".format(key, len(csvdata)))
        datalen = len(csvdata)
        weight = [1] * datalen
        for _, row in csvdata.iterrows():
            row['Month'] = filepaths[key]
        csvdata['Weight'] = weight
        totalData = totalData.append(csvdata, ignore_index = True)
    return totalData

filelist = GetFilelist("greater-manchester-street")
total = MergeCsv(filelist)
total.to_csv("greater-manchester-street-merge.csv")
print(len(total))