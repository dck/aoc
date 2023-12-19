import os
import sys
from functools import reduce
from dataclasses import dataclass
from collections import defaultdict
from typing import Optional


@dataclass
class Part:
  x: int
  m: int
  a: int
  s: int

  def sum(self) -> int:
    return self.x + self.m + self.a + self.s

@dataclass
class PartRange:
  x: list[int]
  m: list[int]
  a: list[int]
  s: list[int]

  def apply_condition(self, clause: str):
    field, sign, val = clause[0], clause[1], int(clause[2:])
    left, right = getattr(self, field)

    if sign == '<' and val > left:
      new_val = [left, min(val-1, right)]
    elif sign == '>' and val < right:
      new_val = [max(val+1, left), right]
    else:
      new_val = []

    newPart = PartRange(self.x, self.m, self.a, self.s)
    setattr(newPart, field, new_val)
    return newPart

  def apply_reversed_condition(self, clause: str):
    field, sign, val = clause[0], clause[1], int(clause[2:])
    if sign == '>':
      sign = '<'
      val = val + 1
    else:
      sign = '>'
      val = val - 1
    return self.apply_condition(field + sign + str(val))

  def is_empty(self) -> bool:
    return not self.x or not self.m or not self.a or not self.s

  def cost(self) -> int:
    return (self.x[1] - self.x[0] + 1) * (self.m[1] - self.m[0] + 1) * (self.a[1] - self.a[0] + 1) * (self.s[1] - self.s[0] + 1)

@dataclass
class Condition:
  target: str
  clause: Optional[str] = None


matrix = defaultdict(list)
parts = []

for line in sys.stdin:
  if not line.split():
    continue

  if line.startswith('{'):
    part = Part(0,0,0,0)
    for field in line[1:-2].split(','):
      name, val = field.split('=')
      setattr(part, name.strip(), int(val))
    parts.append(part)
  else:
    name, conds = line.split('{')
    name = name.strip()
    conds = conds[:-2].split(',')
    for c in conds:
      f = c.split(':')
      if len(f) > 1:
        matrix[name].append(Condition(f[1].strip(), f[0].strip()))
      else:
        matrix[name].append(Condition(f[0].strip()))


def check_workflow(name: str, part: Part) -> bool:
  if name == 'A':
    return True
  if name == 'R':
    return False

  for cond in matrix[name]:
    if cond.clause:
      if eval(f'part.{cond.clause}'):
        return check_workflow(cond.target, part)
    else:
      return check_workflow(cond.target, part)
  return True

def check_workflow_range(name: str, part: PartRange):
  if name == 'A':
    return [part]
  if name == 'R':
    return []

  ranges = []

  final_cond = part
  for i, cond in enumerate(matrix[name]):
    if cond.clause:
        if i>0:
          next_cond = final_cond.apply_condition(cond.clause)
        else:
          next_cond = part.apply_condition(cond.clause)

        final_cond = final_cond.apply_reversed_condition(cond.clause)
        ranges.extend(check_workflow_range(cond.target, next_cond))
    else:
      ranges.extend(check_workflow_range(cond.target, final_cond))
  return ranges


total = 0
for part in parts:
  if check_workflow('in', part):
    total += part.sum()
print("Part 1:", total)


total = 0
r = PartRange([1, 4000], [1, 4000], [1, 4000], [1, 4000])
applied_ranges = check_workflow_range('in', r)

for r in sorted(applied_ranges, key=lambda r: r.cost()):
  print(r)
  total += r.cost()

print("Part 2:", total)
