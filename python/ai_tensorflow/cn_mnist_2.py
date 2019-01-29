#!/usr/bin/env python
# -*- coding: utf-8 -*-
import tensorflow as tf

"""
一个线性层的softmax回归模型, 
    将其扩展为一个拥有多层卷积网络的softmax回归模型
"""

# 导入数据
print u'\n----------> # 导入数据'
import input_data
mnist = input_data.read_data_sets("MNIST_data/", one_hot=True)

print u'\n----------> # 创建模型'
x = tf.placeholder("float", shape=[None, 784])
W = tf.Variable(tf.zeros([784,10]))
b = tf.Variable(tf.zeros([10]))
y_ = tf.placeholder("float", shape=[None, 10])

print u'\n----------> # 创建卷机网络'
# 卷机
def weight_variable(shape):
    initial = tf.truncated_normal(shape, stddev=0.1)
    return tf.Variable(initial)

def bias_variable(shape):
    initial = tf.constant(0.1, shape=shape)
    return tf.Variable(initial)

def conv2d(x, W):
    return tf.nn.conv2d(x, W, strides=[1, 1, 1, 1], padding='SAME')

def max_pool_2x2(x):
    return tf.nn.max_pool(x, ksize=[1, 2, 2, 1],
                          strides=[1, 2, 2, 1], padding='SAME')

print u'\n-----------> # 第一层卷积'
W_conv1 = weight_variable([5, 5, 1, 32])
b_conv1 = bias_variable([32])

x_image = tf.reshape(x, [-1,28,28,1])

h_conv1 = tf.nn.relu(conv2d(x_image, W_conv1) + b_conv1)
h_pool1 = max_pool_2x2(h_conv1)

print u'\n-----------> # 第二层卷积'
W_conv2 = weight_variable([5, 5, 32, 64])
b_conv2 = bias_variable([64])

h_conv2 = tf.nn.relu(conv2d(h_pool1, W_conv2) + b_conv2)
h_pool2 = max_pool_2x2(h_conv2)

print u'\n-----------> # 密集连接层'
W_fc1 = weight_variable([7 * 7 * 64, 1024])
b_fc1 = bias_variable([1024])

h_pool2_flat = tf.reshape(h_pool2, [-1, 7*7*64])
h_fc1 = tf.nn.relu(tf.matmul(h_pool2_flat, W_fc1) + b_fc1)

print u'\n-----------> # Dropout'
keep_prob = tf.placeholder("float")
h_fc1_drop = tf.nn.dropout(h_fc1, keep_prob)

print u'\n-----------> # 输出层'
W_fc2 = weight_variable([1024, 10])
b_fc2 = bias_variable([10])

y_conv=tf.nn.softmax(tf.matmul(h_fc1_drop, W_fc2) + b_fc2)

# with tf.name_scope('output_act'):
#     hidden = tf.nn.relu6(tf.matmul(reshape, output_weights[0]) + output_biases)
#     tf.histogram_summary('output_act', hidden)
# 

print u'\n----------> # 训练 & 评估'
# 使用之前简单的单层SoftMax神经网络模型，
# 但用更加复杂的ADAM优化器来做梯度最速下降
cross_entropy = -tf.reduce_sum(y_*tf.log(y_conv))
# 梯度下降算法
# train_step = tf.train.GradientDescentOptimizer(0.01).minimize(cross_entropy)
train_step = tf.train.AdamOptimizer(1e-4).minimize(cross_entropy)

# 评估模型
correct_prediction = tf.equal(tf.argmax(y_conv,1), tf.argmax(y_,1))
accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))

# 初始化
init = tf.global_variables_initializer()
sess = tf.InteractiveSession()
sess.run(init)

# tensorflow tensorboard
file_writer = tf.summary.FileWriter('/tmp/aaaa/', sess.graph)

# 训练
for i in range(20000):
    batch = mnist.train.next_batch(50)
    if i%100 == 0:
        train_accuracy = accuracy.eval(feed_dict={
            x:batch[0], y_: batch[1], keep_prob: 1.0})
        print "step %d, training accuracy %g"%(i, train_accuracy)
    train_step.run(feed_dict={x: batch[0], y_: batch[1], keep_prob: 0.5})

# print "test accuracy %g"%accuracy.eval(feed_dict={
#     x: mnist.test.images, y_: mnist.test.labels, keep_prob: 1.0})
tmp = accuracy.eval(feed_dict={
    x: mnist.test.images, y_: mnist.test.labels, keep_prob: 1.0})
