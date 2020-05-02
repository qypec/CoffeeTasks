#include <iostream>
#include <map>
#include <vector>

using namespace std;

int							main()
{
    int                     Q;
    string                  operation, country, country_new, capital;
    map<string, string>     capitalList;

    cin >> Q;

	for (auto i = 0; i < Q; i++)
	{
		cin >> operation;
		if (operation == "CHANGE_CAPITAL")
		{
            cin >> country >> capital;
            if (capitalList.count(country) == 0)
            {
                capitalList[country] = capital;
                cout << "Introduce new country " << country << " with capital " << capital << endl;
            }
            else if (capitalList[country] == capital)
                cout << "Country " << country << " hasn't changed its capital" << endl;
            else
            {
                cout << "Country " << country << " has changed its capital from " << capitalList[country] << " to " << capital << endl;
                capitalList[country] = capital;
            }
		}
		else if (operation == "RENAME")
		{
            cin >> country >> country_new;
            if (country == country_new || capitalList.count(country) == 0 || capitalList.count(country_new) != 0)
                cout << "Incorrect rename, skip" << endl;
            else
            {
                cout << "Country " << country << " with capital " << capitalList[country] << " has been renamed to " << country_new << endl;
                capital = capitalList[country];
                capitalList.erase(country);
                capitalList[country_new] = capital;
            }

		}
		else if (operation == "ABOUT")
		{
            cin >> country;
            if (capitalList.count(country) == 0)
                cout << "Country " << country << " doesn't exist" << endl;
            else
                cout << "Country " << country << " has capital " << capitalList[country] << endl;
		}
        else if (operation == "DUMP")
		{
            if (capitalList.size() == 0)
                cout << "There are no countries in the world";
            for (auto item : capitalList)
                cout << item.first << "/" << item.second << " ";
            cout << endl;
		}
	}
}
