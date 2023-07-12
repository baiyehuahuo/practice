// defining your own function
#include <iostream>

void simon(int);

int main()
{
    using std::cin;
    using std::cout;
    using std::endl;
    int count;
    simon(3);
    cout << "Pick an integer: ";
    cin >> count;
    simon(count);
    cout << "Done!" << endl;
    return 0;
}

void simon(int n)
{
    using std::cout;
    using std::endl;
    cout << "Simon says touch your toes " << n << " times." << endl;
}
