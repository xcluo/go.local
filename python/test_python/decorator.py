#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>


def test_desc(func):
    def _deco(*args, **kwargs):
        print 'desc = 1'
        cc = func(*args, **kwargs)
        print 'desc = 2'
        if 'c' in kwargs and kwargs['c'] is True:
            print 'remove session'
        return cc
    return _deco


@test_desc
def test(a, b, c=None):
    print 'a * b = ', a * b
    return 'OK'


print test(1,2, c = True)

print '------------'

print test(1,2)
