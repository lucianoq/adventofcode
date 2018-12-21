#!/usr/local/bin/python3


# def dump(sb, elf1, elf2):
#     for i, x in enumerate(sb):
#         if i == elf1:
#             print(f"({x})", end='')
#             continue
#         if i == elf2:
#             print(f"[{x}]", end='')
#             continue
#         print(f" {x} ", end='')
#     print()


def next_10_after(input):
    elf1 = 0
    elf2 = 1
    sb = [3, 7]
    for i in range(input + 10):
        sb += list(map(int, list(str(sb[elf1] + sb[elf2]))))
        elf1 = (elf1 + 1 + sb[elf1]) % len(sb)
        elf2 = (elf2 + 1 + sb[elf2]) % len(sb)
        # dump(sb, elf1, elf2)
    return "".join(list(map(str, sb[input:input + 10])))


if __name__ == '__main__':
    input = 286051
    print(next_10_after(input))
