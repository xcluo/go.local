

之前用for循环实现了输出100之内的素数，看看用while怎么实现！

    >>> i = 2
    >>> while (i < 100):
    >>>     j = 2
    >>>     while (j <= (i/j)):
    >>>         if not(i%j):  break
    >>>         j = j + 1
    >>>     if (j > i/j) : print i, "是素数"
    >>>     i = i + 1

###以上实例输出结果

    2 是素数
    3 是素数
    5 是素数
    7 是素数
    11 是素数
    13 是素数
    17 是素数
    19 是素数
    23 是素数
    29 是素数
    31 是素数
    37 是素数
    41 是素数
    43 是素数
    47 是素数
    53 是素数
    59 是素数
    61 是素数
    67 是素数
    71 是素数
    73 是素数
    83 是素数
    89 是素数
    97 是素数


嵌套循环输出2~100之间的素数：

    >>> print "2 is prime"
    >>> for i in range(3,100): 
    >>>     for j in range(2, i):
    >>>         if not i%j:
    >>>             break
    >>>         j = j + 1 
    >>>     if j > i/j:
    >>>         print i, " is prime"
    >>>     else:
    >>>         print i, " is not prime"
    >>>     i = i + 1
    >>> print "The end!"
    
