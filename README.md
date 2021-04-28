# TXSolarPlanCalculator
A small tool to help decide which retail provider will give the best deal for solar buy-back.

# This is a tool I wrote for my own usage, I make no guarantee that it is correct or fit for any purpose.
Please, check my math and logic. Do not trust that I did things correctly. This was a quick-and-dirty calculator to help me decide which plan will be best for me. Your milage may vary.

# Usage
* Install Go for your operating system
* ``go get github.com/cloudkucooland/TXSolarPlanCalculator``
* Log into your SmartMeterTexas.com account
** Select Report Type "Energy Data 15 Min Interval"
** Select the start and end dates you wish to evaluate
** "Submit Update" button
** "Export My Report" button
* update the plans.json file with new data from https://www.texaspowerguide.com/solar-buyback-plans-texas/
* edit the interval.csv file to fill in any missing data (search for ,, and put in 0 or the approprite value (I have 6 null entries in my output files...)
* ``TXSolarPlanCalculator interval.csv``

# Example output
```
./TXSolarPlanCalculator example.csv | more
2020/11	Import: 603.64kwh	Export: 0.00kwh	Net: -603.64kwh
2020/12	Import: 781.40kwh	Export: 0.00kwh	Net: -781.40kwh
2021/1	Import: 668.91kwh	Export: 0.00kwh	Net: -668.91kwh
2021/2	Import: 446.95kwh	Export: 410.28kwh	Net: -36.67kwh
2021/3	Import: 345.02kwh	Export: 1181.01kwh	Net: 835.99kwh
2021/4	Import: 232.01kwh	Export: 1044.73kwh	Net: 812.72kwh

ATG Texas Solar Home
month	base	TDUbase	import	TDU	export	net
2020/11	$4.95	$3.42	$38.03	$23.48	$0.00	$-69.87
2020/12	$4.95	$3.42	$49.23	$30.39	$0.00	$-87.99
2021/1	$4.95	$3.42	$42.14	$26.01	$0.00	$-76.53
2021/2	$4.95	$3.42	$28.16	$17.38	$25.85	$-28.06
2021/3	$4.95	$3.42	$21.74	$13.42	$74.40	$30.88
2021/4	$4.95	$3.42	$14.62	$9.02	$65.82	$33.81

Chariot Rise Solar
month	base	TDUbase	import	TDU	export	net
2020/11	$4.95	$3.42	$37.43	$23.48	$0.00	$-69.27
2020/12	$4.95	$3.42	$48.45	$30.39	$0.00	$-87.21
2021/1	$4.95	$3.42	$41.47	$26.01	$0.00	$-75.86
2021/2	$4.95	$3.42	$27.71	$17.38	$25.44	$-28.03
2021/3	$4.95	$3.42	$21.39	$13.42	$73.22	$30.04
2021/4	$4.95	$3.42	$14.38	$9.02	$64.77	$33.00

GEX Solar Home
month	base	TDUbase	import	TDU	export	net
2020/11	$4.95	$3.42	$30.18	$23.48	$0.00	$-62.03
2020/12	$4.95	$3.42	$39.07	$30.39	$0.00	$-77.83
2021/1	$4.95	$3.42	$33.45	$26.01	$0.00	$-67.83
2021/2	$4.95	$3.42	$22.35	$17.38	$20.51	$-27.59
2021/3	$4.95	$3.42	$17.25	$13.42	$59.05	$20.01
2021/4	$4.95	$3.42	$11.60	$9.02	$52.24	$23.24

Green Mtn Renewable Rewards (12 mo)
month	base	TDUbase	import	TDU	export	net
2020/11	$0.00	$0.00	$89.94	$0.00	$0.00	$-89.94
2020/12	$0.00	$0.00	$116.43	$0.00	$0.00	$-116.43
2021/1	$0.00	$0.00	$99.67	$0.00	$0.00	$-99.67
2021/2	$0.00	$0.00	$66.60	$0.00	$61.13	$-5.46
2021/3	$0.00	$0.00	$51.41	$0.00	$175.97	$124.56
2021/4	$0.00	$0.00	$34.57	$0.00	$155.66	$121.10

Infuse Lean Green Infusion
month	base	TDUbase	import	TDU	export	net
2020/11	$4.95	$3.42	$39.84	$23.48	$0.00	$-71.69
2020/12	$4.95	$3.42	$51.57	$30.39	$0.00	$-90.33
2021/1	$4.95	$3.42	$44.15	$26.01	$0.00	$-78.53
2021/2	$4.95	$3.42	$29.50	$17.38	$27.08	$-28.17
2021/3	$4.95	$3.42	$22.77	$13.42	$77.95	$33.39
2021/4	$4.95	$3.42	$15.31	$9.02	$68.95	$36.25

Reliant Simple Solar Sell Back
month	base	TDUbase	import	TDU	export	net
2020/11	$0.00	$0.00	$77.87	$0.00	$0.00	$-77.87
2020/12	$0.00	$0.00	$100.80	$0.00	$0.00	$-100.80
2021/1	$0.00	$0.00	$86.29	$0.00	$0.00	$-86.29
2021/2	$0.00	$0.00	$57.66	$0.00	$52.93	$-4.73
2021/3	$0.00	$0.00	$44.51	$0.00	$152.35	$107.84
2021/4	$0.00	$0.00	$29.93	$0.00	$134.77	$104.84

TXU Renewal Buyback (24)
month	base	TDUbase	import	TDU	export	net
2020/11	$9.95	$0.00	$89.94	$0.00	$0.00	$-99.89
2020/12	$9.95	$0.00	$116.43	$0.00	$0.00	$-126.38
2021/1	$9.95	$0.00	$99.67	$0.00	$0.00	$-109.62
2021/2	$9.95	$0.00	$66.60	$0.00	$48.82	$-27.72
2021/3	$9.95	$0.00	$51.41	$0.00	$140.54	$79.18
2021/4	$9.95	$0.00	$34.57	$0.00	$124.32	$79.80

TXU Renewal Buyback (0)
month	base	TDUbase	import	TDU	export	net
2020/11	$9.95	$0.00	$83.91	$0.00	$0.00	$-93.86
2020/12	$9.95	$0.00	$108.62	$0.00	$0.00	$-118.57
2021/1	$9.95	$0.00	$92.98	$0.00	$0.00	$-102.93
2021/2	$9.95	$0.00	$62.13	$0.00	$36.52	$-35.56
2021/3	$9.95	$0.00	$47.96	$0.00	$105.11	$47.20
2021/4	$9.95	$0.00	$32.25	$0.00	$92.98	$50.78
```
