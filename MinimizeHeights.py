def getMinDiff(arr, n, k):
    # code here
    min_ele = []
    for i in range(n):
        a = arr[i] + k
        s = 0
        if  arr[i] - k > -1:
            s = arr[i] - k
        min_ele.append(min(a,s))
    for i in min_ele:
        if i==0:
            min_ele.remove(i)
    return max(min_ele)-min(min_ele)
arr= [1,5,8,10]
n = len(arr) 
k = 2
n = getMinDiff(arr, n, k)
print(n)
        