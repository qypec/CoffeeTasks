
#include <iostream>
#include <map>
#include <vector>

using namespace std;

int                                 main()
{
    int                             Q, n, number_counter;
    map<vector<string>, int>        route;
    vector<string>                  stop_list;
    string                          stop;

    cin >> Q;
    number_counter = 1;
    for (auto i = 0; i < Q; i++)
    {
        stop_list.clear();
        cin >> n;
        for (auto j = 0; j < n; j++)
        {
            cin >> stop;
            stop_list.push_back(stop);
        }
        if (route.count(stop_list) != 0)
            cout << "Already exists for " << route[stop_list] << endl;
        else
        {
            route[stop_list] = number_counter++;
            cout << "New bus " << route[stop_list] << endl;
        }
    }
}
