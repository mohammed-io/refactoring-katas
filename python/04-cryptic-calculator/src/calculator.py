class Calculator:
    def __init__(self):
        pass

    def normalize(self, a, b, c):
        d = []
        e = float("inf")
        f = float("-inf")

        for g in a:
            if g < e:
                e = g
            if g > f:
                f = g

        h = f - e
        i = c - b

        for g in a:
            try:
                k = b + ((g - e) / h) * i
            except ZeroDivisionError:
                k = float("nan")
            d.append(round(k, 2))

        return d
