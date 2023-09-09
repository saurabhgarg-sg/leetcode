"""
    Solution for finding numbers in the list that have even digits count.
"""


class Solution:
    """
        Leetcode solution class.
    """

    def findNumbers(self, nums: list[int]) -> int:
        """
            Function to count the numbers with even digits.
        """

        # List comprehension is slower for this problem as we don't need the numbers.
        # return len([num for num in nums if len(str(num)) % 2 == 0])

        even_digit_numbers = 0
        for num in nums:
            if len(str(num)) % 2 == 0:
                even_digit_numbers += 1

        return even_digit_numbers


if __name__ == "__main__":
    sol = Solution()
    test_cases = (
        [12, 345, 2, 6, 7896],
        [555, 901, 482, 1771],
        [12, 112345, 1000],
        [101, 220, 67220, 1231210, 110, 99990],
        [0, 1, 0, 1, 0, 1],
        []
    )
    expected_output = (2, 1, 3, 0, 0, 0)
    for index, input_list in enumerate(test_cases):
        output = sol.findNumbers(nums=input_list)
        result = "PASS" if output == expected_output[index] else "FAIL"
        print(f"{result}: for {input_list} -> expected: {expected_output[index]}, got: {output}")
