import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

data = pd.read_csv("./data/all.csv")
print(data.head())

data['frame'] = data['frame'].astype('float')
data['errframe'] = data['errframe'].astype('float')
data['errbits'] = data['errbits'].astype('float')
data['Eb/N0'] = data['snr'].astype('float')
data['EBR'] = data['errbits'] / data['frame'] / 1008
data['EFR'] = data['errframe'] / data['frame']
sns.set_style("whitegrid")
f, axes = plt.subplots(nrows=1, ncols=2, figsize=(14, 6))
axes[0].set(yscale='log')
axes[1].set(yscale='log')
sns.lineplot(x='Eb/N0', y='EBR', hue="type", style="type", data=data, markers=True, ax=axes[0])
sns.lineplot(x='Eb/N0', y='EFR', hue="type", style="type", data=data, markers=True, ax=axes[1])
plt.show()