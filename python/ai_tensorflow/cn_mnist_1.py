#!/usr/bin/env python
# -*- coding: utf-8 -*-
import tensorflow as tf

# 导入数据
print u'\n----------> # 导入数据'
import input_data
mnist = input_data.read_data_sets("MNIST_data/", one_hot=True)

print u'\n----------> # 创建模型'
x = tf.placeholder("float", [None, 784])
W = tf.Variable(tf.zeros([784,10]))
b = tf.Variable(tf.zeros([10]))
y = tf.nn.softmax(tf.matmul(x,W) + b)

print u'\n----------> # 计算模型'
y_ = tf.placeholder("float", [None,10])     # 占位符，输入正确
cross_entropy = -tf.reduce_sum(y_*tf.log(y))    # 交叉熵
# 梯度下降算法
train_step = tf.train.GradientDescentOptimizer(0.01) \
    .minimize(cross_entropy)

init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init)

# 模型循环训练1000次
# 每个步骤中，我们都会随机抓取训练数据中的100个批处理数据点
for i in range(1000):
    batch_xs, batch_ys = mnist.train.next_batch(100)
    sess.run(train_step, feed_dict={x: batch_xs, y_: batch_ys})


print u'\n----------> # 评估模型'
correct_prediction = tf.equal(tf.argmax(y,1), tf.argmax(y_,1))
# correct_prediction = tf.equal(tf.argmax(y,1), tf.argmax(y_,1))
accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))
res = sess.run(accuracy, feed_dict={x: mnist.test.images, y_: mnist.test.labels})
print 'results = ', res

