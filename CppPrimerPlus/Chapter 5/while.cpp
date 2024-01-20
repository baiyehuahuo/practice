#include<iostream>
int main()
{
    using namespace std;
    cout << "Your first name, please: ";
    char first_name[20];
    cin >> first_name;
    cout << "Here is your name, verticalized and ASCIIized: " << endl;
    int i = 0;
    while (first_name[i] != '\0')
    {
        cout << first_name[i] << ": " << int(first_name[i]) << endl;
        i++;
    }
    return 0;
}