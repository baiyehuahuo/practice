#include<iostream>
int main()
{
    using namespace std;
    cout << "This program may reformat your hard dist and destroy all your data." << endl;
    cout << "Do you wish to continue?<y/n>" ;
    char ch;
    cin >> ch;
    if (ch == 'y' || ch == 'Y')
        cout <<"You were warned!\a\a" << endl;
    else if (ch =='n' || ch == 'N')
        cout <<"A wise choice ... bye" << endl;
    else
        cout << "That wasn't a y or n! Apparently you can't follow instructions, so I'll trash your disk anyway.\a\a\a";
    return 0;
}