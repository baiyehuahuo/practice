// storing strings in an array
#include <iostream>
#include <cstring>

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    const int Size = 15;
    char name1[Size];
    char name2[Size] = "C++owboy";

    cout << "Howdy! I'm " << name2 << "! What's your name?" << endl;
    cin >> name1;
    cout << "Well, " << name1 << ", your name has " << strlen(name1) << " letters and is stored"
         << "in an array of " << sizeof(name1) << " bytes." << endl;
    cout << "Your initial is " << name1[0] << ".\n";
    name2[3] = '\0';
    cout << "Here are the first 3 characters of my name: " << name2 << endl;
    return 0;
}