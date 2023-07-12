// food processing program uses and displays a variable
#include <iostream>
int main()
{
    using std::cout;
    using std::endl;
    int carrots;
    carrots = 25;
    cout << "I have " << carrots << " carrots." << endl;
    cout << "Crunch, crunch, Now Ihava " << carrots - 1 << " carrots." << endl;
    return 0;
}