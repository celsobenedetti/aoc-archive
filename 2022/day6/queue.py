class FixedLenQueue:
    maxLen: int
    queue: list[str]

    def __init__(self, length: int) -> None:
        self.maxLen = length
        self.queue = list[str]()

    def indexOf(self, char: str) -> int:
        try:
            return self.queue.index(char)
        except ValueError:
            return -1

    def enqueue(self, char: str):
        existingIndex = self.indexOf(char)

        if existingIndex != -1:
            self.queue = self.queue[existingIndex + 1 :]

        self.queue.append(char)

    def size(self) -> int:
        return len(self.queue)

    def full(self) -> bool:
        return self.size() == self.maxLen
