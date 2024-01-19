#include<iostream>
int main()
{
    using namespace std;
    cout << "The Amazing Accounto will sum and average five numbers for you." << endl;
    cout << "Please enter five values:" << endl;
    double num, sum;
    for (int i = 0; i < 5; i++)
    {
        cout << "Value " << i+1 << ": ";
        cin >> num;
        sum += num;
    }
    cout << "Five exquisite choices indeed! They sum to " << sum << " and average to " << sum / 5 << ".\n";
    cout << "The Amazing Accounto bids you adieu!" << endl;
    return 0;
}