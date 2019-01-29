# -*- coding:utf-8 -*-
#!/usr/bin/python

fp = open('data.txt', 'r')
tt = fp.readlines()
fp.close()

#import pdb; pdb.set_trace()

my_list = [int(val) for val in tt if val]

my_list = sorted(my_list)

my_list = [str(val)+'\n' for val in my_list if val]

fp = open('data2.txt', 'w+')
fp.writelines(my_list)
#for val in my_list:
#    fp.write(val)
fp.close()

print 'len=', len(my_list)
