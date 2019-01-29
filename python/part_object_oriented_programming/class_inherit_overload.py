#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Learn in Cloud <contact@pythonpie.com>
# Filename: class_inherit_by_depth.py


# Coding here
class Prod:
    def __init__(self, value):
        self.value = value

    def __call__(self, other):
        return self.value * other


p = Prod(2)     # call __init__
print p(2)      # call __call__
print p(4)
