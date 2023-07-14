// uses % operator to convert lbs to stone
#include <iostream>
int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    const int Lbs_per_stn = 14;
    int lbs;

    cout << "Enter your weight in pounds: ";
    cin >> lbs;

    int stone = lbs / Lbs_per_stn;
    int pounds = lbs % Lbs_per_stn;

    cout << lbs << " pounds are " << stone << " stone, " << pounds << " pounds(s)." << endl;
    return 0;
}