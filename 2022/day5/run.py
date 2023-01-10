import re
from common import input  # type: ignore[attr-defined]

from stack import Stack


def splitCrates(crates: str, crateLength=4) -> list[str]:
    def removeBrackets(string):
        return re.sub(r"\[|\]", "", string)

    def getCrate(i):
        return removeBrackets(crates[i : (i + crateLength)].strip())

    crateStacks = [getCrate(i) for i in range(0, len(crates) - 1, crateLength)]
    return crateStacks


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


def interpretInstruction(crateStacks: list[Stack], instruction: str):
    parts = [part for part in instruction.split(" ")]
    parts = list(filter(str.isdigit, parts))
    parts = [int(part) for part in parts]

    amount, origin, to = parts[0], parts[1] - 1, parts[2] - 1

    for _ in range(0, amount):
        crateStacks[to].push(crateStacks[origin].pop())


def run():
    inputStr = input.readFile("day5/input.txt")
    crateInput, instructions = readInput(inputStr)
    crateStacks = stackCrates(crateInput)

    for instruction in instructions:
        interpretInstruction(crateStacks, instruction)

    p1 = "".join([stack.getTop() for stack in crateStacks])
    return p1, -1
