s=input()
def brute_force(s):
    n=len(s)    
    pi=[0]*n
    for i in range(n):
        for L in range(i,0,-1):
            if(s[:L]==s[i-L+1:i+1]):
                pi[i]=L
                break
    return pi

a=brute_force(s)
print(a)

