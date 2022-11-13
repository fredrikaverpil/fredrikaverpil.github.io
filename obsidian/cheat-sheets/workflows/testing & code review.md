# Test-driven development
## Rules
1.  You may not write any production code when you do _**not**_ have a failing test
2.  You may not write any test code when you have a failing test
3.  You may only write the least possible amount of produiction code to make the test pass

## Red-green-refactor
1.  Write failing test
2.  Write production code (making test pass)
3.  Refactor
4. (Rinse and repeat)

# Behavior driven development
Steps:
1. Given
2. When
3. Then
There's a Python package called [behave](https://github.com/behave/behave).

# Code review

## Assert using previous production code
All tests added in PR, must fail:
[[git]]
```bash
# 1. Restore the production code from master
git restore -s master <path to production file>

# 2. Re-run tests
# 3. Verify all new tests fail
# 4. Restore files back to current PR branch's state
git restore <path to production file>
```

