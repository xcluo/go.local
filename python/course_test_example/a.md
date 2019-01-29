### 字符，数字相关

#### 过滤器 lower

转化字符串中的字母为小写.

范例:

    {{ value|lower }}

value = 'THIS Is a Test!', 输出为 'this is a test!'.

#### 过滤器 upper

转化字符串中的字母为大写.

范例:

    {{ value|upper }}

value = 'this is a test!', 输出为 'THIS IS A TEST!'.


#### 过滤器 add

增加一个值给模版变量

范例：

    {{ value|add:"2" }}

如果 value = 4，那么输出就是 4 + 2 = 6。

需要注意的是，如果 value 和后面的参数，都是数字，就是加法，如果不是，过滤器 add 将尝试把两个值连接起来。


范例：

    {{ first|add:second }}

first = [1, 2, 3] and second = [4, 5, 6], 那么输出就是 [1, 2, 3, 4, 5, 6].


#### 过滤器 cut

移除模版变量中所有指定的字符串。

范例：

    {{ value|cut:" " }}


value = "String with spaces", 输出就是 "Stringwithspaces".

