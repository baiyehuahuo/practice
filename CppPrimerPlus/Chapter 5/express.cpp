#include<iostream>
int main() 
{
    using namespace std;
    int x;

    cout << "The expression x = 100 has the value 100" << endl;
    cout << (x=100) << endl;
    cout << "The expression x < 3 has the value " << (x < 3) << endl;
    cout << "The expression x > 3 has the value " << (x > 3) << endl;
    cout.setf(ios_base::boolalpha);
    cout << "The expression x < 3 has the value " << (x < 3) << endl;
    cout << "The expression x > 3 has the value " << (x > 3) << endl;
    return 0;
}