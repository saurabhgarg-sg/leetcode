"""
    Solution for replicating each 0 occurence in the list.
    Has to be done in-place, without modifying array length.
"""


class Solution:
    """
        Leetcode solution class.
    """

    def __init__(self):
        self.output = None

    def duplicateZeros(self, arr: list[int]) -> None:
        """
            leetcode: Do not return anything, modify arr in-place instead.
        """
        current_index = 0
        arr_size = len(arr)
        while current_index < arr_size - 1:
            if arr[current_index] == 0:
                for mover in reversed(range(current_index + 2, arr_size)):
                    arr[mover] = arr[mover - 1]
                arr[current_index + 1] = 0
                current_index += 2
            else:
                current_index += 1

        # Remove this during submission.
        self.output = arr


if __name__ == "__main__":
    sol = Solution()
    test_cases = (
        [1, 0, 2, 3, 0, 4, 5, 0],
        [51, 3, 0, 10, 19, 34],
        [11, 0, 32, 54, 0, 61],
        [12, 13, 14, 15, 16, 17],
        [21, 33, 45, 55, 23, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 31, 0, 17, 0, 101],
        []
    )
    expected_output = (
        [1, 0, 0, 2, 3, 0, 0, 4],
        [51, 3, 0, 0, 10, 19],
        [11, 0, 0, 32, 54, 0],
        [12, 13, 14, 15, 16, 17],
        [21, 33, 45, 55, 23, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 31, 0, 0, 17],
        []
    )
    for index, input_list in enumerate(test_cases):
        temp = list(input_list)
        sol.duplicateZeros(arr=input_list)
        result = "PASS" if sol.output == expected_output[index] else "FAIL"
        print(f"{result}: for {temp} -> expected: {expected_output[index]}, got: {sol.output}")
