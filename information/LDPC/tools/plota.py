import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

alpha = pd.read_csv("./data/alpha.csv")
beta = pd.read_csv('./data/beta.csv')
print(alpha.head())
alpha['α'] = alpha['alpha']
alpha['EBR'] = alpha['ebr']
beta['β'] = beta['beta']
beta['EBR'] = beta['ebr']

sns.set_style("whitegrid")
plt.subplot(121)
plt.semilogy(alpha['α'], alpha['EBR'], 'o-')
plt.xlabel('α')
plt.ylabel('EBR')
plt.subplot(122)
plt.semilogy(beta['β'], beta['EBR'], 'o-')
plt.xlabel('β')
plt.ylabel('EBR')
plt.show()