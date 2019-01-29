#!/usr/bin/env python
# -*- coding: utf-8 -*-

import tensorflow as tf

'''
# 计算图 Computational Graph
node1 = tf.constant(3.0, dtype=tf.float32)
node2 = tf.constant(4.0)
print '\n node1, node2 = ',  (node1, node2)

session = tf.Session()
res = session.run([node1, node2])
print '\n session.run = ', res

node3 = tf.add(node1, node2)
print '\n node3 = ', node3
res = session.run(node3)
print '\n session.run = ', res
'''

'''
a = tf.placeholder(tf.float32)
b = tf.placeholder(tf.float32)
adder_node = a + b
session = tf.Session()
print '\n session run = ', session.run(adder_node, {a: 2, b: 3})
print '\n session run = ', session.run(adder_node, {a: [1, 3], b: [2, 4]})

add_and_triple = adder_node * 3
print '\n session run = ', session.run(add_and_triple, {a: [1, 3], b: [2, 4]})
'''


'''
'''
# Model parameters
W = tf.Variable([.3], dtype=tf.float32)
b = tf.Variable([-.3], dtype=tf.float32)

# Model input and output
x = tf.placeholder(tf.float32)
linear_model = W * x + b
y = tf.placeholder(tf.float32)

# Loss
squared_deltas = tf.square(linear_model - y)
loss = tf.reduce_sum(squared_deltas)

# Optimizer
optimizer = tf.train.GradientDescentOptimizer(0.01)
train = optimizer.minimize(loss)

# Training data
data_x = [1.0, 2.0, 3.0, 4.0]
data_y = [0, -1, -2, -3]

# Training loop
init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init)	# reset values to wrong
for i in range(1000):
	sess.run(train, {x: data_x, y: data_y})

print 'W = ', sess.run(W)
print 'b = ', sess.run(b)
# exit()


fixW = tf.assign(W, [-1.])
fixb = tf.assign(b, [1.])

# Evaluate training accuracy
curr_W, curr_b, curr_loss = sess.run([W, b, loss], {x: data_x, y: data_y})
print("W: %s b: %s loss: %s"%(curr_W, curr_b, curr_loss))
