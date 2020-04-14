import datetime
import os
import pandas as pd
import numpy as np

cityname = '重庆'
inputpath = "./data"
outputpath = './merged_daily_chongqing.csv'
sitepath = './sites.csv'
startindex= 0

df = pd.read_csv(sitepath, encoding='gbk')
df = df[df['city'] == cityname]
sitename = []
for _, row in df.iterrows():
    sitename.append(row['code'])
metname = ['date', 'hour', 'type']
addname = ['day', 'weekour']
totalcols = sitename + metname + addname

def dayToInt(day, hour):
    result = 24*day
    result += hour
    return result

def intToDay(hour):
    return int(hour / 24)

def dateToInt(date, hour = 0):
    return dayToInt(dateToDay(date), hour)

def dateToDay(inputdate):
    strDate = str(inputdate)
    dDate = datetime.datetime.strptime(strDate, "%Y%m%d")
    return int(dDate.strftime("%w"))



def newDataFrame():
    return pd.DataFrame(columns=totalcols)


def mergeCsv():
    firsttime = True
    count = 0
    errfile = []
    df = newDataFrame()
    if startindex != 0:
        firsttime = False
    elif os.path.exists(outputpath):
        os.remove(outputpath)
        firsttime = True
    for root, _, files in os.walk(inputpath):
        for filename in files:
            if not filename.endswith('csv'):
                continue
            if count < startindex:
                count += 1
                continue
            if count % 100 == 0 and count != 0:
                df.to_csv(outputpath, mode = 'a', header = firsttime, index=False)
                firsttime = False
                df = newDataFrame()
                with open('ckpt.txt', 'w') as f:
                    f.write(str(count))

            print(count, len(errfile))
            csvdata = pd.read_csv(os.path.join(root, filename))
            filtereddata = pd.DataFrame()

            siteerr = 0
            for site in sitename:
                try:
                    filtereddata[site] = csvdata[site]
                except:
                    siteerr += 1
            if siteerr == len(sitename):
                print(filename, 'skip')
                errfile.append(filename)
                continue
            else:
                print(filename, siteerr)
            for met in metname:
                filtereddata[met] = csvdata[met]

            rday = -1
            whour = -1
            for _, row in filtereddata.iterrows():
                if rday == -1:
                    rday = dateToDay(row['date'])
                if whour == -1:
                    whour = dateToInt(row['date'])
                row['day'] = rday
                row['weekhour'] = whour + int(row['hour'])
                df = df.append(row, ignore_index=True)
            count += 1
    return count, errfile

mergeCsv()