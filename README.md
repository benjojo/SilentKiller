SilentKiller
============

A tool that will kill the process you launch if it falls silent

The default timeout is 15 secs

you can change this with the -t <int> option

Usage:
```
ben@daring:~/SilentKiller$ ./SilentKiller sleep 10

ben@daring:~/SilentKiller$ ./SilentKiller sleep 16
2014/03/01 13:53:47 Timed out

ben@daring:~/SilentKiller$ ./SilentKiller -t 10 sleep 16
2014/03/01 13:54:10 Timed out

ben@daring:~/SilentKiller$ ./SilentKiller -t 2 sleep 15
2014/03/01 13:54:22 Timed out
```