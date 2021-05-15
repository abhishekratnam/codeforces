def minJumps(arr, n):
        count = 0
        for i in arr:
            k = i
            if n%2 == 0:
                for j in range(k):
                    arr.pop(0)
            else:
                for j in range(k-1):
                    arr.pop(0)
            count += 1
        return count

# n = 11
# arr = [1 ,3 ,5 ,8 ,9 ,2 ,6 ,7 ,6 ,8 ,9]

n = 6
arr = [1 ,4 ,3 ,2 ,6 ,7]
ans = minJumps(arr,n)
print(ans)
