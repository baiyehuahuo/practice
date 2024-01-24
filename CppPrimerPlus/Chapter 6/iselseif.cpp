#include<iostream>
const int Fave = 27;

int main()
{
    using namespace std;
    int n;
    cout << "Enter a number in the range 1-100 to find my favorite number: ";
    do 
    {
        cin >> n;
        if (n > Fave)
            cout << "Too high -- guess again: ";
        else if (n < Fave)
            cout << "Too low -- guess again: ";
        else
            cout << n << " is right!" << endl;
    }while(n != Fave);
    return 0;
}