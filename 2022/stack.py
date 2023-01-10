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

    def pushMany(self, data: list[T]):
        self.stack.extend(data)

    def popMany(self, amount: int):
        result = self.stack[-amount:]
        self.stack = self.stack[:-amount]
        return result

    def getTop(self) -> T:
        return self.stack[-1]

    def print(self):
        print(self.stack)
