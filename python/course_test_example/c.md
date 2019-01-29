### 序列相关


#### 过滤器 length

返回长度，包括字符串和列表的所有序列都适用。

范例:

    {{ value|length }}

value = ['a', 'b', 'c', 'd']，输出为 4.

value =  "abcd", 输出为 4.

`Django 1.8 后，如果变量 value 未定义，则返回 0`

`Django 1.8 前的版本，返回空字符`

