# GIBDD Sender

GIBDD Sender is a utility to warn police about parking law violations.

## Work Principles

The utility works with https://mvd.ru API which is also used by official MVD Android application. It doesn't require any secrets or credentials.

Used things:

* sender details: name, phone, email
* receiver details: region and MVD unit responsible for the district where the violation happened
* message
* photos 

The utility:

* searches for files with `jpg` and `jpeg` extensions
* resizes them to 2000 px width, so that more than 50 files can be sent in a single request
* creates a zip archive with these files
* sends it to MVD

When it succeeds a confirmation email will be received from MVD with message details.

## Common Usage

* take photos of the parking law violation
* place them to a directory with name of pattern: `date address`, where date format is `YYYY-MM-DD HH-MM` like `2020-01-15 18-41`
* create `config.json` file with proper configuration
* run `gibdd-sender "files/directory/path" "config/file/path"`    

## Configuration 

```json
{
  "lastName": "",
  "name": "",
  "middleName": "",
  "email": "",
  "phone": "",
  "regionId": "",
  "subunit": "",
  "receiver": "",
  "archiveType": "",
  "messageTemplate": ""
}
```

Supported values for `receiver`: 
* `mvd`

Supported values for `archiveType`:
* `zip`
* `pdf`

To get list of regions run `GET https://mvd.ru/api/address/subject`, choose yours and use its `code`

To get list of subunits run `GET https://mvd.ru/api/request/get_sub_by_region?regions=%region%`, choose yours and use its `key` 

`messageTemplate` is processed by [Go Template](https://golang.org/pkg/text/template/) package with the following context:
```go
package main 

import (
	"time"
)

type TimeAddress struct {
	Time    time.Time
	Address string
}
``` 
This structure is parsed from images directory name.