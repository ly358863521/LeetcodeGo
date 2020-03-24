#%%

import random

MaxLevel = 16
p = 0.25

class node:
    def __init__(self,value,level):
        self.value = value
        self.level = level
        self.forward = [None for _ in range(level)]

def randlevel():
    for level in range(1,MaxLevel):
        if random.random()>p:
            return level
class Skiplist:

    def __init__(self):
        self.header = node(-1,MaxLevel)
        self.level = 1

    def search(self, target: int) -> bool:
        cur = self.header
        for i in range(self.level-1,-1,-1):
            while cur.forward[i]!=None:
                if cur.forward[i].value == target:
                    return True
                elif cur.forward[i].value>target:
                    break
                else:
                    cur = cur.forward[i]
        return False

    def add(self, num: int) -> None:
        cur = self.header
        update = [None for _ in range(MaxLevel)]
        for i in range(self.level-1,-1,-1):
            if cur.forward[i]==None or cur.forward[i].value>num:
                update[i] = cur
            else:
                while cur.forward[i]!=None and cur.forward[i].value<num:
                    cur = cur.forward[i]
                update[i] = cur
        level = randlevel()
        if level>self.level:
            for i in range(self.level,level):
                update[i] = self.header
            self.level = level
        n = node(num,level)
        for i in range(level):
            n.forward[i] = update[i].forward[i]
            update[i].forward[i] = n

    def erase(self, num: int) -> bool:
        cur = self.header
        flag = False
        for i in range(self.level-1,-1,-1):
            while cur.forward[i]!= None:
                if cur.forward[i].value == num:
                    flag = True
                    cur.forward[i] = cur.forward[i].forward[i]
                    break
                elif cur.forward[i].value>num:
                    break
                else:
                    cur = cur.forward[i]
        if flag:
            return True
        return False


# %%
m = Skiplist()
m.add(1)

# %%
m.search(1)

# %%
