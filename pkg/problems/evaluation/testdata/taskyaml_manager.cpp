#include <iostream>
#include <signal.h>
#include <cassert>
#include <fstream>

using namespace std;

int main(int argc, char **argv) {
  signal(SIGPIPE, SIG_IGN);

  ifstream fin("input.txt");
  ofstream to_sol(argv[2]);
  ifstream from_sol(argv[1]);

  int n;

  fin >> n;

  to_sol << n << endl;

  bool ok = true;


  string er2;

  for (int i = 0; i < n; i++) {
    long long a, b;
    fin >> a >> b;
    to_sol << a << ' ' << b << endl;

    from_sol >> er2;
    if (to_string(a+b) != er2) {
      ok=false;
    }
  }

    if(ok){
      printf("1.0\n");
      fprintf(stderr, "translate:success\n");
    } else {
        printf("0.0\n");
        fprintf(stderr, "translate:wrong\n");
    }

}