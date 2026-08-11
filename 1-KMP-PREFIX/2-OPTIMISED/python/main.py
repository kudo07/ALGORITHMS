def compute_KMP(str):  
    pi=[0]* len(str)
    j=0
    
    for i in range(1,len(str)):
        
        while( j>0 and s[i] != s[j] ):
            j=pi[j-1]
        
        if(s[i]==s[j]){
            j+=1
        }
        pi[i]=j
    return pi