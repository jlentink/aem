package main

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
	"regexp"
	"strconv"
)

func newBundlePicker() bundlePicker {
	return bundlePicker{
		sliceUtil: new(sliceUtil),
	}
}

type bundlePicker struct {
	sliceUtil *sliceUtil
}

func (b *bundlePicker) addBundlesToTable(t table.Writer, bundles []bundle) {
	for i, bundle := range bundles {
		t.AppendRow(table.Row{i + 1, bundle.ID, bundle.Name, bundle.SymbolicName, bundle.Version, bundle.State})
	}
}

func (b *bundlePicker) appendSelected(list []int64, selected int64) []int64 {
	if !b.sliceUtil.inSliceInt64(list, selected) {
		list = append(list, selected)
	}
	return list
}

func (b *bundlePicker) getSelectedBundles(selected []int64, bundles []bundle) []bundle {
	selectedBundles := make([]bundle, 0)
	for _, bundle := range selected {
		selectedBundles = append(selectedBundles, bundles[bundle-1])
	}
	return selectedBundles
}

func (b *bundlePicker) picker(instance aemInstanceConfig) []bundle {
	http := new(httpRequests)
	bundles := http.listBundles(instance)
	pageSize := 20
	selected := make([]int64, 0)
	selectedBundles := make([]bundle, 0)
	writer := new(tableWriter)

	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Id", "bundle", "Symbolic name", "Version", "Status"})
	t.SetPageSize(pageSize)
	t.SetOutputMirror(writer)
	b.addBundlesToTable(t, bundles.Data)
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
			return make([]bundle, 0)
		case "d":
			return b.getSelectedBundles(selected, bundles.Data)
		default:
			r, _ := regexp.Compile(`\d`)
			if r.MatchString(input) {
				id, _ := strconv.ParseInt(input, 10, 32)
				if int(id) < len(bundles.Data)-1 && id > 0 {
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
	return selectedBundles
}
