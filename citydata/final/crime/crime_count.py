import pandas as pd
import seaborn as sns
import numpy as np
import matplotlib.pyplot as plt

sns.set_style("white")

totaldata = pd.read_csv("greater-manchester-street-merge.csv")

def CountCrime(data, offset = 0, top = 30):
    record = {}
    for _, row in data.iterrows():
        month = row['Month']
        crime = row['Crime type']
        if month not in record:
            record[month] = {}
        if crime not in record[month]:
            record[month][crime] = 0
        record[month][crime] += 1
    result = []
    print(record)
    for month in record:
        onekind = []
        for crime in record[month]:
            if record[month][crime] > offset:
                onekind.append({'Month': month, 'Crime type':crime, 'Weight':record[month][crime]})
        onekind.sort(key=lambda x:x['Weight'])
        for i in range(top):
            if i >= len(onekind):
                break
            result.append(onekind[i])
    return pd.DataFrame(result)

#CountData = CountCrime(totaldata)
#CountData.to_csv("crime_count.csv")
CountData = pd.read_csv("crime_count.csv")
#print(CountData)
sns.barplot(x = 'Month', y = 'Weight', hue='Crime type', data=CountData, dodge=True)
plt.show()
sns.barplot(x = 'Month', y = 'Weight', estimator=np.sum, data=CountData, ci=0)
plt.show()