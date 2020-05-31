# 不同地区的不同犯罪的占比是不是接近的
# 总体来说比例是定的，但存在个别异常值
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt
import random
crimes = ["Other theft", "Robbery", "Violence and sexual offences", "Burglary", "Public order", "Other crime", "Criminal damage and arson", "Vehicle crime", "Anti-social behaviour", "Shoplifting", "Bicycle theft", "Drugs", "Possession of weapons", "Theft from the person"]
shows = ["Other theft", "Robbery", "Violence \nand \nsexual offences", "Burglary", "Public order", "Other crime", "Criminal damage \nand \narson", "Vehicle \ncrime", "Anti-social \nbehaviour", "Shoplifting", "Bicycle theft", "Drugs", "Possession \nof \nweapons", "Theft \nfrom the person"]

def GetPercent(data):
    result = []
    for _, row in data.iterrows():
        for i in range(len(crimes)):
            crime = crimes[i]
            showcrime = shows[i]
            val = row[crime]/row['allcrime']
            if val != 0 and val != 1 and row['allcrime'] > 20:
                result.append({'lsoa': row['lsoa'], 'crime type': showcrime, 'percentage':val})
    return pd.DataFrame(result)

yeardata = pd.read_csv('lsoa_crime_count_year.csv')
percentdata = GetPercent(yeardata)
percentdata.to_csv('lsoa_crime_percent.csv')
sns.boxplot(x='crime type', y = 'percentage', data = percentdata)
#sns.violinplot(x='type', y = 'value', data = percentdata)
plt.show()

# 异常点
## E01005228 Violence and sexual offences 872/1280
## 
