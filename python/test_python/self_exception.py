#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>



class WarningException(Exception):
    def __init__(self, value='warning exception'):
        self.status = 1
        self.message = value


class TestException(Exception):
    def __init__(self, value='test exception'):
        self.status = 2
        self.message = value




try:
    # raise WarningException('aaaaaaaa')
    raise TestException('aaaaaaaa')

except (WarningException, TestException), e:
    print 'aaaaaaaaaa = ', e.status, '   bbbbbbbbb=', e.message

