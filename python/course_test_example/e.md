### HTML 处理相关过滤器

#### 过滤器 title

转换一个字符串为标题格式（英文中，是每个单词的第一个字母大写）。

范例:

    {{ value|title }}

value = "my FIRST post", 输出为 "My First Post".

#### 过滤器 truncatechars

截取一个字符串指定的长度。如果字符串的长度大于截取的长度，自动附加 "..." 。


范例:

    {{ value|truncatechars:5 }}

截取 5 个字符。

value = "abcdef", 输出就是 "ab..."。

`注意，5个字符并不是 value 的 5 个字符，而是截取后整个的长度是 5 个。`

所以:

    {{ value|truncatechars:4 }}

相当于 3 个 ...，和 1 个 value 的字符，输出是 'a...':

    {{ value|truncatechars:4 }}

这样，输出只有三个点 "..."了，这个计算方法需要注意，不然容易产生歧义！

    {{ value|truncatechars:6 }}

输出竟然是 "abcdef"，而不是 “abc...“，在定义中已经说明，只有截取的长度超过值的长度，才去截取，而 value 本来就是 6 个字符，所以这里直接输出内容，不做截取！

对于中文，也支持：

    {{ value|truncatechars:6 }}

value = "这是一个测试"，输出就是 “这是一...”

`这是个常用技能，界面上无法展示所有内容时，大多在后面加上 ... 作为省略`

#### 过滤器 truncatewords

截取一个字符串指定个数的单词。

在英文单词(之间是有空格的)组成的字符串中，截取指定数量的单词，如果字符串单词数超过截取的，多出来的部分使用省略号。

在中文语句汇总，由于 truncatewords 按照空格识别单词，无法处理中文连贯的表达方式，可以使用 truncatechars 处理。


范例:

    {{ value|truncatewords:2 }}

这里指定 2， 作为截取的长度。

value = "Joel is a slug", 输出就是 "Joel is ...".

`不同于 truncatechars，truncatewords 不存在复杂的计算，`
`截取2个单词就是2个单词，超过的部分加 ...，不超过直接输出原文！`


由于中文的限制，对我们来说，truncatechars 要更实用的多。

#### 过滤器 safe

如果模版变量是个HTML，比如 `value = '<h1> Title </h1>'，正常情况下 Django 模版并不会将 value 当初 HTML 语句，而是直接输出。

如果模版是：

    <body>
        <h1> Test </h1>
        {{ value }}
    </body>

那么网页展示的效果是：

    Test
    <h1> Title </h1>

这里的模版变量原样输出，并没有像 `<h1> Test </h1>` 那样被渲染。

过滤器 safe，处理可以渲染的模版变量中的 HTML 代码。

    {{ value|safe }}

这时候的输出就是：

    Test
    Title

