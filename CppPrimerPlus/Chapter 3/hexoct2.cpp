// display values in hex and octal
#include <iostream>

int main()
{
    using std::cout;
    using std::endl;
    int value = 42;

    cout << "Monsieur cuts a striking figure!" << endl;
    cout << "chest = " << value << " (decimal for 42)" << endl;
    cout << std::hex;
    cout << "waist = " << value << " (hex for 42)" << endl;
    cout << std::oct;
    cout << "inseam = " << value << " (oct for 42)" << endl;
    return 0;
}