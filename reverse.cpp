#include<bits/stdc++.h> 
using namespace std;

int main(){
   
   //"Welcome to this Javascript Guide!" should be become "emocleW ot siht tpircsavaJ !ediuG"
    string s ; 
    cin>>s ; 
    string ans = "" ;  
     string temp ; 
    for (int i = 0; i < s.length(); i++) { 
        if (s[i] == ' ') { 
            ans = temp + ' ' + ans; 
            temp = "";   
        } else {
            temp = s[i] + temp;
        }
    }  

    ans = temp + ' ' + ans; 

    cout << ans << endl; 
    return 0 ;
}
