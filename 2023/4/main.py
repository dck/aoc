import sys
from collections import defaultdict

def main():
    total = 0
    cards = defaultdict(int)
    for i, line in enumerate(sys.stdin):
        cards[i] += 1
        line = line.strip()

        if not line:
            continue
        first, rest = line.split('|', 1)
        _, first = first.split(':')
        winning = {int(x.strip()) for x in first.strip().split()}
        numbers = [int(x.strip()) for x in rest.strip().split()]
        count = sum(1 for x in numbers if x in winning)
        if count:
          total += 2 ** (count - 1)
        for j in range(count):
            cards[i+1+j] += cards[i]
    print(total)
    print(sum(cards.values()))

if __name__ == "__main__":
    main()
