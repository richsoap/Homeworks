import scipy.io as scio
import numpy as np

class Builder:
    record = []
    z = 0
    def __init__(self, z):
        self.record = np.zeros((z + 1, z), dtype=int)
        for i in range(z):
            self.record[i+1][i] = 1
        self.z = z
    def getLine(self, v, l):
        if v == 0:
            return self.record[0]
        if v + l <= self.z:
            return self.record[v + l]
        else:
            return self.record[(v + l) % self.z]
    def extra(self, matrix):
        result = np.zeros((self.z * len(matrix), self.z * len(matrix[0])), dtype=int)
        for i in range(len(matrix)):
            for j in range(self.z):
                for k in range(len(matrix[0])):
                    line = self.getLine(matrix[i][k], j)
                    for index in range(self.z):
                        result[i*self.z + j][k*self.z + index] = line[index]
        offset = int(len(result[0])/2) - self.z
        for i in range(self.z):
            result[0][offset + i] = 0
        return result

def main():
    dataFile = './Matrix(2016,1008)Block56.mat'
    data = scio.loadmat(dataFile)
    data = data['H_block']
    bu = Builder(56)
    result = bu.extra(data)
    with open("./H_block.txt", "w") as f:
        f.write("{} {} {}\n".format(len(result), len(result[0]), 56)) 
        for line in result:
            for val in line:
                f.write(str(val))
                f.write(" ")
            f.write("\n")

def test():
    bu = Builder(4)
    data = [[0], [1], [2], [3], [4]]
    result = bu.extra(data)
    print(result)

main()