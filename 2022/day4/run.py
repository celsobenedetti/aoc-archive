from common import input  # type: ignore[attr-defined]
from functools import reduce

debugInput = """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"""

ElfSections = list[str]
Pair = tuple[ElfSections, ElfSections]


def getPairs(inputString: str) -> list[Pair]:
    elfs = [line.split(",") for line in inputString.splitlines()]
    pairs = [(elf[0].split("-"), elf[1].split("-")) for elf in elfs]
    return pairs


def pairsCompletelyOverlap(pair: Pair) -> bool:
    elf1, elf2 = pair[0], pair[1]
    oneContainsTwo = int(elf1[0]) <= int(elf2[0]) and int(elf1[1]) >= int(elf2[1])
    twoContainsOne = int(elf2[0]) <= int(elf1[0]) and int(elf2[1]) >= int(elf1[1])
    return oneContainsTwo or twoContainsOne


def countCompleteOverlaps(pairs: list[Pair]) -> int:
    return reduce(
        lambda overlapCount, pair: overlapCount + 1
        if pairsCompletelyOverlap(pair)
        else overlapCount,
        pairs,
        0,
    )


def run():
    inputStr = input.readFile("day4/input.txt")
    pairs = getPairs(inputStr)
    p1 = countCompleteOverlaps(pairs)
    return p1, 0


run()
