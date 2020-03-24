### RMQ 区间最大(最小)值

- 求[l,r]区间内的最大值
- 令 f(i,j) = [i,i+2^j-1]区间内的最大值
  - f(i,0) = ai
  - f(i,j) = max(f(i,j-1),f(i+2^(j-1),j-1))
- [l,r]区间内最大值 = max(f(l,s),f(r-2^s+1,r)) 令s = log2(r-l+1)

- 优化
  - 建议用数组保存log
    - log[i] = log[i/2]+1