#include<bits/stdc++.h> 
using namespace std;

int main(){
   
   //"Welcome to this Javascript Guide!" should be become "emocleW ot siht tpircsavaJ !ediuG"
    string s ; 
    cin>>s ; 
    string ans = "" ;  
     string temp ; 
    for(int i=0;i<s.length();i++){ 
    	temp  = s[i] +  temp  ;  
          if (s[i] == ' '){ 
          	i++;
           	 ans = ans + temp ; 
           	ans = ans + ' ' ; 
           	temp = "" ;   
           } 
          
      
    }  
    ans = ans + temp ;  
  
    for(int i=0;i<ans.length();i++){
    	cout<<ans[i] ;
    }
    //return ans ; 
    return 0 ;
}