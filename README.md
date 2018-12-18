# ptad
Papertrail Archives Downloader

Support concurrent downloading multiple archives `-n`, local timezone `-l` (defalt UTC) and only hourly format (not daily).

## Install
Use `go get`,
```
go get -u github.com/giraffate/ptad
```
or use `brew`,
```
brew tap giraffate/ptad
brew install ptad
```

## Usage
After setting Papertrail API token, run the following command (local timezone `JST`),
```
$ ptad -n 3 -d -l 2018-11-29-16
2018/12/10 10:37:58 [DEBUG] [Run as debug mode]
2018/12/10 10:37:58 [DEBUG] num: [3]
2018/12/10 10:38:01 [DEBUG] Completed: [2018-11-29-09]
2018/12/10 10:38:01 [DEBUG] Completed: [2018-11-29-08]
2018/12/10 10:38:01 [DEBUG] Completed: [2018-11-29-07]
```

### Papertrail API token
Set it via environmental variable,
```
$ export PAPERTRAIL_API_TOKEN=xxxxxxxx
