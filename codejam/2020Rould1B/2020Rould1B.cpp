#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void printResult(int casenum, vector<pair<int, int>>& res) {
	cout << "Case #" << casenum << ": " << res.size() << endl;
	for (auto& val : res)
		cout << val.first << " " << val.second << endl;
}

struct point {
	int r, s;
	point* prev, *next;
	point(int _r, int _s) : r(_r), s(_s), prev(nullptr), next(nullptr) {}
};

point* append(point* tail, int r, int s) {
	point* newtail = new point(r, s);
	if (tail != nullptr) {
		tail->next = newtail;
		newtail->prev = tail;
	}
	return newtail;
}

int countPoint(point* s, point* e) {
	int result = 0;
	while (s != e)
		s = s->next;
	return result;
}

void printlist(point* head) {
	cerr << "list" << endl;
	for (point* p = head; p != nullptr; p = p->next)
		cerr << p->r << " " << p->s << endl;
	cerr << "-----" << endl;
}

point* swap(point* fh, point* ft, point* bh, point* bt) {
	//cerr << "swap" << endl << fh->r << ":" << fh->s << endl << ft->r << ":" << ft->s << endl << bt->r << ":" << bt->s << endl;
	ft->next = bt->next;
	ft->next->prev = ft;
	bt->next = fh;
	fh->prev = bt;
	bh->prev = nullptr;
	return bh;
}

void solve(int casenum) {
	int r, s, n;
	cin >> r >> s;
	n = r * s;
	point* head = nullptr;
	point* tail = nullptr;
	vector<point*> record;
	for (int i = 0; i < s; i++) {
		for (int j = 0; j < r; j++) {
			tail = append(tail, j, i);
			if (head == nullptr)
				head = tail;
			record.push_back(tail);
		}
	}
	vector<pair<int, int>> result;
	int sortedcount = 1;
	while (sortedcount < n) {
		//printlist(head);
		if (tail->prev->r != tail->r && sortedcount % s != 0) {
			point* ft;
			int count = 1;
			for (ft = head; ft->r != tail->r; ft = ft->next) {
				count++;
			}
			head = swap(head, ft, ft->next, tail->prev);
			result.push_back(pair<int, int>(count, n - count - sortedcount));
		}
		tail = tail->prev;
		//cout << "new tail:" << tail->r << ":" << tail->s << endl;
		sortedcount++;
	}
	//printlist(head);
	printResult(casenum, result);
	for (auto val : record)
		delete val;
}

int main() {
	int T;
	cin >> T;
	for (int i = 1; i <= T; i++)
		solve(i);
	return 0;
}