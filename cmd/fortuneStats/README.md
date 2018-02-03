# FortuneStats

FortuneStats is a commandline program that calculates the frequencies and
distribution of byte values within all fortunes in the system's installed
fortune. It prints out Go var statements for frequencies and the distribution
data.

## Installation

To install, `go get` is sufficient on its own.

```shell
go get github.com/eseymour/goutils/cmd/fortuneStats
```

## Usage

Ensure fortune is installed on your system. FortuneStats takes no flags nor
parameters.

```shell
$ fortuneStats
var byteFreqs = [256]int{1, 1, 1, 1, 1, 1, 1, 60, 363, 31463, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 522560, 4225, 18451, 678, 157, 104, 169, 14823, 2192, 2214, 1006, 68, 35160, 23848, 42113, 623, 1653, 2423, 1290, 1012, 726, 889, 580, 633, 842, 1081, 5155, 1562, 80, 551, 146, 3781, 94, 9067, 5585, 5314, 4185, 5675, 5487, 7092, 4274, 11125, 4341, 1277, 4735, 4421, 7666, 5538, 4673, 1887, 4880, 7482, 11051, 3705, 4928, 5676, 729, 3412, 1725, 476, 325, 476, 74, 1295, 229, 176984, 81812, 54351, 57732, 243095, 78012, 94762, 105385, 116316, 17348, 16714, 87285, 41554, 172920, 143729, 46589, 31603, 188494, 115475, 157166, 93220, 64170, 32208, 10947, 68194, 19104, 33, 128, 12, 75, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var byteDist = [256]float64{3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 1.8884819909636137e-05, 0.00011425316045329863, 0.009902884813614695, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07,3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 0.164474191532991, 0.0013298060686368779, 0.005807396869211606, 0.00021339846497888835, 4.941527876354789e-05, 3.2733687843369305e-05, 5.319224274547512e-05, 0.004665494758675608, 0.0006899254206987069, 0.0006968498546655735, 0.0003166354804848992, 2.140279589758762e-05, 0.011066504467046776, 0.007506086420083376, 0.01325494034757511, 0.00019608738006172188, 0.0005202767885104755, 0.000762631977350806, 0.00040602362805717694, 0.0003185239624758628, 0.00022850632090659726, 0.0002798100816611088, 0.00018255325912648265, 0.00019923485004666124, 0.0002650169727318938, 0.0003402415053719444, 0.0016225207772362382, 0.0004916348116475274, 2.517975987951485e-05, 0.00017342559617015852, 4.59530617801146e-05, 0.0011900584013055705, 2.9586217858429946e-05, 0.002853811035344514, 0.0017578619865886303, 0.0016725655499967738, 0.0013172161886971206, 0.0017861892164530845, 0.0017270167807362246, 0.0022321857133189914, 0.0013452286715630807, 0.003501560358245034, 0.0013663167204621745, 0.0004019319170767558, 0.0014903270378687852, 0.0013914964803416893, 0.0024128504904545106, 0.0017430688776594154, 0.0014708127239621612, 0.0005939275861580565, 0.0015359653526504059, 0.002354937042731626, 0.0034782690803564823, 0.0011661376294200315, 0.0015510732085781147, 0.0017865039634515785, 0.00022945056190207906, 0.0010739167588613083, 0.0005429385724020389, 0.00014981957128311335, 0.00010229277451052908, 0.00014981957128311335, 2.3291277888551235e-05, 0.0004075973630496466, 7.207706265511125e-05, 0.0557051827814507, 0.02575008144078586, 0.017106814115143896, 0.01817097371705189, 0.07651342159888327, 0.024554042846508906, 0.029826055071282328, 0.033169612436283406, 0.03661011187682062, 0.005460230929872795, 0.00526068133282764, 0.02747269176354317, 0.013078996775417, 0.054426050979571344, 0.045238271346534874, 0.014663747912833966, 0.009946949393403848, 0.0593279207341159, 0.036345409651087214, 0.049467526765297885, 0.029340715199604676, 0.02019731489335585, 0.010137371327492678, 0.003445535392513113, 0.021463856815295446, 0.006012926659228146, 1.0386650950299875e-05, 4.028761580722376e-05, 3.776963981927227e-06, 2.360602488704517e-05, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07, 3.147469984939356e-07}
```