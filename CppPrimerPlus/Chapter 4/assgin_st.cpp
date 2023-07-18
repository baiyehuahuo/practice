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

    inflatable bouquet = {"sunflowers", 0.20, 12.49};
    inflatable choice;

    cout << "bouquet: " << bouquet.name << " for $" << bouquet.price << endl;
    choice = bouquet;
    cout << "choice: " << choice.name << " for $" << choice.price << endl;
    return 0;
}