import math

def Ivec(p):
    result = [0] * len(p)
    for i in range(len(p)):
        result[i] = math.log2(1/p[i])
    return result

def H(p):
    I = Ivec(p)
    result = 0
    for i in range(len(p)):
        result += p[i] * I[i]
    return result

def shannon(p):
    syms = [0] * len(p)
    for i in range(len(p)):
        syms[i] = math.ceil(math.log2(1/p[i]))
    average = 0
    for i in range(len(syms)):
        average += syms[i] * p[i]
    eta = H(p)/average
    return (syms, average, eta)


p = [0.20, 0.19, 0.18, 0.17, 0.15, 0.10, 0.01]
l = [2,2,3,3,3,4,4]
res = shannon(p)
res = 0
for i in range(len(p)):
    res += p[i] * l[i]

res = 0
for i in range(8):
    res += 0.1 * (i + 1) * 0.9**i
res += 8 * 0.9**8
print(res)