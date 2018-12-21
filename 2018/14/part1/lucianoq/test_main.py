#!/usr/local/bin/python3

import unittest

from .main import next_10_after


class TestInput(unittest.TestCase):

    def test_9(self):
        self.assertEqual(next_10_after(9), "5158916779")

    def test_5(self):
        self.assertEqual(next_10_after(5), "0124515891")

    def test_18(self):
        self.assertEqual(next_10_after(18), "9251071085")

    def test_2018(self):
        self.assertEqual(next_10_after(2018), "5941429882")


if __name__ == '__main__':
    unittest.main()
