// show hex and octal literals
#include <iostream>
int main()
{
    using std::cout;
    using std::endl;
    int chest = 42;
    int waist = 0x42;
    int insream = 042;

    cout << "Monsieur cuts a striking figure!" << endl;
    cout << "chest = " << chest << " (42 in decimal)" << endl;
    cout << "waist = " << waist << " (0x42 in hex)" << endl;
    cout << "insream = " << insream << " (042 in octal)" << endl;
    return 0;
}