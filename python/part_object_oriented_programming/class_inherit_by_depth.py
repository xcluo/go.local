#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_depth.py


# Coding here
class Grandpa(object):
    var2 = '爷爷类的属性'

    def get_var2(self):
        return self.var2


class Parent(Grandpa):
    var1 = '父类的属性'

    def get_var1(self):
        return self.var1


class Son(Parent):
    var = '子类的属性'


son = Son()
print son.var   # 调用实例对象的属性

print '- 实例调用父类+爷爷类的属性 - '
print son.var1  # 调用父类的属性
print son.var2  # 调用爷爷类的属性

print '- 实例调用父类+爷爷类方法 - '
print son.get_var1()
print son.get_var2()
