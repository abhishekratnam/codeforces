n = int(input())
cal_ans = []
for i in range(n):
    cal_ans.append(int(input()))
for i in cal_ans:
    j = i%10    
    digits = 10*(j-1)
    no = i
    count = 0
    while no != 0:
        no //=10
        count += 1
    s = 0
    while count !=0:
        s += count
        count -= 1
    print(digits+s)

        


    



    





    



    



