#!/usr/bin/env python
# -*- coding:utf-8 -*-

import random,time

num=1000000

# lis=[]
# for i in range(0,num):
#     lis.append(random.randint(0,num))

fo = open('data.txt','r')
str_num = fo.read()
fo.close()
str_num = str_num[0:-1]

lis = str_num.split('\n')





start=time.time()

# lis = set(lis)
# lis = list(lis)

lis=sorted(lis)

# for i in range(0,num):
# 	j=1
# 	if i==num-2:
# 		break
# 	while lis[i]==lis[i+j]:
# 		print j
# 		lis.pop(i+j)
# 		j+=1
# print lis

# n=range(0,num)
# a=[]
# for i in n:
# 	if i>0:
# 		if lis[i] == lis[i-1]:
# 			a.append(lis[i])

for i in range(num):
	a=num-(i+1)
	try:
		if lis[a]==lis[a-1]:
			lis.pop(a)
			# del(lis[a])
	except Exception, e:
		pass

# lis = set(lis)
# lis = list(lis)


end = time.time()
# print end-start
print len(lis),',',end-start

