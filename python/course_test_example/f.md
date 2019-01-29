

### 定义模版继承

模版继承是很 easy 的技术，但很强大，我们看个范例：


    <!DOCTYPE html>
    <html lang="en">
    <head>
        <link rel="stylesheet" href="style.css" />
        <title>{% block title %}My amazing site{% endblock %}</title>
    </head>
    
    <body>
        <div id="header">
            <ul>
                <li><a href="/">Home</a></li>
                <li><a href="/blog/">Blog</a></li>
            </ul>
        </div>
    
        <div id="content">
            {% block content %}{% endblock %}
        </div>

        <div id="footer">
            Welcome, and contact by contact@fuvism.com.
        </div>
    </body>
    </html>


这个文件一般被命名为 base.html，称为父模版。不同于普通的 HTML 模版， 它多了：

    {% block title %} {% endblock %}
    {% block content %} {% endblock %}


block 和 endblock 标签组成一对，title 是名称，在这里定义了两个名称为 title, content 的区域，这也是我们这个父模版作为“骨架”模版的关键。

通过 block 标签 “挖了两个坑 title, content"，这样我们的子模版（除父模版外其他模版），就可以将自己的内容填充到这两个坑中，而不需要再次编写 `<html>/<head>/header/footer` 等部分。

`这里的父模版保存在模版目录中，也就是 templates/base.html`


### 使用模版继承

子模版只需要这样使用即可：

    {% extends "base.html" %}
    {% block title %}
       我的Title，会替换掉父模版的 title
    {% endblock %}
    {% block content %}
       子页面需要展示的内容，会替换掉父模版中相同区域。
    {% endblock %}

Done!

