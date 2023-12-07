import sys
from collections import defaultdict
from functools import cmp_to_key

C = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T",  "Q", "K", "A"]

hands = []
for line in sys.stdin:
    cards, bid = line.strip().split(" ")
    hands.append((cards, int(bid)))

def rank1(hand):
    d = defaultdict(int)
    for c in hand:
        d[c] += 1
    values = sorted(d.values(), reverse=True)
    if values == [5]:
        return 6
    if values == [4, 1]:
        return 5
    if values == [3, 2]:
        return 4
    if values == [3, 1, 1]:
        return 3
    if values == [2, 2, 1]:
        return 2
    if values == [2, 1, 1, 1]:
        return 1
    return 0

def rank2(hand):
    d = defaultdict(int)
    for c in hand:
        d[c] += 1
    values = sorted(d.values(), reverse=True)
    if values == [5]:
        return 6
    if values == [4, 1]:
        if d["J"] == 4 or d["J"] == 1:
            return 6
        else:
            return 5
    if values == [3, 2]:
        if d["J"] == 3 or d["J"] == 2:
            return 6
        else:
            return 4
    if values == [3, 1, 1]:
        if d["J"] == 3 or d["J"] == 1:
            return 5
        else:
            return 3
    if values == [2, 2, 1]:
        if d["J"] == 2:
            return 5
        elif d["J"] == 1:
            return 4
        else:
            return 2
    if values == [2, 1, 1, 1]:
        if d["J"] == 2 or d["J"] == 1:
            return 3
        else:
            return 1
    if d["J"] == 1:
        return 1
    else:
        return 0

def compare(h1, h2):
    r1 = rank2(h1[0])
    r2 = rank2(h2[0])
    if r1 > r2:
        return 1
    elif r1 < r2:
        return -1
    else:
        return compare_high_card(h1[0], h2[0])

def compare_high_card(c1, c2):
    for c1, c2 in zip(c1, c2):
        c1 = C.index(c1)
        c2 = C.index(c2)
        if c1 > c2:
            return 1
        elif c1 < c2:
            return -1
    return 0

total = 0
for i, t in enumerate(sorted(hands, key=cmp_to_key(compare))):
  total += (i+1) * t[1]
print(total)
