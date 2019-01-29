#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_breadth.py


class Parent(object):
    def __init__(self):
        print "实例化父类"


class Son(Parent):
    def __init__(self):
        print "实例化子类"
        super(Son, self).__init__()


son = Son()
