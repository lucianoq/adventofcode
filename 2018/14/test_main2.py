#!/usr/local/bin/python3

import unittest

from .main import first_appears


class TestInput(unittest.TestCase):

    def test_9(self):
        self.assertEqual(first_appears("51589"), 9)

    def test_5(self):
        self.assertEqual(first_appears("01245"), 5)

    def test_18(self):
        self.assertEqual(first_appears("92510"), 18)

    def test_2018(self):
        self.assertEqual(first_appears("59414"), 2018)


if __name__ == '__main__':
    unittest.main()
