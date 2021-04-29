package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

type plan struct {
	Name      string
	Base      float64 // base charge
	ExportKwh float64 // price per kwh
	ImportKwh float64 // rebate per kwh
	TDUBase   float64 // base delivery cost
	TDUKwh    float64 // delivery cost per kwh
	MaxKwhNet int     // maximum net surplus the provider will purchase, -1 for no limit
	Rollover  bool    // does the net surplus roll over month-to-month
	Cashout   bool    // cash out on contract end
}

type data struct {
	Date   time.Time
	Import float64
	Export float64
}

func main() {
	var datafile, planfile string

	flag.Parse()
	args := flag.Args()
	argc := len(args)
	if argc != 1 {
		panic("usage: txsolarplans [datafile]")
	}
	datafile = args[0]

	// let this be a cli flag (-p)
	planfile = "plans.json"
	plans, err := loadPlans(planfile)
	if err != nil {
		panic(err)
	}

	data, err := loadData(datafile)
	if err != nil {
		panic(err)
	}
	y := dataToMonthly(data)

	for _, v := range plans {
		fmt.Printf("\n%s\n", v.Name)
		simulation(v, y)
	}
}

func simulation(p plan, data map[int]map[int]data) {
	sortYears := make([]int, 1)
	for year := range data {
		sortYears = append(sortYears, year)
	}
	sort.Slice(sortYears, func(i, j int) bool {
		return sortYears[i] < sortYears[j]
	})

	fmt.Printf("month\tbase\tTDUbase\timport\tTDU\texport\tnet\n")
	for _, year := range sortYears {
		yeardata, ok := data[year]
		if !ok {
			continue
		}
		for month := 0; month <= 12; month++ {
			monthdata, ok := yeardata[month]
			if !ok {
				continue
			}
			base := p.Base + p.TDUBase
			imp := p.ImportKwh * monthdata.Import
			tdu := p.TDUKwh * monthdata.Import

			limitedexport := monthdata.Export
			// if a plan has a net export ceiling
			if p.MaxKwhNet != -1 && (monthdata.Export-monthdata.Import) > float64(p.MaxKwhNet) {
				limitedexport = float64(p.MaxKwhNet)
			}
			// TXU: no net positive export -- no, only if the past year was net total
			/* if p.MaxKwhNet == 0 && monthdata.Export > monthdata.Import {
				limitedexport = monthdata.Import
			} */
			exp := limitedexport * p.ExportKwh
			net := exp - base - imp - tdu
			fmt.Printf("%d/%d\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\n", year, month, p.Base/100, p.TDUBase/100, imp/100, tdu/100, exp/100, net/100)
		}
	}
}

func loadData(file string) (map[time.Time]data, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	d := make(map[time.Time]data)
	for _, v := range lines {
		if v[0] == "ESIID" {
			continue
		}
		date := fmt.Sprintf("%s %s", v[1], v[3])
		parsed, err := time.Parse("01/02/2006 15:04", date)
		if err != nil {
			panic(err)
		}
		val, err := strconv.ParseFloat(v[5], 64)
		if err != nil {
			panic(err)
		}

		f, ok := d[parsed]
		if ok {
			if v[7] == "Consumption" {
				f.Import = val
			} else {
				f.Export = val
			}
			d[parsed] = f
		} else {
			var e data
			e.Date = parsed
			if v[7] == "Consumption" {
				e.Import = val
			} else {
				e.Export = val
			}
			d[parsed] = e
		}
	}

	return d, nil
}

func dataToMonthly(d map[time.Time]data) map[int]map[int]data {
	sortYears := make([]int, 1)

	years := make(map[int]map[int]data)
	for k, v := range d {
		year := k.Year()
		y, ok := years[year]
		if !ok {
			y = make(map[int]data)
			years[year] = y
			sortYears = append(sortYears, year)
		}

		month := int(k.Month())
		m, ok := y[month]
		if !ok {
			m = data{}
		}
		m.Import += v.Import
		m.Export += v.Export
		y[month] = m
	}

	sort.Slice(sortYears, func(i, j int) bool {
		return sortYears[i] < sortYears[j]
	})

	for _, year := range sortYears {
		yeardata, ok := years[year]
		if !ok {
			continue
		}
		for month := 0; month <= 12; month++ {
			monthdata, ok := yeardata[month]
			if !ok {
				continue
			}
			fmt.Printf("%d/%d\tImport: %4.2fkwh\tExport: %4.2fkwh\tNet: %4.2fkwh\n", year, month, monthdata.Import, monthdata.Export, monthdata.Export-monthdata.Import)
		}
	}
	return years
}

func loadPlans(file string) ([]plan, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var plans []plan
	err = json.Unmarshal(f, &plans)
	if err != nil {
		panic(err)
	}

	return plans, nil
}
