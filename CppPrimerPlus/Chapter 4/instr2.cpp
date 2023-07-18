// reading more than one word with getline
#include <iostream>

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    const int ArSize = 20;
    char name[ArSize], dessert[ArSize];

    cout << "Enter your name: " << endl;
    cin.getline(name, ArSize);
    cout << "Enter your favorite dessert: " << endl;
    cin.getline(dessert, ArSize);
    cout << "I have some delicious " << dessert << " for you, " << name << endl;
    return 0;
}