#include<iostream>
int main()
{
    using namespace std;
    // cout << "Enter characters; enter # to quic: " << endl;
    
    int count = 0;
    char ch;
    while(cin.get(ch))
    {
        cout << ch;
        count++;
    }
    cout << endl << count << " characters read\n";
    return 0;
}