# -*- coding:utf-8 -*-
#!/usr/bin/python

fp = open('data.txt', 'r')
tt = fp.readlines()
fp.close()

import pdb; pdb.set_trace()

my_set = set(tt)
my_list = list(my_set)

fp = open('data2.txt', 'rw')
fp.writelines(my_list)
fp.close()

print 'len=', len(my_list)
