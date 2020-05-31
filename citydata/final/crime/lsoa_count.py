import pandas as pd
import seaborn as sns
import numpy as np
import matplotlib.pyplot as plt
import copy

sns.set_style("white")

totaldata = pd.read_csv("greater-manchester-street-merge.csv")
lsoadata = pd.read_csv("Lower_Layer_Super_Output_Areas_2011.csv")
posdata = pd.read_csv('lsoa_position.csv')
scores = []

def GetLSOADict(data, pos):
    result = {}
    for _, row in data.iterrows():
        lsoa = {}
        lsoaname = row['lsoa11cd']
        if lsoaname in pos:
            lsoa = pos[lsoaname]
        else:
            lsoa = {'latitude': -1, 'longitude':-1}
        lsoa['area'] = row['Shape__Area']
        result[lsoaname] = lsoa
    return result

def GetLSOAPosDict(data):
    result = {}
    for _, row in data.iterrows():
        result[row['lsoa']] = {'latitude': row['lat'], 'longitude':row['lon']}
    return result

monthlist = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
monthdict = {}
for i in range(len(monthlist)):
    monthdict[monthlist[i]] = i

def CountLSOACrimeWithMonth(data, lsoa):
    record = {}
    crimes = {}
    for _, row in data.iterrows():
        crime = row['Crime type']
        if crime not in crimes:
            crimes[crime] = 0
        lsoaid = row['LSOA code']
        if lsoaid not in record:
            record[lsoaid] = lsoa[lsoaid]
            record[lsoaid]['month'] = {}
        month = monthdict[row['Month'][:3]]
        if month not in record[lsoaid]['month']:
            record[lsoaid]['month'][month] = {}
        if crime not in record[lsoaid]['month'][month]:
            record[lsoaid]['month'][month][crime] = 0
        record[lsoaid]['month'][month][crime] += 1


    for lsoaid in record:
        yearcrime = {}
        for crime in crimes:
            yearcrime[crime] = 0
        for month in record[lsoaid]['month']:
            for crime in record[lsoaid]['month'][month]:
                yearcrime[crime] += record[lsoaid]['month'][month][crime]
        record[lsoaid]['year'] = yearcrime
    
    yearresult = []
    result = []
    for lsoaid in record:
        lsoastate = {'lsoa': lsoaid}
        if record[lsoaid]['latitude'] == -1 and record[lsoaid]['longitude'] == -1:
            continue
        for key in record[lsoaid]:
            if key != 'month' and key != 'year':
                lsoastate[key] = record[lsoaid][key]
        # month data
        for month in range(12):
            lsoastate['month'] = month
            allcrime = 0
            if month in record[lsoaid]['month']:
                for crime in crimes:
                    if crime not in record[lsoaid]['month'][month]:
                        lsoastate[crime] = 0
                    else:
                        lsoastate[crime] = record[lsoaid]['month'][month][crime]
                        allcrime += lsoastate[crime]
            else:
                for crime in crimes:
                    lsoastate[crime] = 0
            lsoastate['allcrime'] = allcrime
            result.append(copy.deepcopy(lsoastate))
        # year data
        allcrime = 0
        for crime in crimes:
            lsoastate[crime] = record[lsoaid]['year'][crime]
            allcrime += lsoastate[crime]
        lsoastate['allcrime'] = allcrime
        lsoastate['month'] = -1
        yearresult.append(copy.deepcopy(lsoastate))
    return pd.DataFrame(result), pd.DataFrame(yearresult)

pos = GetLSOAPosDict(posdata)
lsoa = GetLSOADict(lsoadata, pos)
countdata, countyear = CountLSOACrimeWithMonth(totaldata, lsoa)
countdata.to_csv("lsoa_crime_count_withmonth.csv")
countyear.to_csv("lsoa_crime_count_year.csv")
print(countdata.head())