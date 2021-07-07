def max(n,m):
    if n > m:
        return n,m
    return m,n
n = int(input())
for _ in range(n):
    arr = tuple(input().split())
    red,blue,dif = int(arr[0]),int(arr[1]),int(arr[2])
    m,n= max(red,blue)
    diff = int(m)-int(n)
    quo = diff/n
    rem = diff%n
    if quo+rem <= int(diff):
        if int(red)==int(blue):
            print("Yes")
    
    print("No")


