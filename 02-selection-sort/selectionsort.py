lists = [10,20,30,5,14,66,8,41,55]

def find_biggest(lists):
    biggest = lists[0]
    biggest_index = 0

    for i in range(1,len(lists)):
        if(lists[i] > biggest):
            biggest = lists[i]
            biggest_index = i
    return biggest_index

def selection_sort(lists):
    new_list = []
    for i in range(len(lists)):
        biggest = find_biggest(lists)
        new_list.append(lists.pop(biggest))
    return new_list
    

print(selection_sort(lists))