
old_list = [1,3,4,6,7,9,10]

index = None; prev_index=None; next_index=None
count = len(old_list) - 1
input_data = int(raw_input('Please input your number:'))

i = 0
if input_data < old_list[0]:
    prev_index = None
    next_index = 0
if input_data  old_list[-1]:
    prev_index = count
    next_index = None

for var in old_list:
    if var == input_data:
        index =  i
        break

    if input_data > var:
        if i+1 <= count and input_data < old_list[i+1]:
            prev_index = i; next_index = i+1
            break

    if input_data < var:
        if i-1 >= 0 and input_data > old_list[i-1]:
            prev_index = i-1; next_index = i
            break
    i += 1


print 'index=', index
print 'prev_index = ', prev_index
print 'next_index = ', next_index
