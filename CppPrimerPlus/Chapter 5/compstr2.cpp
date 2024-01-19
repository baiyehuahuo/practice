#include<iostream>
#include<cstring>
int main()
{
    using namespace std;
    string word = "?ate";
    for (char ch = 'a'; ch <= 'z' && word != "mate"; ch++)
    {
        cout << word << endl;
        word[0] = ch;
    }
    cout << "After loop ends, word is mate." << endl;
    return 0;
}