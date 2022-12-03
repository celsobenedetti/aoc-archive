from common.input import readFile

testInput = '''A Y
B X
C Z
'''

playerMoves = {
    'X': 1,
    'Y': 2,
    'Z': 3
}

win = {
    'A': 2,
    'B': 3,
    'C': 1,
}

loss = {
    'A': 3,
    'B': 1,
    'C': 2,
}

draw = {
    'A': 1,
    'B': 2,
    'C': 3,
}

def matchHandsPart1(opponent: str, player: str):
    if opponent == 'A':
        if player == 'X':
            return 3
        elif player == 'Y':
            return 6
        return 0

    if opponent == 'B':
        if player == 'X':
            return 0
        elif player == 'Y':
            return 3
        return 6
    
    if player == 'X':
        return 6
    elif player == 'Y':
        return 0
    return 3

def matchHandsPart2(opponent: str, player: str):
    if player == 'X':
        return 0 + loss[opponent] 

    if player == 'Y':
        return 3 + draw[opponent]
    
    return 6 + win[opponent]

def run():
    part1, part2 = 0, 0
    input = readFile("day2/input.txt")

    for line in input.split('\n'):
        if len(line) == 0:
            continue

        moves = line.split(' ')
        opponent, player = moves[0], moves[1]
        
        part1 += matchHandsPart1(opponent, player) + playerMoves[player]
        part2 += matchHandsPart2(opponent, player)
    
    return part1, part2


    



    


