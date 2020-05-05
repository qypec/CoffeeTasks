/*
    Реализуйте систему хранения автобусных маршрутов. Вам нужно обрабатывать следующие запросы:
        NEW_BUS bus stop_count stop1 stop2 ... — добавить маршрут автобуса с названием bus
        и stop_count остановками с названиями stop1, stop2, ...
        BUSES_FOR_STOP stop — вывести названия всех маршрутов автобуса, проходящих через остановку stop.
        STOPS_FOR_BUS bus — вывести названия всех остановок маршрута bus со списком автобусов,
        на которые можно пересесть на каждой из остановок.
        ALL_BUSES — вывести список всех маршрутов с остановками.
*/

#include <iostream>
#include <map>
#include <vector>

using namespace std;

int                                 main()
{
    int                             Q;
    map<string, vector<string>>     route, stops;
    string                          operation, bus, stop;
    int                             stop_count;


    cin >> Q;
	for (auto i = 0; i < Q; i++)
	{
		cin >> operation;
		if (operation == "NEW_BUS")
		{
            cin >> bus >> stop_count;
            for (auto i = 0; i < stop_count; i++)
            {
                cin >> stop;
                route[bus].push_back(stop);
                stops[stop].push_back(bus);
            }
		}
		else if (operation == "BUSES_FOR_STOP")
		{
            cin >> stop;
            if (stops.count(stop) == 0)
                cout << "No stop";
            for (auto item : stops[stop])
                cout << item << " ";
            cout << endl;
		}
		else if (operation == "STOPS_FOR_BUS")
		{
            cin >> bus;
            if (route.count(bus) == 0)
                cout << "No bus" << endl;
            else
            {
                for (auto stop_item : route[bus])
                {
                    cout << "Stop " << stop_item << ": ";
                    if (stops[stop_item].size() == 1)
                        cout << "no interchange";
                    else
                    {
                        for (auto bus_item : stops[stop_item])
                        {
                            if (bus_item == bus)
                                continue ;
                            cout << bus_item << " ";
                        }
                    }
                    cout << endl;
                }
            }
		}
        else if (operation == "ALL_BUSES")
		{
            if (route.size() == 0)
                cout << "No buses" << endl;
            else
            {
                for (auto item : route)
                {
                    cout << "Bus " << item.first << ": ";
                    for (auto stop : item.second)
                        cout << stop << " ";
                    cout << endl;
                }
            }
		}
	}
}
