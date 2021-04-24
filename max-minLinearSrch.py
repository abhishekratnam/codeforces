class pair:
    def __init__(self):
        self.min = 0
        self.max = 0
def getMazMin(arr, n):
    minmax = pair()
    if n==1:
        minmax.max = arr[0]
        minmax.min = arr[0]
        return minmax
    if arr[0] > arr[1]:
        minmax.max = arr[0]
        minmax.min = arr[1]
    else:
        minmax.min = arr[1]
        minmax.max = arr[0]
    for i in range(2,n):
        if arr[i]>minmax.max:
            minmax.max = arr[i]
        elif arr[i]<minmax.min:
            minmax.min = arr[i]
    return minmax

if __name__=="__main__":
    arr = [100,11,445,1]
    arr_size = 6
    minmax = getMazMin(arr, arr_size)
    print("Minimum element is", minmax.min)
    print("Maximum element is", minmax.max)
    