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
* ``TXSolarPlanCalculator interval.csv``

# Example output
```
scot@covert:/home/scot/TXSolarPlanCalculator % ./TXSolarPlanCalculator example.csv 
2021/2	Import: 370.44kwh	Export: 384.88kwh	Net: 14.44kwh
2021/3	Import: 23.50kwh	Export: 82.49kwh	Net: 58.99kwh

ATG Texas Solar Home
month	base	TDUbase	import	TDU	export	net
2021/2	$4.95	$3.42	$23.34	$14.41	$24.25	$-21.87
2021/3	$4.95	$3.42	$1.48	$0.91	$5.20	$-5.57

Chariot Rise Solar
month	base	TDUbase	import	TDU	export	net
2021/2	$4.95	$3.42	$22.97	$14.41	$23.86	$-21.88
2021/3	$4.95	$3.42	$1.46	$0.91	$5.11	$-5.63

GEX Solar Home
month	base	TDUbase	import	TDU	export	net
2021/2	$4.95	$3.42	$18.52	$14.41	$19.24	$-22.05
2021/3	$4.95	$3.42	$1.18	$0.91	$4.12	$-6.33

Green Mtn Renewable Rewards (12 mo)
month	base	TDUbase	import	TDU	export	net
2021/2	$0.00	$0.00	$55.20	$0.00	$57.35	$2.15
2021/3	$0.00	$0.00	$3.50	$0.00	$3.73	$0.22

Infuse Lean Green Infusion
month	base	TDUbase	import	TDU	export	net
2021/2	$4.95	$3.42	$24.45	$14.41	$25.40	$-21.82
2021/3	$4.95	$3.42	$1.55	$0.91	$5.44	$-5.39

Reliant Simple Solar Sell Back
month	base	TDUbase	import	TDU	export	net
2021/2	$0.00	$0.00	$47.79	$0.00	$49.65	$1.86
2021/3	$0.00	$0.00	$3.03	$0.00	$10.64	$7.61

TXU Renewal Buyback (24)
month	base	TDUbase	import	TDU	export	net
2021/2	$9.95	$0.00	$55.20	$0.00	$44.08	$-21.06
2021/3	$9.95	$0.00	$3.50	$0.00	$2.80	$-10.66
```
