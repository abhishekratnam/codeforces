def max(n,m):
    if n > m:
        return n,m
    return m,n
n = int(input())
for _ in range(n):
    arr = tuple(input().split())
    red,blue,diff = arr[0],arr[1],arr[2]
    m,n= max(red,blue)
    diff = int(m)-int(n)
    quo = diff/int(n)
    rem = diff%int(n)
    if quo+rem <= int(diff) or int(red)==int(blue)<=int(diff):
        print("Yes")
    else:
        print("No")


