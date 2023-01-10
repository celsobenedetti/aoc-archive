import re
import copy
from common import input  # type: ignore[attr-defined]

from stack import Stack


def run():
    inputStr = input.readFile("day5/input.txt")
    crateInput, instructions = readInput(inputStr)
    crateStacksPart1 = stackCrates(crateInput)
    crateStacksPart2 = copy.deepcopy(crateStacksPart1)

    for instruction in instructions:
        origin, to, amount = parseInstruction(instruction)
        interpretInstructionPart1(crateStacksPart1, origin, to, amount)
        interpretInstructionPart2(crateStacksPart2, origin, to, amount)

    p1 = "".join([stack.getTop() for stack in crateStacksPart1])
    p2 = "".join([stack.getTop() for stack in crateStacksPart2])
    return p1, p2


def readInput(inputStr: str) -> tuple[list[list[str]], list[str]]:
    crateStack, instructions = [], []
    readingCreates = True

    for line in inputStr.splitlines():
        if len(line) == 0:
            readingCreates = False
            continue

        if readingCreates:
            crateStack.append(splitCrates(line))
        else:
            instructions.append(line)

    return crateStack[:-1], instructions


def splitCrates(crates: str, crateLength=4) -> list[str]:
    def removeBrackets(string):
        return re.sub(r"\[|\]", "", string)

    def getCrate(i):
        return removeBrackets(crates[i : (i + crateLength)].strip())

    crateStacks = [getCrate(i) for i in range(0, len(crates) - 1, crateLength)]
    return crateStacks


def stackCrates(crateInput: list[list[str]]) -> list[Stack]:
    crateStacks = []

    numberOfStacks = len(crateInput[0])
    for i in range(0, numberOfStacks):
        crateStacks.append(Stack())

    crateInput.reverse()
    for step in crateInput:
        for i, crate in enumerate(step):
            if len(crate) > 0:
                crateStacks[i].push(crate)

    return crateStacks


def parseInstruction(instruction: str) -> tuple[int, int, int]:
    parts = [int(part) for part in list(filter(str.isdigit, instruction.split(" ")))]
    amount, origin, to = parts[0], parts[1] - 1, parts[2] - 1
    return origin, to, amount


def interpretInstructionPart1(crateStacks: list[Stack], origin, to, amount: int):
    for _ in range(0, amount):
        crateStacks[to].push(crateStacks[origin].pop())


def interpretInstructionPart2(crateStacks: list[Stack], origin, to, amount: int):
    crateStacks[to].pushMany(crateStacks[origin].popMany(amount))
