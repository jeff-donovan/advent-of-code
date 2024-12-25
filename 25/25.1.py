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

if __name__ == '__main__':
    with open('25/day_25_input.txt', 'r') as f:
        contents = f.read()

    locks, keys = parse_input(contents)
    print(len(locks))
    print(len(keys))
