// a simple structure
#include <iostream>
struct inflatable
{
    /* data */
    char name[20];
    float volumn;
    double price;
};

int main()
{
    using std::cout;
    using std::endl;

    inflatable guest = {
        "Glorious Gloria",
        1.88,
        29.99};

    inflatable pal = {
        "Audacious Arthur",
        3.12,
        32.99};

    cout << "Expand your guest list with " << guest.name << " and " << pal.name << "!" << endl;
    cout << "You can have both for $" << guest.price + pal.price << "!" << endl;
    return 0;
}