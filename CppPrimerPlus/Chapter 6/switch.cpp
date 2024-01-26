#include<iostream>
using namespace std;
void show_menu();
void report();
void comfort();
int main()
{
    show_menu();
    int choice;
    cin >> choice;
    while (choice != 5)
    {
        switch (choice)
        {
        case 1:
            cout << "\a\n";
            break;
        case 2:
            report();
            break;
        case 3:
            cout << "The boss was in all day." << endl;
            break;
        case 4:
            comfort();
            break;
        default:
            cout << "That's not a choice." << endl;
        }
        show_menu();
        cin >> choice;
    }
    cout << "Bye!" << endl;
    return 0;
}

void show_menu() 
{
    cout <<"Please enter 1, 2, 3, 4, or 5:\n"
         <<"1) alarm        2) report\n" 
         <<"3) alibi        4) comfort\n"
         <<"5) quit" << endl;
}

void report()
{
    cout <<"It's been an excellent week for business." << endl;
    cout <<"Sales are up 120%. Expenses are down 35%." << endl;
}

void comfort()
{
    cout << "You employees think you are the finest CEO in the dustry." << endl;
    cout << "The board of directions think you are the finest CEO in the industry." << endl;
}