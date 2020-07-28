gotspl
======

Golang client for TSC printers (TSPL).

This library is being developed for communication and printing on TSPL/TSPL2 (by TSC) based printers.

Currently supported communication types:
* Ethernet
* RS232 (TODO)
* USB (TODO)

> **Disclaimer:** This is not official or supported by TSC Auto ID Technology Co., Ltd.
>
> I started developing it because I could not find an appropriate library for Golang. Contributions of any type are welcome. Please contact via issues to discuss further.


How to
======

Create printer connection

```go
# Initialize library 
gotspl.TSPLInitialize(gotspl.MEASUREMENT_SYSTEM_METRIC)

client := gotspl.NewEthernetTSPLClient("printer:9100")

err := client.Connect()
if err != nil {
    panic(err)
}

defer client.Disconnect()
```

Create label

```go
label := gotspl.NewTSPLLabel()
label = label.Cmd(gotspl.SizeCmd().
        LabelWidth(30).
        LabelLength(20)).
    Cmd(gotspl.GapCmd().
        LabelDistance(20).
        LabelOffsetDistance(0)).
    Cmd(gotspl.SpeedCmd().PrintSpeed(4)).
    Cmd(gotspl.ClsCmd()).
    Cmd(gotspl.DataMatrixCmd().
        XCoordinate(100).
        YCoordinate(100).
        Width(300).
        Height(300).
        Content("THIS IS DATAMATRIX TEST")).
    Cmd(gotspl.PrintCmd().NumberLabels(1).NumberCopies(1))
	
```

Send commands to printer

```go
err = client.SendCommandSequence(label)
if err != nil {
    panic(err)
}
```

TSPL Code will send to printer

```tspl
SIZE 30 mm,20 mm
GAP 20 mm,0
SPEED 4
CLS
DMATRIX 100,100,300,300, "THIS IS DATAMATRIX TEST"
PRINT 1,1
```

Based on documentation
======================

[Official documentation about TSPL programming](https://www.tscprinters.com/EN/DownloadFile/readpdf/support/4353/TSPL_TSPL2_Programming.pdf?file_type=0)

[Looked at this github repository (tspl2-driver) for library design patterns](https://github.com/fintrace/tspl2-driver)