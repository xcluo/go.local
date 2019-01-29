#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <xcluo@thstack.com>


import code
import readline
import rlcompleter

# do something here

vars = globals()
vars.update(locals())
readline.set_completer(rlcompleter.Completer(vars).complete)
readline.parse_and_bind("tab: complete")
shell = code.InteractiveConsole(vars)
shell.interact()
