import pandas as pd
import numpy as np
import math
import matplotlib.pyplot as plt
import seaborn as sns
def LoadScore(data, key):
  result = np.zeros(len(data))
  for index, row in data.iterrows():
    result[index] = row[key]
  return result

totaldata = pd.read_csv('distance.csv')
prices = LoadScore(totaldata, 'bus distance')
sns.distplot(prices, kde=True, bins=20)
plt.show()