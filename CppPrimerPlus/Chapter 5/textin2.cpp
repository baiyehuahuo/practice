#include<iostream>
int main()
{
    using namespace std;
    cout << "Enter characters; enter # to quic: " << endl;
    
    int count = 0;
    char ch;
    cin.get(ch);
    while(ch != '#')
    {
        cout << ch;
        count++;
        cin.get(ch);
    }
    cout << endl << count << " characters read\n";
    return 0;
}