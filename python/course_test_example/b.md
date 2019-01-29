### 列表相关

#### 过滤器 first

返回模版变量中第一个值（模版变量必须是个序列，包含字符串，列表，元组）。

范例：

    {{ value|first }}

value = ['a', 'b', 'c'], 输出就是 'a'.


#### 过滤器 last

返回一个列表的最后一个元素.

范例:

    {{ value|last }}

value = ['a', 'b', 'c', 'd'], 输出就是 "d".


#### 过滤器 join

连接一个列表成为一个字符串，和 Python 的 str.join(list) 一样的效果。

范例:

    {{ value|join:" // " }}

value = ['a', 'b', 'c'], 输出就是 "a // b // c".

#### 过滤器 random

返回一个列表中元素随机的一个。

范例:

    {{ value|random }}

value = ['a', 'b', 'c', 'd'], 输出可能就是 "b".

#### 过滤器 slice

返回一个列表的切片。

范例:

    {{ value|slice:":2" }}

value = ['a', 'b', 'c'], 输出就是 ['a', 'b'].
