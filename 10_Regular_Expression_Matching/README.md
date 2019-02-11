'.' Matches any character.
'*' Matches zero or more of the preceding element.

cover the entire input

s could be empty and contains only lowercase letters a-z
p could be empty and contains only lowercase letters a-z, and character like '*' and '.'

Example1:
Input:
s = "aa"
p = "a"
Output: false

Example2:
Input:
s = "aa"
p = "a*"
Output: true

Example3:
Input:
s = "ab":
p = ".*"
Output: true

Example4:
Input:
s = "aab"
p = "c*a*b"
Output: true

Example5:
Input:
s = "mississippi"
p = "mis*is*p"
Output: false

