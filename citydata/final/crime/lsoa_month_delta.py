import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt
from scipy.stats import kstest
sns.set_style("white")
monthdata = pd.read_csv('lsoa_crime_count_withmonth.csv')
crimes = ["Other theft", "Robbery", "Violence and sexual offences", "Burglary", "Public order", "Other crime", "Criminal damage and arson", "Vehicle crime", "Anti-social behaviour", "Shoplifting", "Bicycle theft", "Drugs", "Possession of weapons", "Theft from the person", "allcrime"]
# 绘制月数据的平均值和方差，被归一化过

def Normalize(data):
    return (data - np.mean(data))/(np.max(data) - np.min(data))

def GetMonthDiff(data, crimelist):
    record = {}
    for _, row in data.iterrows():
        if row['lsoa'] not in record:
            record[row['lsoa']] = {}
            for crime in crimelist:
                record[row['lsoa']][crime] = np.zeros(12)
        for crime in crimelist:
            record[row['lsoa']][crime][row['month']] = row[crime]
    for lsoa in record:
        for crime in crimelist:
            record[lsoa][crime] = Normalize(record[lsoa][crime])
    result = []
    for lsoa in record:
        for crime in crimelist:
            for i in range(12):
                result.append({'month': i + 1, 'value': record[lsoa][crime][i], 'crime': crime})
    return pd.DataFrame(result)

allcrime = GetMonthDiff(monthdata, crimes)
#testres = kstest(allcrime['Y'], 'norm', (u, std))
#print(testres)
#sns.jointplot(x='X', y='Y', data=allcrime, kind='hex')
sns.boxplot(x='month', y='value', data=allcrime, hue = 'crime', palette='Paired')
plt.show()