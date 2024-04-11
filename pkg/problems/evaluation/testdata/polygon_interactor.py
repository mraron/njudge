#!/usr/bin/python3
import sys
with open(sys.argv[1], 'r') as f:
    a,b = list(map(int, f.readline().split()))

print(a,b)
res = input()
with open(sys.argv[2], 'w') as f:
    f.write(res)