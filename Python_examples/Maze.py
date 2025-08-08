maze = {
    0: {'north': None,     'east': 1},
    1: {'east': 2,         'south': 4, 'west': 0},
    2: {'south': 5,        'west': 1},
    3: {'east': 4,         'south': 6},
    4: {'north': 1,        'west': 3},
    5: {'north': 2,        'south': 8},
    6: {'north': 3},
    7: {'east': 8},
    8: {'north': 5,        'west': 7}
}

def walk_maze(start, moves):
    pos = start
    for move in moves:
        if move in maze[pos]:
            next_pos = maze[pos][move]
            if next_pos is None:
                print(f"{pos} --{move}--> STOP")
                return
            print(f"{pos} --{move}--> {next_pos}")
            pos = next_pos
        else:
            print(f"{pos} --{move}--> invalid")
            return

walk_maze(0, ['east', 'east', 'south', 'south', 'north', 'north', 'west'])
