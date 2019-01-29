#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_breadth.py


# Coding here
class Parent1(object):
    var1 = '父类1的属性'

    def get_var1(self):
        return self.var1


class Parent2(object):
    var2 = '父类2的属性'

    def get_var2(self):
        return self.var2


class Son(Parent1, Parent2):
    var = '子类的属性'


son = Son()
print son.var   # 调用实例对象的属性

print '-实例调用所有父类的属性 - '
print son.var1  # 调用父类1的属性
print son.var2  # 调用父类2的属性

print '-实例调用父类方法 - '
print son.get_var1()
print son.get_var2()
