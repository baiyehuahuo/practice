#include<iostream>
int main()
{
    using namespace std;
    char ch;
    int count = 0;
    cout << "Enter characters; enter # to quic: " << endl;
    cin >> ch;
    while( ch != '#')
    {
        cout << ch;
        count++;
        cin >> ch;
    }     
    cout << endl << count << " characters read" << endl;
    return 0;
}