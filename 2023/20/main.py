import os
import sys
from functools import reduce
from dataclasses import dataclass
from collections import defaultdict
from enum import Enum

class Pulse(Enum):
    HIGH = 1
    LOW = 2


class FlipFlop:
  def __init__(self, bus, name, children) -> None:
    self.bus = bus
    self.name = name
    self.children = children
    self._state = False

  def on_high(self, _source):
    pass

  def on_low(self, _source):
    if self._state:
      self._state = False
      pulse_to_send = Pulse.LOW
    else:
      self._state = True
      pulse_to_send = Pulse.HIGH
    for child in self.children:
      self.bus.send(self.name, pulse_to_send, child)

class Conjunction:
  def __init__(self, bus, name, children) -> None:
    self.bus = bus
    self.name = name
    self.children = children
    self.sources_types = {}

  def declare_source(self, source):
    self.sources_types[source] = Pulse.LOW

  def on_high(self, source):
    self.sources_types[source] = Pulse.HIGH
    self.on_signal()

  def on_low(self, source):
    self.sources_types[source] = Pulse.LOW
    self.on_signal()

  def on_signal(self):
    if any(t == Pulse.LOW for t in self.sources_types.values()):
      signal_to_send = Pulse.HIGH
    else:
      signal_to_send = Pulse.LOW
    for child in self.children:
      self.bus.send(self.name, signal_to_send, child)

class Broadcast:
  def __init__(self, bus, name, children) -> None:
    self.bus = bus
    self.name = name
    self.children = children

  def on_high(self, _source):
    for child in self.children:
      self.bus.send(self.name, Pulse.HIGH, child)

  def on_low(self, _source):
    for child in self.children:
      self.bus.send(self.name, Pulse.LOW, child)

class Bus:
  def __init__(self) -> None:
    self.messages = []
    self.modules = {}
    self.stats = {}

  def send(self, source, pulse, target):
    self.messages.append((source, pulse, target))
    self.stats[pulse] = self.stats.get(pulse, 0) + 1

  def register(self, name, module):
    self.modules[name] = module

class Button:
  def __init__(self, bus, child) -> None:
    self.bus = bus
    self.child = child

  def click(self):
    self.bus.send('button', Pulse.LOW, self.child)


bus = Bus()
button = None

for line in sys.stdin:
  line = line.strip()
  if not line:
    continue

  name, children = line.split(' -> ')
  children = children.split(', ')

  if name == 'broadcaster':
    broadcast = Broadcast(bus, name, children)
    bus.register(name, broadcast)
    button = Button(bus, broadcast.name)
    continue

  if name.startswith('%'):
    flipflop = FlipFlop(bus, name[1:], children)
    bus.register(flipflop.name, flipflop)
    continue

  if name.startswith('&'):
    conj = Conjunction(bus, name[1:], children)
    bus.register(conj.name, conj)
    continue

for source, module in bus.modules.items():
  for child in module.children:
    if isinstance(bus.modules.get(child), Conjunction):
      bus.modules[child].declare_source(source)

for i in range(1000):
  button.click()

  while bus.messages:
    source, pulse, target = bus.messages.pop(0)
    destination = bus.modules.get(target)

    if pulse == Pulse.HIGH:
      print(f'{source} -high-> {target}')
      if destination is not None:
        destination.on_high(source)
    else:
      print(f'{source} -low-> {target}')
      if destination is not None:
        destination.on_low(source)

print("Part 1:", reduce(lambda x, y: x * y, bus.stats.values()))

total = 0
while True:
  button.click()
  total += 1

  while bus.messages:
    source, pulse, target = bus.messages.pop(0)
    destination = bus.modules.get(target)

    if pulse == Pulse.HIGH:
      if destination is not None:
        destination.on_high(source)
    else:
      if destination is not None:
        destination.on_low(source)

    if pulse == Pulse.HIGH and source in ['fm', 'dk', 'fg', 'pq']:
      print(f'{source} -high-> {target} ({total})')

print("Part 2:", total)
