 
def coinChangeFast(coins, amount):
    dp = [0] + [-1] * amount
    for x in range(amount):
      if dp[x] < 0:
          continue
      for c in coins:
          if x + c > amount:
              continue
          if dp[x + c] < 0 or dp[x + c] > dp[x] + 1:
              dp[x + c] = dp[x] + 1
              print dp
    return dp[amount]

def coinChange(coins, amount):
    dp = [amount + 1 for _ in range(amount + 1)]
    dp[0] = 0
    for i in range(amount + 1):
        for coin in coins:
            if i - coin >= 0:
                dp[i] = min(dp[i], dp[i - coin] + 1)

    return dp[amount] if dp[amount] <= amount else -1


print coinChangeFast([1, 2, 5],11)
# print coinChange([2],3)