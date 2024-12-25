def parse_input(contents):
    locks = []
    keys = []

    for lock_or_key in contents.split('\n\n'):
        if lock_or_key:
            schematic = []
            for line in lock_or_key.split('\n'):
                if line:
                    new_row = []
                    for char in line:
                        new_row.append(char)
                    schematic.append(new_row)
            if is_lock(schematic):
                locks.append(schematic)
            elif is_key(schematic):
                keys.append(schematic)
            else:
                print('neither! ', schematic)
    return locks, keys

def is_lock(schematic):
    return all([schematic[0][j] == '#' for j in range(len(schematic[0]))]) and all([schematic[-1][j] == '.' for j in range(len(schematic[0]))])

def is_key(schematic):
    return all([schematic[0][j] == '.' for j in range(len(schematic[0]))]) and all([schematic[-1][j] == '#' for j in range(len(schematic[0]))])

def get_lock_pin_heights(lock):
    pin_heights = []
    for j in range(len(lock[0])):
        height = 0
        for i in range(len(lock)):
            if lock[i][j] == '#':
                height = i
        pin_heights.append(height)
    return pin_heights

def get_key_pin_heights(key):
    pin_heights = []
    for j in range(len(key[0])):
        height = 0
        for i in range(len(key)):
            if key[i][j] == '.':
                height = i
        pin_heights.append(len(key) - 2 - height)
    return pin_heights

if __name__ == '__main__':
    with open('25/day_25_test.txt', 'r') as f:
        contents = f.read()

    locks, keys = parse_input(contents)

    lock_heights = [get_lock_pin_heights(lock) for lock in locks]
    print(lock_heights)

    key_heights = [get_key_pin_heights(key) for key in keys]
    print(key_heights)
