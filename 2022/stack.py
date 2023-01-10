from typing import TypeVar, Generic

T = TypeVar("T")


class Stack(Generic[T]):
    stack: list[T]

    def __init__(self) -> None:
        self.stack = []

    def push(self, data: T):
        self.stack.append(data)

    def pop(self) -> T:
        return self.stack.pop()

    def getTop(self) -> T:
        return self.stack[-1]

    def print(self):
        print(self.stack)
