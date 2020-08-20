#Boyer Moore 算法

-   匹配方式是从 pattern 的末尾开始
-   坏字符移动位数的计算规则就是：
    -   移动位数 = pattern 匹配到坏字符的位置 – 坏字符在 pattern 中的位置

-   好后缀位移规则
-   移动位数 = 好后缀在 pattern 中出现的位置 – 好后缀在 pattern 中的重复出现的位置
