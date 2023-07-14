// type changes on initialization
#include <iostream>

int main()
{
    using std::cout;
    using std::endl;

    cout.setf(std::ios_base::fixed, std::ios_base::floatfield);

    float tree = 3;
    int guess(3.9832);
    int debt = 7.2E12;

    cout << "tree = " << tree << endl;
    cout << "guess = " << guess << endl;
    cout << "debt = " << debt << endl;
    return 0;
}