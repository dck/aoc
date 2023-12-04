import sys
from collections import defaultdict
from functools import reduce

limits = {
    'red': 12,
    'green': 13,
    'blue': 14
}

def main():
    total1 = 0
    total2 = 0
    failed = 0
    for g, line in enumerate(sys.stdin):
        try:
            line = line.strip()
            if not line:
                continue
            _, line = line.split(': ', 1)
            rounds = line.split('; ')
            d = defaultdict(int)
            for round in rounds:
                balls = round.split(', ')
                for ball in balls:
                    n, color = ball.split(' ')
                    n = int(n)
                    if n > d[color]:
                        d[color] = n
                    # if n > limits[color]:
                    #     raise ValueError('Invalid ball number')
            total1 += (g+1)
            power = reduce(lambda x, y: x * y, d.values(), 1)
            total2 += power
            print(line, power)
        except ValueError:
            continue
    print(total1)
    print(total2)

if __name__ == "__main__":
    main()
