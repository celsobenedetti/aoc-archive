import day6.run as day6
from day6.queue import FixedLenQueue

TestInput1 = "bvwbjplbgvbhsrlpgdmjqwftvncz"
TestInput2 = "nppdvjthqldpwncqszvftbrmjlhg"
TestInput3 = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
TestInput4 = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"


def test_getMarkerIndex():
    assert day6.getMarkerIndex(FixedLenQueue(4), "abcd") == 4
    assert day6.getMarkerIndex(FixedLenQueue(4), "abcad") == 5
    assert day6.getMarkerIndex(FixedLenQueue(4), TestInput1) == 5
    assert day6.getMarkerIndex(FixedLenQueue(4), TestInput2) == 6
    assert day6.getMarkerIndex(FixedLenQueue(4), TestInput3) == 10
    assert day6.getMarkerIndex(FixedLenQueue(4), TestInput4) == 11


def test_queue():
    queue = FixedLenQueue(4)
    queue.enqueue("a")
    queue.enqueue("b")
    queue.enqueue("c")
    queue.enqueue("d")
    assert queue.queue == ["a", "b", "c", "d"]
    assert queue.indexOf("a") == 0
    assert queue.indexOf("b") == 1
    assert queue.indexOf("c") == 2
    assert queue.indexOf("d") == 3
    queue = FixedLenQueue(4)
    queue.enqueue("e")
    assert queue.queue == ["e"]
    queue.enqueue("a")
    assert queue.queue == ["e", "a"]
    queue.enqueue("a")
    assert queue.queue == ["a"]
