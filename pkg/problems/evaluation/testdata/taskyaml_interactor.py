#!/usr/bin/python3
import sys
a,b = map(int, input().split())
with open(sys.argv[1], 'w') as f:
    print(a,b,file=f)
res = 0
with open(sys.argv[2], 'r') as f:
    res = int(f.readline())

if res == a+b:
	print(1.0)
	print('correct', file=sys.stderr)
else:
	print(0.0)
	print('incorrect', file=sys.stderr)
