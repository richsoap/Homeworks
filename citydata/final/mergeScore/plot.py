import pandas as pd
import numpy as np
import math
import matplotlib.pyplot as plt
import seaborn as sns
sns.set_style('whitegrid')
plt.rcParams['font.sans-serif'] = ['SimHei']


def PlotVio(data, key, savePath = "", title=""):
  result = [0]*len(data)
  for index, row in data.iterrows():
    result[index] = row[key]
  fig = plt.figure(figsize=(9.0, 6.0))
  ax = sns.violinplot(result, inner='point')
  if len(title) > 0:
    ax.set_title(title)
  if len(savePath) > 0:
    plt.savefig(savePath)
  plt.show()
  return result

def PlotBox(data, key, savePath = ""):
  result = [0]*len(data)
  for index, row in data.iterrows():
    result[index] = row[key]
  fig = plt.figure(figsize=(9.0, 6.0))
  sns.boxplot(result)
  if len(savePath) > 0:
    plt.savefig(savePath)
  plt.show()
  return result

def PlotReg(data, key, savePath = "", title="", label = ''):
  result = [0]*len(data)
  for index, row in data.iterrows():
    result[index] = row[key]
  fig = plt.figure(figsize=(9.0, 6.0))
  ax = sns.distplot(result, kde=False, bins = 20, rug=True)
  ax.set_xlabel(label)
  ax.set_ylabel('计数(个)')
  if len(title) > 0:
    ax.set_title(title)
  if len(savePath) > 0:
    plt.savefig(savePath)
  plt.show()
  return result

totaldata = pd.read_csv('totalscore.csv')

attrs = ['rating','price','crime','commute','commute distance','price bonus','total','bus distance']
titles = ['评分分布图', '价格分布图','报案数量分布图','通勤得分分布图','直线距离分布图','价差奖励分布图','总分分布图']
labels = ['Airbnb总评(分)','单日住宿价格(欧元)','案件数量(件)','通勤得分(分)','直线距离(千米)','价差奖励(分)','总分(分)','最近的公交站直线距离(千米)']

for index in range(len(attrs)):
  #prices = PlotReg(totaldata, attrs[index], title=titles[index], label=labels[index], savePath='./img/'+attrs[index]+'.png')
  prices = PlotReg(totaldata, attrs[index], label=labels[index], savePath='./img/'+attrs[index]+'.png')