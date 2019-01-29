
查找算法：
    条件： 有序序列（升， 降）
    顺序查找：

1. 值是存在的：
    for val in list:
        if input == input:
            print 'index=', list.index(val)

    for i in range(len(list)):
        if input ==  list[i]
            print 'index=', i

    i = 0
    for val in list:
        if input ==  val]
            print 'index=', i
        i += 1

2. 优化的顺序查找：
    值不存在的情况下，求前置和后置索引：

index = None
count = len(old_list) - 1
i = 0
for var in old_list:
    if var == input_data:
        index =  i
    if input_data > var:
        if i+1 <= count and input_data < old_list[i+1]:
            prev_index = i; next_index = i+1

    if input_data < var:
        if i-1 >= 0 and input_data > old_list[i-1]:
            prev_index = i-1; next_index = i
        else:
    
    i += 1




















old_list = [1,3,4,6,7,9,10]

input_data = 3
index =  None
prev_index = None
next_index = None


1. 判断是否相等, 返回当前的index
2. 如果比最小值（左边）小，返回 p_index=None, next_index=0
3. 如果比最大值（右边）大，返回 p_index=最后一个值的index, next_index=None
4. 如果在中间的情况：
    遍历列表，找 < <
if input <= old_list[0]
    index = 0
    prev_index = None
    next_index = 0
elif input >= old_list[-1]
    prev_index = len(list)-1
    next_index = None

i = 0 
for val in old_list:
    if val == input:
        index = old_list.index(val)
    if  input > val:
        if input < old_list[i+1]
            prev_index = i
            next_index = i+1
    if  input < val:
        index = old_list.index(val)
    i += 1
    
        
遍历old_list的每个value, 找到值的index:
    1. 拿到以第一个value：
        判断输入的 比 value 大：
                        是否比后面的值小： 是：  return p_ = i, n = i+1 break
                                            否： 继续
                小：
                        是否比前一个值大： 是： rerturn p_= , n = , break
                                           否：  
