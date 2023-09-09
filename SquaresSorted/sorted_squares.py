"""
    Solution for sorting squares of a list in decreasing order.
"""


class Solution:
    """
        Leetcode solution class.
    """

    def sortedSquares(self, nums: list[int]) -> list[int]:
        """
            Function to square and sort the list elements.
        """
        squares = [num * num for num in nums]
        squares.sort()
        return squares


if __name__ == "__main__":
    sol = Solution()
    test_cases = ([-4, -1, 0, 3, 10], [-7, -3, 2, 3, 11],)
    expected_output = ([0, 1, 9, 16, 100], [4, 9, 9, 49, 121])
    for index, input_list in enumerate(test_cases):
        output = sol.sortedSquares(nums=input_list)
        result = "PASS" if output == expected_output[index] else "FAIL"
        print(f"{result}: for {input_list} -> expected: {expected_output[index]}, got: {output}")
