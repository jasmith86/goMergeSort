#!/usr/bin/python3

from typing import List


def merge(left: List[int], right: List[int]) -> (List[int], int):
    """
    Merge two arrays of ints.
    :return: Merged (sorted) array of ints and number of inversions
    """
    results = []
    inversions = 0
    i = 0
    j = 0
    count = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            results.append(left[i])
            i += 1
        else:
            results.append(right[j])
            j += 1
            inversions += len(left) - i
    results += left[i:] + right[j:]
    return results, inversions


def merge_sort(array: List[int]) -> (List[int], int):
    """
    Use mergesort algorithm to sort array.

    :return: Return sorted array of ints and number of inversions.
    """
    n = len(array)
    if n < 2:
        return array, 0
    mid = n // 2
    left = array[:mid]
    right = array[mid:]
    left, linv = merge_sort(left)
    right, rinv = merge_sort(right)
    results, inv = merge(left, right)
    return results, inv + linv + rinv


if __name__ == '__main__':
    input = {
        0:  [1, 1, 1, 2, 2],
        4:  [2, 1, 3, 1, 2],
        15: [5, 4, 3, 2, 1, 0],
    }

    for ans in input.keys():
        arr = input[ans]
        n = len(arr)

        arr = merge_sort(arr)
        print(arr)

