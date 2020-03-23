"""
代码思路:
heapq动态获取
"""
import heapq
from io import StringIO

def efficent_readint():
    from sys import stdin
    s = StringIO()
    for c in stdin.read():
        if c >= '0' and c <= '9':
            s.write(c)
        else:
            break
    return int(s.getvalue())

cases = int(input())

for casen in range(cases):
    pnum = input()
    papers = map(int,input().split())
    
    current = 0
    caseheap = []
    print("Case #%d:" % (int(casen) + 1), end="")
    for n,cit in enumerate(papers):
        # n paper order, cit: citation 
        # analysis if should be changed
        if cit > current:
            heapq.heappush(caseheap, cit)
            print(caseheap)
        while len(caseheap) > current:
            current += 1
            while len(caseheap) > 0 and caseheap[0] == current:
                heapq.heappop(caseheap)
        print(" %d" % current, end="")
    print("")
