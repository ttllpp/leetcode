# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```text
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解题思路1

```c
a + b = target
```

也可以看成是

```c
a = target - b
```

然后依次寻找差值a

## 解题思路2

```c
a + b = target
```

最暴力的方法，直接两个循环