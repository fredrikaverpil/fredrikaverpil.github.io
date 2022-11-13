# git bisect
[[git]]
1.  Have a failing test
2.  `git bisect start`
3.  `git bisect bad`
4.  Run test again

    1.  If OK, run  `git bisect good`
    2.  if test is still failing, run `git bisect bad`
