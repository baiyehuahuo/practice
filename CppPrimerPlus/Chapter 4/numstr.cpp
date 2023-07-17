// following number input with line input
#include <iostream>
int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    cout << "What year your house built?" << endl;
    int year;
    (cin >> year).get();

    cout << "What is its street address?" << endl;
    char address[80];
    cin.getline(address, 80);

    cout << "Year built: " << year << endl;
    cout << "Address: " << address << endl;
    cout << "Done!" << endl;
    return 0;
}