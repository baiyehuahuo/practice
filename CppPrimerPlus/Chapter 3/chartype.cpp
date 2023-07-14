// the char type
#include <iostream>

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    char ch;

    cout << "Enter a character: " << endl;
    cin >> ch;
    cout << "Hola! Thank you for the " << ch << " character." << endl;
    return 0;
}