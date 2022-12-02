from functools import reduce
from common.input import readFile

test_input = '''1000
2000
3000

4000

5000
6000

7000
8000
9000

10000'''


def parseInput(input: str):
    elfs = []
    for i, line in enumerate(input.split("\n")):
        if len(line) == 0:
            elfs.append([])
            continue

        calories = int(line)
        if len(elfs) == 0:
            elfs.append([calories])
        else:
            elfs[-1].append(calories)
    return elfs


def sortCalories(elfs):
    elfCalories = [reduce(lambda acc, c: acc+c, elf) for elf in elfs]
    return sorted(elfCalories)


def run():
    elfs = parseInput(readFile("day1/input.txt"))
    sortedCalories = sortCalories(elfs)
    part1 = sortedCalories[-1]
    part2 = sortedCalories[-1]+sortedCalories[-2]+sortedCalories[-3]
    return (part1, part2)
