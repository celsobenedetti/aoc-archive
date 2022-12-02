from functools import reduce

def readFile(fileName: str):
    f = open(fileName, "r")
    return reduce(lambda string, line: string+line, f.readlines())
