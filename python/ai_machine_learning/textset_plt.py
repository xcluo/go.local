# -*- coding:utf-8 -*-
import os
import pprint
from matplotlib import pyplot as plt


filename = os.path.dirname(os.path.abspath(__file__)) + "/textset.txt"
with open(filename) as pf:
    data = pf.readlines()
    
dataset = [(lambda x=line.split(): ((x[0]), float(x[1]), int(x[2])))()
           for line in data if line.strip()]
# pprint.pprint(dataset)

plt.figure()
for point in dataset:
    plt.scatter(point[0], point[1], c='red' if point[2] > 0 else 'blue')
plt.show()
