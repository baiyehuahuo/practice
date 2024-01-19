#include<iostream>
int main()
{
    using namespace std;
    string word;
    char tmp;
    cout << "Enter a word: ";
    cin >> word;

    for (int i = 0, j = word.size(); i < j; i++, j--) 
    {  
        tmp = word[i];
        word[i] = word[j];
        word[j] = tmp;
    }
    cout << word << endl;
    cout << "Done" << endl;
    return 0;
}