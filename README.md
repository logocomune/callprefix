# Callsign information

This golang library give general information of ham callsign by its prefix

## Example
```golang
package main

import (
	"fmt"
	"github.com/logocomune/callprefix"
)

func main() {

	info, ok := callprefix.Info("iu5pmp")
	if !ok {
		fmt.Println("No info found on callsing")
	}

	fmt.Printf("%+v\n", info)

}

```
Response:
```
{
    CountryName:Italy 
    CQZone:15 
    ITUZone:28 
    Continent:EU 
    Lat:42.82 
    Lng:-12.58 
    TimeOffset:-1 
    Prefix:I 
    IsFullCallsign:false 
    Flag:🇮🇹 
    CountryCode:IT
}
```
