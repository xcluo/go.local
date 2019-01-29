# -*- coding:utf-8 -*-
import os
import numpy as np
from matplotlib import pyplot as plt
from sklearn.svm import SVC

"""
SVC, 该函数的时间复杂度超过了O（N^2）对于较大的数据集（10000样本以上）不建议使用


方法
    decision_function（X）样本X与分离超平面的距离。
    fit（X，y [，sample_weight]）根据给定的训练数据拟合SVM模型。
    get_params（[deep]）获取此估计器的参数。
    Predict（X）对X中的样本进行分类
    score（X，y [，sample_weight]）返回给定测试数据和标签的平均精度。
    set_params（\ * \ * params）设置此估计器的参数。
"""

# vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv DataSet1
# X: 训练样本；Y：样本标签；T：测试样本
X = np.array([[1,1],[1,2],[1,3],[1,4],[2,1],[2,2],[3,1],[4,1],[5,1],
              [5,2],[6,1],[6,2],[6,3],[6,4],[3,3],[3,4],[3,5],[4,3],
              [4,4],[4,5]])
Y = np.array([1]*14 + [-1]*6)
X_T=np.array([[0.5,0.5],[1.5,1.5],[3.5,3.5],[4,5.5]])
Y_T=np.array([1, 1, -1, -1])
# ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

# vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv DataSet2
filename = os.path.dirname(os.path.abspath(__file__)) + "/textset.txt"
with open(filename) as pf:
    data = pf.readlines()
dataset = [(lambda x=line.split(): ((x[0]), float(x[1]), int(x[2])))()
           for line in data if line.strip()]

X = np.array([p[:2] for p in dataset])
Y = np.array([p[-1] for p in dataset])

X, X_T = X[10:], X[:10]
Y, Y_T = Y[10:], Y[:10]
# ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

# vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv 执行模型
model = SVC(kernel='rbf', probability=True)
# model = SVC(kernel='poly', degree=2, gamma=1, coef0=0)
model.fit(X, Y)
# ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

# vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv 模型预测测试
predict = model.predict(X_T)
predict_r = np.equal(predict, Y_T)
# ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

print model.score(X_T, Y_T)
print u"\n模型：\n", model
print u"\nSVC参数：\n", model.get_params()
print u"\n测试数据：\n", X_T
print u"\n输出预测结果：\n", predict
print u"\n输出预测结果对比：", np.alltrue(predict_r)
print predict_r
print u"\n输出正类和负类支持向量总个数：\n", model.n_support_
print u"\n输出正类和负类支持向量索引：\n", model.support_
# print u"\n输出正类和负类支持向量：\n", model.support_vectors_    


plt.figure()
for i in range(len(X)):
    plt.scatter(X[i][0], X[i][1],
                c='red' if Y[i] >0 else 'blue') # 样本
for i in range(len(X_T)):
    plt.scatter(X_T[i][0], X_T[i][1], s=100, marker='s',
                c='green' if predict[i] else 'gray') # 测试结果
plt.show()



# 模型持久化
import pickle
s = pickle.dumps(model)
print type(s)     # <type 'str'> 存文件，比如 svm-sklearn.pkl
with open('svm-sklearn.pkl') as pf:         # 重新装载
    # model = pickle.load(pf)
    pass
# 或
from sklearn.externals import joblib
joblib.dump(model, 'svm-sklearn.pkl')      # 存文件中
model = joblib.load('svm-sklearn.pkl')     # 重新装载
# 生产环境使用 pmml (sklearn-pmml)
# PMML 是一种事实标准语言，用于呈现数据挖掘模型, 
# 利用XML描述和存储数据挖掘模型
