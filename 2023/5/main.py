import sys

seeds = []
seed_ranges = []
maps = []

sys.stdin = open('input.txt', 'r')
for line in sys.stdin:
    if line.startswith('seeds'):
        seeds = [int(n) for n in line.split(': ')[1].split(' ')]
        seed_ranges = [*zip(seeds[::2], seeds[1::2])]
        continue
    if 'map' in line:
        m = []
        for entry in sys.stdin:
            if not entry.strip():
                break
            m.append(tuple(int(n) for n in entry.split(' ')))
        maps.append(sorted(m, key=lambda x: x[1]))

for m in maps:
    for i, seed in enumerate(seeds):
        for dest, source, length in reversed(m):
            if source <= seed < source + length:
                seeds[i] = seed - source + dest
                break
        else:
            seeds[i] = seed
print("Part 1:", min(seeds))

min_seed = float('inf')
total = sum(end for _, end in seed_ranges)
count = 0
for start, end in seed_ranges:
    for seed in range(start, start+end):
        count += 1
        if count % 1000000 == 0:
            print(f"{count}/{total} ({round(count / total * 100, 2)}%)", end='\r')
        curr = seed
        for m in maps:
            for dest, source, length in reversed(m):
                if source <= curr < source + length:
                    curr = curr - source + dest
                    break

        if curr < min_seed:
            min_seed = curr
print("Part 2:", min_seed)
