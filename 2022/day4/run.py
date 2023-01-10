from common import input  # type: ignore[attr-defined]
from functools import reduce
from typing import Callable


debugInput = """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"""

ElfSections = list[str]
ElfPair = tuple[ElfSections, ElfSections]


def getPairs(inputString: str) -> list[ElfPair]:
    elfs = [line.split(",") for line in inputString.splitlines()]
    pairs = [(elf[0].split("-"), elf[1].split("-")) for elf in elfs]
    return pairs


def pairCompletelyOverlaps(pair: ElfPair) -> bool:
    elf1, elf2 = pair[0], pair[1]
    e1p1, e1p2 = int(elf1[0]), int(elf1[1])
    e2p1, e2p2 = int(elf2[0]), int(elf2[1])
    oneContainsTwo = e1p1 <= e2p1 and e1p2 >= e2p2
    twoContainsOne = e2p1 <= e1p1 and e2p2 >= e1p1
    return oneContainsTwo or twoContainsOne


def pairOverlaps(pair: ElfPair) -> bool:
    elf1, elf2 = pair[0], pair[1]
    bordersOverlap = elf1[0] == elf2[0] or elf1[1] == elf2[1]
    oneBeforeTwo = elf1[0] <= elf2[0] and elf1[1] >= elf2[0]
    twoBeforeOne = elf2[0] <= elf1[0] and elf2[1] >= elf1[0]
    return (
        bordersOverlap or oneBeforeTwo or twoBeforeOne or pairCompletelyOverlaps(pair)
    )


def countOverlaps(pairs: list[ElfPair], pairOverlaps: Callable[[ElfPair], bool]) -> int:
    return reduce(
        lambda overlapCount, pair: overlapCount + 1
        if pairOverlaps(pair)
        else overlapCount,
        pairs,
        0,
    )


def run():
    inputStr = input.readFile("day4/input.txt")
    pairs = getPairs(inputStr)
    p1 = countOverlaps(pairs, pairCompletelyOverlaps)
    p2 = countOverlaps(pairs, pairOverlaps)

    for pair in pairs:
        if pairOverlaps(pair):
            print(pair)
    return p1, p2
