// using escape sequences
#include <iostream>

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    cout << "\aOperation \"HyperHype\" is now activated!" << endl;
    cout << "Enter your agent code: ________\b\b\b\b\b\b\b\b";

    int code;
    cin >> code;
    cout << "\aYou entered " << code << "..." << endl;
    cout << "\aCode verified! Proceed with Plan Z3" << endl;
    return 0;
}