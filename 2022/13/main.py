import sys
import itertools


def chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i + n]


def compare(a, b):
  match a, b:
    case int(), int():
      return a - b
    case int(), list():
      return compare([a], b)
    case list(), int():
      return compare(a, [b])
    case list(), list():
      for i, item in enumerate(b):
        if i >= len(a):
          return -1
        res = compare(a[i], b[i])
        if res:
          return res
      return len(a) - len(b)


lines = itertools.chain.from_iterable(x.split() for x in sys.stdin.read().split("\n\n"))
packets = chunks([eval(l) for l in lines], 2)

total = 0
for i, (a, b) in enumerate(packets):
  if compare(a,b) <= 0:
    total += i + 1

print(total)
