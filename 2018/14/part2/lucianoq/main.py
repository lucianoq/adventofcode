#!/usr/local/bin/python3


def first_appears(input):
    elf1 = 0
    elf2 = 1
    sb = "37"

    l = len(input)
    while True:
        sum = int(sb[elf1]) + int(sb[elf2])

        if sum > 9:
            sb += "1" + str(sum % 10)
        else:
            sb += str(sum)

        print(sb[:2] == "37")
        elf1 = (elf1 + 1 + int(sb[elf1])) % len(sb)
        elf2 = (elf2 + 1 + int(sb[elf2])) % len(sb)
        # dump(sb, elf1, elf2)
        if len(sb) >= l and sb[-l:] == input:
            # dump(sb, elf1, elf2)
            return len(sb) - l


if __name__ == '__main__':
    input = "286051"
    print(first_appears(input))
