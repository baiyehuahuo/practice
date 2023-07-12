// converts stone to pound
#include <iostream>

int stonetolb(int);
int main()
{
    using std::cin;
    using std::cout;
    using std::endl;

    int stone;

    cout << "Enter the weight in stone: ";
    cin >> stone;

    int pounds = stonetolb(stone);
    cout << stone << " stone = " << pounds << " pounds." << endl;
    return 0;
}

int stonetolb(int stone)
{
    return 14 * stone;
}