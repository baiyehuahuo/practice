// using sqrt function
#include <iostream>
#include <cmath>
int main()
{
    using std::cin;
    using std::cout;
    using std::endl;
    double area, side;
    cout << "Enter the floor area, in square feet, of your home: ";
    cin >> area;
    side = sqrt(area);
    cout << "That's the equivalent of a square " << side << " feet to the side." << endl;
    cout << "How fascinating!" << endl;
    return 0;
}