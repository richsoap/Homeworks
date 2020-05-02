import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt
import math

data = np.zeros(1)
with open('norresult.txt') as f:
    resstr = f.read()
    resstrlist = resstr.split()
    data = np.zeros(len(resstrlist))
    for i in range(len(resstrlist)):
        data[i] = float(resstrlist[i])

gauss = lambda x:1/math.sqrt(2*math.pi)*math.e**(-x*x/2)
x = np.arange(-3.5, 3.5, 0.01)
y = np.zeros(len(x))
for i in range(len(x)):
    y[i] = gauss(x[i])
sns.distplot(data)
plt.plot(x, y)
plt.show()