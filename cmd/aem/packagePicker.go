package main

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
	"regexp"
	"strconv"
)

func newPackagePicker() packagePicker {
	return packagePicker{
		sliceUtil: new(sliceUtil),
	}
}

type packagePicker struct {
	sliceUtil *sliceUtil
}

func (b *packagePicker) addPackagesToTable(t table.Writer, packages []packageDescription) {
	for i, pkg := range packages {
		t.AppendRow(table.Row{i + 1, pkg.Name, pkg.Version})
	}
}

func (b *packagePicker) appendSelected(list []int64, selected int64) []int64 {
	if !b.sliceUtil.inSliceInt64(list, selected) {
		list = append(list, selected)
	}
	return list
}

func (b *packagePicker) getSelectedPackages(selected []int64, bundles []packageDescription) []packageDescription {
	selectedPackages := make([]packageDescription, 0)
	for _, bundle := range selected {
		selectedPackages = append(selectedPackages, bundles[bundle-1])
	}
	return selectedPackages
}

func (b *packagePicker) picker(instance aemInstanceConfig) []packageDescription {
	http := new(httpRequests)
	pkgs := http.getListForInstance(instance)
	pageSize := 20
	selected := make([]int64, 0)
	selectedPackages := make([]packageDescription, 0)
	writer := new(tableWriter)

	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Package", "Version"})
	t.SetPageSize(pageSize)
	t.SetOutputMirror(writer)
	b.addPackagesToTable(t, pkgs)
	t.Render()
	tables := writer.getTables()

	for i := 0; i < len(tables); i++ {
		fmt.Print(tables[i])
	choose:
		fmt.Printf("Selected %d\n", selected)
		fmt.Print("d: done selecting, q: quit, c: continue, package id ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[0 : len(input)-1]
		switch input {
		case "c":
			continue
		case "q":
			return make([]packageDescription, 0)
		case "d":
			return b.getSelectedPackages(selected, pkgs)
		default:
			r, _ := regexp.Compile("\\d")
			if r.MatchString(input) {
				id, _ := strconv.ParseInt(input, 10, 32)
				if int(id) < len(pkgs)-1 && id > 0 {
					selected = b.appendSelected(selected, id)
				} else {
					fmt.Printf("Invalid id: %s\n", input)
					goto choose
				}
			} else {
				fmt.Printf("Unknown option: %s\n", input)
				goto choose
			}
			i = i - 1
		}
	}

	return selectedPackages
}
