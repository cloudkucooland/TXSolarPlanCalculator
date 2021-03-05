package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
	MaxKwhNet int   // maximum net surplus the provider will purchase, -1 for no limit
	Rollover  bool    // does the net surplus roll over month-to-month
	Cashout   bool    // cash out on contract end
}

type data struct {
    Date time.Time
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

    for _, v := range *plans {
        fmt.Printf("\n%s\n", v.Name)
        simulation(v, y)
    }
}

func simulation(p plan, data *map[int]*map[int]data) {
    fmt.Printf("month\tbase\tTDUbase\timport\tTDU\texport\tnet\n")
    for year, yeardata := range *data {
        for month, monthdata := range *yeardata {
            base := p.Base + p.TDUBase
            imp := p.ImportKwh * monthdata.Import
            tdu := p.TDUKwh * monthdata.Import

            limitedexport := monthdata.Export
            // Green Mtn, only buyback up to 25kwh over import
            if p.MaxKwhNet != -1 && (monthdata.Export - monthdata.Import) > float64(p.MaxKwhNet) {
                limitedexport = float64(p.MaxKwhNet)
            }
            // TXU: no net positive export
            if p.MaxKwhNet == 0 && monthdata.Export > monthdata.Import {
                limitedexport = monthdata.Import
            }
            exp := limitedexport * p.ExportKwh
            net := exp - base - imp - tdu
            fmt.Printf("%d/%d\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\t$%4.2f\n", year, month, p.Base / 100, p.TDUBase / 100, imp / 100, tdu / 100, exp / 100, net / 100)
        }
    }
}

func loadData(file string) (*map[time.Time]data, error) {
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

	return &d, nil
}

func dataToMonthly(d *map[time.Time]data) *map[int]*map[int]data {
    years := make(map[int]*map[int]data)
    for k, v := range *d {
       year := k.Year()
       y, ok := years[year]
       if !ok {
           tmp := make(map[int]data)
           y = &tmp
           years[year] = y
       }

       month := int(k.Month())
       tt := *y
       m, ok := tt[month]
       if !ok {
           m = data{}
       }
       m.Import += v.Import
       m.Export += v.Export
       tt[month] = m
    }

    for ky, vy := range years {
        for km, vm := range *vy {
            fmt.Printf("%d/%d\tImport: %4.2fkwh\tExport: %4.2fkwh\tNet: %4.2fkwh\n", ky, km, vm.Import, vm.Export, vm.Export - vm.Import)
        }
    }
    return &years
}

func loadPlans(file string) (*[]plan, error) {
    f, err := os.ReadFile(file)
    if err != nil {
        panic(err)
    }

    var plans []plan
    err = json.Unmarshal(f, &plans)
    if err != nil {
        panic(err)
    }

	return &plans, nil
}
