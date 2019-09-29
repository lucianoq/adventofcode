#!/usr/bin/env python3


def first_appears(sequence):
    elf_1 = 0
    elf_2 = 1
    scoreboard = "37"
    input_len = len(sequence)

    while True:
        recipe1 = int(scoreboard[elf_1])
        recipe2 = int(scoreboard[elf_2])
        recipe_sum = recipe1 + recipe2
        if recipe_sum > 9:
            scoreboard += "1"
            if len(scoreboard) >= input_len and scoreboard[-input_len:] == sequence:
                return len(scoreboard) - input_len
        scoreboard += str(recipe_sum % 10)
        if len(scoreboard) >= input_len and scoreboard[-input_len:] == sequence:
            return len(scoreboard) - input_len

        elf_1 = (elf_1 + 1 + recipe1) % len(scoreboard)
        elf_2 = (elf_2 + 1 + recipe2) % len(scoreboard)


if __name__ == '__main__':
    assignment = "286051"
    print(first_appears(assignment))
