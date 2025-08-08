False_ = 0
True_ = 1
N = 2

class TupleSpace:
    def __init__(self):
        self.space = {}

    def out(self, t):
        c = self.space.get(t, 0)
        if c < N:
            self.space[t] = c + 1
            print(f"out({t}) → {self.space[t]}")

    def _in(self, t):
        c = self.space.get(t, 0)
        if c > 0:
            self.space[t] = c - 1
            if self.space[t] == 0:
                del self.space[t]
            print(f"in({t}) → {self.space.get(t, 0)}")

    def inp(self, b, t):
        if b == True_:
            self._in(t)
        else:
            print(f"inp(False, {t}) → no change")

    def rd(self, t):
        c = self.space.get(t, 0)
        if c > 0:
            print(f"rd({t}) → exists")
        else:
            print(f"rd({t}) → not present")

    def rdp(self, b, t):
        if b == True_:
            self.rd(t)
        else:
            print(f"rdp(False, {t}) → no check")

    def __str__(self):
        return f"{self.space}"

ts = TupleSpace()
ts.out("any")
ts.rd("any")
ts.inp(True_, "any")
ts.rdp(True_, "any")
ts.inp(False_, "any")
ts.out("any")
ts.out("any")
ts.out("any")
print(ts)
