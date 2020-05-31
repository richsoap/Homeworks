# 不同地区的不同犯罪的数量是不同的
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
            result.append({'lsoa': row['lsoa'], 'Crime type': showcrime, 'Quantity':row[crime]})
    return pd.DataFrame(result)

yeardata = pd.read_csv('lsoa_crime_count_year.csv')
percentdata = GetPercent(yeardata)
percentdata.to_csv('lsoa_crime_box.csv')
sns.boxplot(x='Crime type', y = 'Quantity', data = percentdata)
#sns.violinplot(x='type', y = 'value', data = percentdata)
plt.show()

# 异常点
## E01005228 Violence and sexual offences 872/1280
## 
