#include<iostream>
int main()
{
    using namespace std;
    int a, b;
    cout << "Enter two integers: ";
    cin >> a >> b;
    cout << "The larger of " << a << " and " << b << " is " << (a > b ? a : b) << endl;
    return 0;
}