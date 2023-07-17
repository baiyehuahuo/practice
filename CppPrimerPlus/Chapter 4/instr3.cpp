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
    cin.get(name, ArSize).get();
    cout << "Enter your favorite dessert: " << endl;
    cin.get(dessert, ArSize).get();
    cout << "I have some delicious " << dessert << " for you, " << name << endl;
    return 0;
}