MAX = 9
N = 4

def generator():
    for x in range(2, MAX + 1):
        yield x
    yield 'eos'

def filter_stage(input_gen, p=None):
    for x in input_gen:
        if x == 'eos':
            yield 'eos'
            break
        if p is None:
            p = x
            print(p)
            yield x
        elif x % p != 0:
            yield x

def run_pipeline():
    g = generator()
    for _ in range(N):
        g = filter_stage(g)
    for _ in g:
        pass

run_pipeline()
