# page_checker

## How To Use.

download.
```
git clone git@github.com:hachi-n/page_checker.git
```

build.
```
make build
```

Image Check.
```
./pkg/page_checker img --json #{JSON_FILE_PATH}
```

Output File
```
./pkg/page_checker img --json #{JSON_FILE_PATH} --output #{DEST/FILENAME}
```

## JSON check
* Error Pattern
```
$ cat result.json | jq '.[] | select(.Judge == false)'
```
