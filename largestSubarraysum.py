def max(a,b):
    if a > b:
        return a 
    return b
def Kadane(arr, n):
    max_so_fr = arr[0]
    curr_max = arr[0]
    for i in range(1,n):
        curr_max = max(arr[i], curr_max+arr[i])
        max_so_fr = max(max_so_fr, curr_max)
    return max_so_fr

n = int(input())
arr = list(map(int,input().split()))
print(Kadane(arr,n))