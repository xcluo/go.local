#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_depth.py


# Coding here
class Grandpa(object):
    def __init__(self):
        print "实例化Grandpa"


class Parent(object):
    def __init__(self):
        print "实例化Parent"


class Son(Parent):
    var = '子类的属性'

    def __init__(self):
        super(Son, self).__init__()


son = Son()
