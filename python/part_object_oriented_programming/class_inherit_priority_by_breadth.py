#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_breadth.py


# Coding here
class Parent1(object):
    def __init__(self):
        print "实例化Parent1"


class Parent2(object):
    def __init__(self):
        print "实例化Parent2"


class Son(Parent1, Parent2):
    var = '子类的属性'

    def __init__(self):
        super(Son, self).__init__()


son = Son()
