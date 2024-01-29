#include<iostream>
int main()
{
    using namespace std;
    const int Max = 5;
    double fish[Max];
    cout << "Please enter the weights of your fish." << endl;
    cout << "You may enter up to " << Max << " fish <q to terminate>." << endl;
    double tmpFish;
    double sumFish = 0;
    int i;
    for (i = 0; i < Max; i++ )
    {
        cout << "fish #" << i+1 << ": ";
        if (!(cin >> tmpFish)) {
            break;
        }
        fish[i] = tmpFish;
        sumFish += tmpFish;
    }
    cout << sumFish / double(i) << " = average weight of " << i << " fish" << endl;
    cout << "Done" << endl;
    return 0; 
}