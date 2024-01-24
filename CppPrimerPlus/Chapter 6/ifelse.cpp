#include<iostream>
int main()
{
    using namespace std;
    char ch;
    cout << "Type, and I shall repeat." << endl;
    cin.get(ch);
    while(ch != '.')
    {
        if (ch == '\n')
            cout << ch;
        else
            cout << ++ch;
        cin.get(ch);
    }
    cout << "\nPlease excuse the slight confusion." << endl;
    return 0;
}