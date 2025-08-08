import random

listeners = ['a', 'b', 'c', 'd']
patterns = {'a': 'pat1', 'b': 'pat2', 'c': 'pat1', 'd': 'pat2'}

state = {l: 'IDLE' for l in listeners}
registered_pattern = {}

def register(listener, pattern):
    if state[listener] == 'IDLE':
        registered_pattern[listener] = pattern
        state[listener] = 'LISTENING'
        print(f"{listener} registered for {pattern}")

def deregister(listener):
    state[listener] = 'STOP'
    print(f"{listener} deregistered")

def announce(pattern):
    print(f"Announce: {pattern}")
    for l in listeners:
        if state[l] == 'LISTENING' and registered_pattern.get(l) == pattern:
            outcome = random.choice(['hit', 'miss', 'timeout'])
            print(f"  {l} reacts: {outcome}")
            if outcome == 'hit':
                deregister(l)
            elif outcome == 'miss':
                pass
            elif outcome == 'timeout':
                pass

def simulate():
    for l in listeners:
        register(l, patterns[l])
    announce('pat1')
    announce('pat2')
    announce('pat1')

simulate()
