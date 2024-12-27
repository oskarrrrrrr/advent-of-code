import sys
import typing as t

Map = list[list[str]]
Pos = t.Tuple[int, int]

UP = (-1, 0)
DOWN = (1, 0)
LEFT = (0, -1)
RIGHT = (0, 1)
# invalid type, should use enum or something
Dir = t.Literal[UP, DOWN, LEFT, RIGHT]

def turn_right(dir: Dir) -> Dir:
    return { UP: RIGHT, RIGHT: DOWN, DOWN: LEFT, LEFT: UP}[dir]

def get_starting_pos(_map: Map) -> Pos:
    for line_nr, line in enumerate(_map):
        for ch_nr, ch in enumerate(line):
            if ch == "^":
                return line_nr, ch_nr
    assert(False)

def at(_map: Map, pos: Pos) -> str:
    return _map[pos[0]][pos[1]]

def set_at(_map: Map, pos: Pos, char: str) -> None:
    _map[pos[0]][pos[1]] = char

def valid_pos(_map: Map, pos: Pos) -> bool:
    return (
        0 <= pos[0] < len(_map[0]) and
        0 <= pos[1] < len(_map)
    )

def move(pos: Pos, dir: Dir) -> Pos:
    return (pos[0] + dir[0], pos[1] + dir[1])

def part1(_map: Map) -> None:
    pos = get_starting_pos(_map)
    dir = UP
    visited: set[Pos] = { pos }
    while valid_pos(_map, pos):
        next_pos = move(pos, dir)
        if not valid_pos(_map, next_pos):
            break
        if at(_map, next_pos) == "#":
            dir = turn_right(dir)
        else:
            visited.add(next_pos)
            pos = next_pos
    print(f"part1: {len(visited)}")

def is_in_loop(_map: Map, pos: Pos, dir: Dir) -> bool:
    visited: set[(Pos, Dir)] = { (pos, dir) }
    while valid_pos(_map, pos):
        next_pos = move(pos, dir)
        if (next_pos, dir) in visited:
            return True
        if not valid_pos(_map, next_pos):
            return False
        if at(_map, next_pos) == "#":
            dir = turn_right(dir)
        else:
            visited.add((next_pos, dir))
            pos = next_pos

def part2(_map: Map) -> None:
    obstacles: set[Pos] = set()
    starting_pos = get_starting_pos(_map)
    pos = starting_pos
    dir = UP
    while valid_pos(_map, pos):
        next_pos = move(pos, dir)
        if not valid_pos(_map, next_pos):
            break
        if at(_map, next_pos) == "#":
            dir = turn_right(dir)
        else:
            curr = at(_map, next_pos)
            set_at(_map, next_pos, "#")
            if is_in_loop(_map, starting_pos, UP):
                obstacles.add(next_pos)
            set_at(_map, next_pos, curr)
            pos = next_pos
    print(f"part2: {len(obstacles)}")

def main() -> None:
    _map = [list(line.strip()) for line in sys.stdin]
    part1(_map)
    part2(_map)

main()
