#include<iostream>
int main()
{
    using namespace std;
    int step;
    cout << "Enter an integer: ";
    cin >> step;
    cout << "Counting by " << step << "s:" << endl;
    for (int i = 0; i < 100; i+=step) {
        cout << i << endl;
    }
    return 0;
}