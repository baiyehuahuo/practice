// some C++ arithmetic
#include <iostream>

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    float num1, num2;

    cout.setf(std::ios_base::fixed, std::ios_base::floatfield);
    cout << "Enter a number: ";
    cin >> num1;
    cout << "Enter another number: ";
    cin >> num2;

    cout << "num1 = " << num1 << "\nnum2 = " << num2 << endl;
    cout << "num1 + num2 = " << num1 + num2 << endl;
    cout << "num1 - num2 = " << num1 - num2 << endl;
    cout << "num1 * num2 = " << num1 * num2 << endl;
    cout << "num1 / num2 = " << num1 / num2 << endl;
    return 0;
}