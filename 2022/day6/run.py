from common import input as fs  # type: ignore[attr-defined]
from day6.queue import FixedLenQueue


def getMarkerIndex(queue: FixedLenQueue, input: str):
    for i, char in enumerate(input):
        queue.enqueue(char)
        if queue.full():
            return i + 1

    return -1


def run():
    input = fs.readFile("day6/input.txt")

    part1Queue = FixedLenQueue(4)
    part2Queue = FixedLenQueue(14)

    part1 = getMarkerIndex(part1Queue, input)
    part2 = getMarkerIndex(part2Queue, input)

    return part1, part2
