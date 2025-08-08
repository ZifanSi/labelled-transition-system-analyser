from queue import Queue

MAX = 9
N = 4

class Generator:
    def __init__(self, max_val, output):
        self.max_val = max_val
        self.output = output

    def run(self):
        for x in range(2, self.max_val + 1):
            self.output.put(x)
        self.output.put('eos')

class Filter:
    def __init__(self, input_q, output_q):
        self.input = input_q
        self.output = output_q
        self.prime = None

    def run(self):
        while True:
            x = self.input.get()
            if x == 'eos':
                self.output.put('eos')
                break
            if self.prime is None:
                self.prime = x
                print(self.prime)
                self.output.put(x)
            elif x % self.prime != 0:
                self.output.put(x)

def run_pipeline(n, max_val):
    queues = [Queue() for _ in range(n + 1)]
    gen = Generator(max_val, queues[0])
    filters = [Filter(queues[i], queues[i + 1]) for i in range(n)]
    gen.run()
    for f in filters:
        f.run()

run_pipeline(N, MAX)
