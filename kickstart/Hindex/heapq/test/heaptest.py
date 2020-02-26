import heapq
import time
import random
h = []
cur = 0

start = time.time()
for i in range(100000000):
    ai = random.randint(0,100000000)
    if ai > cur:
        heapq.heappush(h,ai)
    while len(h)>0:
        cur+=1
        while len(h)>0 and h[0]==cur:
            heapq.heappop(h)

print(time.time()-start)
