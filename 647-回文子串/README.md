### 回文子串

- 回文子串若左右两边加上相同的元素，则仍为回文子串

- ```go
  func countSubstrings(s string) int {
      count:=0
      for i:=0;i<len(s);i++{
          count += countPalindrome(s,i,i)
          count += countPalindrome(s,i,i+1)
      }
      return count
  }
  func countPalindrome(s string,left,right int)int{
      count :=0
      for left>=0 && right<len(s) &&s[left]==s[right]{
          left--
          right++
          count++
      }
      return count
  }
  ```