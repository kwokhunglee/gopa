# Gopa #
[狗爬], [aims to be] A High Performance Distributed  Spider Written in GO .

CHANGES

v0.7
features:
1.add stats api to expose the task info, http://localhost:8001/stats
2.add websocket and simple ui to interact with Gopa, http://localhost:8001/ui/
3.add task api to accept seed
4.dynamic change the seelog config via api, [GET/POST] http://localhost:8001/setting/seelog/
5.follow 301/302 redirect, and continue fetch
6.add boltdb status page, http://localhost:8001/ui/boltdb
7.add pipeline framework to create crawler
8.add command to dynamic change logging level and add seed
8.export metrics to statsD
9.support daemon mode in linux and darwin
improve:
1.add update_ui setup to Makefile in order to build static ui

v0.6
breaking change:
1.remove bloom, use leveldb to store urls
feature:
1.crawling speed control
2.cookie supported
3.brief logging format
bugfix:
1.shutdown nil exception
2.wrong relative link in parse phrase


v0.5
feature:
1.ruled fetch
2.fetch/parse offset can be persisted and reloadable
3.http console

v0.4
improve:
1.refactor storage interface,data path are now configable
2.disable pprof by default
3.use local storage instead of kafka,kafka will be removed later
5.check local file's exists first before fetch the remote page
bugfix:
resolve memory leak caused by sbloom filter
feature:
1.download by url template
2.list page download

v0.3
1.adding golang pprof,http://localhost:6060/debug/pprof/
    go tool pprof http://localhost:6060/debug/pprof/heap
    go tool pprof http://localhost:6060/debug/pprof/profile
    go tool pprof http://localhost:6060/debug/pprof/block

2.integrate with kafka to make task controllable and recoverable
3.parameters configable
4.goroutine can be controlled now


v0.2
1.bloom-filter persistence
2.building script works

v0.1
1.just up and run.


