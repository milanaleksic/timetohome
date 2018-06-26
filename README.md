# timetohome

Compiled using **Go 1.10** and verified only on **Ubuntu 16.04LTS**.

Result: you get an icon in the systray with a time estimated to reach home (including traffic). 

![](screenshot.png)

There is also a link you can click on to get a more detailed map to see the reason for delay 
(you need to get the link manually though to choose the perspective etc).

## Usage:

```
$ timetohome 
Argument start is not set
  -apiKey string
        Developer API key
  -end string
        Coordinate of the ending point (like 49.238197,2.343180)
  -linkToVisit string
        Link to visit via 'Open MyDrive map' when clicked
  -start string
        Coordinate of the starting point (like 50.138197,2.273150)
```

## Installation

### Preconditions 

```bash
sudo apt-get install libgtk-3-dev libappindicator3-dev
```

### Installation (after preconditions have been installed)

```
go get github.com/milanaleksic/
```

