# -*- coding: utf-8 -*-

import matplotlib.pyplot as plt
import numpy as np


""" ------------------ 划线：plot """
# x = np.linspace(-1, 1, 51)
# y1 = x * 2 + 1
# y2 = x ** 2
# print x
# plt.figure()
# plt.plot(x, y1)
# plt.plot(x, y2)
# plt.show()

""" ------------------ 散点图：scatter """
for i in range(1, 10):
    plt.scatter(i * 0.1, i * 0.1, s=60, c='red', alpha=.5)
plt.show()
