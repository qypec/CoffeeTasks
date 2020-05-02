#include <iostream>
#include <map>
#include <vector>

using namespace std;

void						next(vector<vector<string>>& list, int del)
{
	int						newBack = list.size() - del - 1;

	if (del > 0)
	{
		for (auto i = list.size() - 1; i > newBack; i--)
		{
			for (auto t : list[i])
				list[newBack].push_back(t);
		}
	}
	list.resize(list.size() - del);
}

int							main()
{
	int						Q, dayNumber, monthCounter;
	string					operation, task;
	vector<int>				monthDay = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	vector<vector<string>>	list(monthDay[0]);

	cin >> Q;
	monthCounter = 0;
	for (auto i = 0; i < Q; i++)
	{
		cin >> operation;
		if (operation == "ADD")
		{
			cin >> dayNumber >> task;
			list[dayNumber - 1].push_back(task);
		}
		else if (operation == "DUMP")
		{
			cin >> dayNumber;
			cout << list[dayNumber - 1].size() << " ";
			for (auto t : list[dayNumber - 1])
				cout << t << " ";
			cout << endl;
		}
		else if (operation == "NEXT")
		{
			next(list, monthDay[monthCounter] - monthDay[(monthCounter + 1) % 12]);
			monthCounter = (monthCounter + 1) % 12;
		}
	}
}
