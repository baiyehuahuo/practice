#include<iostream>
enum {red, orange, yellow, green, blue, violet, indigo};

int main()
{
    using namespace std;
    cout << "Enter color code 0-6: ";
    int color_code;
    cin >> color_code;
    while(color_code >= red && color_code <= indigo)
    {
        switch (color_code)
        {
        case red:
            cout << "Her lips were red." << endl;
            break;
        case orange:
            cout << "Her hair was orange." << endl;
            break;
        case yellow:
            cout << "Her shoes were yellow." << endl;
            break;
        case green:
            cout << "Her nails were green." << endl;
            break;
        case blue:
            cout << "Her sweatsuit was blue." << endl;
            break;
        case violet:
            cout << "Her eyes were violet." << endl;
            break;
        case indigo:
            cout << "Her mood was indigo." << endl;
        default:
            break;
        }
        cout << "Enter color code (0-6): ";
        cin >> color_code;
    }
    cout << "Bye" << endl;
    return 0;
}