import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd

filelist = [('MS', './data/MS.csv'),
            ('SP', './data/SP.csv'),
            ('NMS', './data/NMS.csv'),
            ('OMS', './data/OMS.csv')]

for filedesc in filelist:
    data = pd.read_csv(filedesc[1])
    snr = []
    efr = []
    ebr = []
    for _, row in data.iterrows():
        print(row)
        snr.append(float(row[4]))
        efr.append(float(row[1]) / float(row[0]))
        ebr.append(float(row[2]) / float(row[0]) / 1008)
    plt.subplot(121)
    plt.semilogy(snr, efr)
    plt.subplot(122)
    plt.semilogy(snr, ebr)
plt.show()