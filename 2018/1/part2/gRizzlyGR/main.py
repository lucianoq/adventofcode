import sys
from io import StringIO


def reset(iterator):
    iterator.seek(0)


def repeat_frequency_list(frequency_list, freq, partial_results):
    for line in frequency_list:
        freq += int(line)

        if freq in partial_results:
            print(str(freq))
            return

        partial_results.append(freq)

    reset(frequency_list)
    repeat_frequency_list(frequency_list, freq, partial_results)


if __name__ == "__main__":
    frequency_list = StringIO(sys.stdin.read())
    repeat_frequency_list(frequency_list, 0, [0])
