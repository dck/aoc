import sys
from collections import defaultdict

def main():
    total1 = 0
    total2 = 0
    D: list[str] = []
    gears = defaultdict(dict)
    for i, line in enumerate(sys.stdin):
        line = line.strip()
        if not line:
            continue
        D.append(line)
    for i in range(len(D)):
        j = 0
        k = 0
        while j < len(D[i]):
          k = j
          while k < len(D[i]) and D[i][k].isdigit():
            k += 1
          if k != j and is_legit(D, j, k-1, i):
            total1 += int(D[i][j:k])

          if k != j:
            gear_i, gear_j = find_gear(D, j, k-1, i)
            if gear_i is not None:
              prev_n, prev_ratio = gears[gear_i].get(gear_j, (0, 1))
              gears[gear_i][gear_j] = (prev_n+1, prev_ratio * int(D[i][j:k]))
          j = k+1

    print(total1)

    for _, d in gears.items():
      for _, v in d.items():
        if v[0] == 2:
          total2 += v[1]
    print(total2)

def is_legit(D: list[str], start: int, end: int, row: int) -> bool:
  for i in range(row-1, row+2):
    for j in range(start-1, end+2):
      if i < 0 or i >= len(D) or j < 0 or j >= len(D[i]):
        continue
      if i == row and start <= j <= end:
        continue
      if D[i][j] != ".":
        return True

  return False

def find_gear(D: list[str], start: int, end: int, row: int) -> bool:
  for i in range(row-1, row+2):
    for j in range(start-1, end+2):
      if i < 0 or i >= len(D) or j < 0 or j >= len(D[i]):
        continue
      if i == row and start <= j <= end:
        continue
      if D[i][j] == "*":
        return (i, j)

  return (None, None)



if __name__ == "__main__":
    main()
