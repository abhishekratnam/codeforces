def sort012(arr,n):
    key = {0:0,1:0,2:0}
    for i in arr:
        for j in key.keys():
            if i==j:
                key[j] += 1
    a = []
    for i in range(0,3):
        k = key[i]
        while k:
            a.append(i)
            k -= 1
    return a
n = 2
arr = [0,2,1,2,0]
print(*(sort012(arr,n)))
