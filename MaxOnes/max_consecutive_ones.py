"""
    Solution for counting max consecutive ones in the list.
"""


class Solution:
    """
        Leetcode solution class.
    """

    def findMaxConsecutiveOnes(self, nums: list[int]) -> int:
        """
            Function to count the biggest group of ones.
        """
        counter = 0
        max_count = 0
        for num in nums:
            if num == 0:
                counter = 0
            else:
                counter += 1
                if counter > max_count:
                    max_count = counter

        return max_count


if __name__ == "__main__":
    sol = Solution()
    test_cases = (
        [1, 1, 0, 1, 1, 1],
        [1, 0, 1, 1, 0, 1],
        [1, 1, 1, 1, 1, 1],
        [0, 0, 0, 0, 0, 0],
        [0, 1, 0, 1, 0, 1],
        []
    )
    expected_output = (3, 2, 6, 0, 1, 0)
    for index, input_list in enumerate(test_cases):
        output = sol.findMaxConsecutiveOnes(nums=input_list)
        result = "PASS" if output == expected_output[index] else "FAIL"
        print(f"{result}: for {input_list} -> expected: {expected_output[index]}, got: {output}")
