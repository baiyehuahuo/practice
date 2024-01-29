#include<iostream>

int main()
{
    using namespace std;
    const int Max = 5;
    int sumScores = 0, tmpScore = 0;

    cout << "Please enter your golf scores." << endl;
    cout << "You must enter " << Max << " rounds." << endl;
    for (int i = 0; i < Max; i++)
    {
        cout << "round #" << i+1 << ": ";
        while(!(cin >> tmpScore))
        {
            cin.clear();
            while (cin.get() != '\n')
                continue; // \n 似乎不会被clear随意清除 但同时又会被 cin >> tmpScore 判定为结尾
            cout << "Please enter a number: ";
        }
        sumScores += tmpScore;
    }
    cout << double(sumScores) / double(Max) << " = average score " << Max << " rounds" << endl;
    return 0;
}