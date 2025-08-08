listeners = ['a', 'b', 'c', 'd']
patterns = ['pat1', 'pat2']

listener_state = {l: {'state': 'IDLE', 'pattern': None} for l in listeners}

event_queue = []

def register(listener, pattern):
    if listener_state[listener]['state'] == 'IDLE':
        listener_state[listener]['state'] = 'MATCH'
        listener_state[listener]['pattern'] = pattern
        print(f"{listener} registered for {pattern}")

def deregister(listener):
    listener_state[listener]['state'] = 'IDLE'
    listener_state[listener]['pattern'] = None
    print(f"{listener} deregistered")

def announce(pattern):
    print(f"Announce: {pattern}")
    event_queue.append(pattern)
    for l in listeners:
        state = listener_state[l]
        if state['state'] == 'MATCH':
            if state['pattern'] == pattern:
                print(f"{l} received event: {pattern}")
                if l in ['b', 'c']: 
                    print(f"{l} auto-deregistered after event")
                    deregister(l)

def simulate():
    register('a', 'pat1')
    register('b', 'pat2')
    register('c', 'pat1')
    register('d', 'pat2')

    announce('pat1')
    announce('pat2')

    register('b', 'pat1')
    announce('pat1')

simulate()
