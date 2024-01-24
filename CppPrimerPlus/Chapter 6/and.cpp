#include<iostream>
const int Arsize = 6;
int main()
{
    using namespace std;
    float naaq[Arsize];
    cout << "Enter the NAAQs (New Age Awareness Quotients) of your neighbors." << endl;
    cout << "Program terminates when you make 6 entries or enter a negative value." << endl;
    int i = 0;
    float tmp ;
    cout << "First value: " ;
    cin >> tmp;
    while(i < Arsize && tmp > 0)
    {
        naaq[i] = tmp;
        i++;
        if (i < Arsize)
        {
            cout << "Next value: ";
            cin >> tmp;
        }
    }
    if (i == 0)
        cout << "No data -- bye" << endl;
    else
    {
        cout <<"Enter your NAAQ: ";
        cin >> tmp;
        int count = 0;
        for (int j = 0; j < i; j++)
            if (naaq[j] > tmp)
                count++;
        cout << count << " of your neighbors have greater awareness of the New Age than you do." << endl;
    }
    return 0;
}