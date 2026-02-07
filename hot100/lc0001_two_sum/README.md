# LC0001 Two Sum

## 题意
给定数组 nums 和 target，返回两个下标，使得 nums[i] + nums[j] = target。

## 思路
- 哈希表记录 value -> index
- 遍历时查找 target-nums[i] 是否已出现

## 复杂度
- 时间：O(n)
- 空间：O(n)

## 易错点
- 同一个元素不能用两次
- 返回下标顺序要求（按题目定义）