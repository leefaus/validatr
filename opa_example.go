package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"io/ioutil"
	"strconv"
)


func opa_example() {
	regoFile, err := ioutil.ReadFile("example.rego")
	regoString := string(regoFile)
	//fmt.Print(string(regoFile))

	ctx := context.TODO()
	query, err := rego.New(
		rego.Query("result := data.example.authz[key]"),
		rego.Module("example.rego", regoString),
	).PrepareForEval(ctx)

	if err != nil {
		// Handle error.
	}
	inputFile, err := ioutil.ReadFile("input.json")
	//fmt.Print(string(inputFile))

	var dat map[string]interface{}
	json.Unmarshal(inputFile, &dat)


	results, err := query.Eval(ctx, rego.EvalInput(dat))

	if err != nil {
		fmt.Errorf(err.Error())
	} else if len(results) == 0 {
		fmt.Println("results were 0")
	} else if result, ok := results[0].Bindings["result"].(bool); !ok {
		fmt.Sprintf("results were of unexpected type: %s", strconv.FormatBool(result))
	} else {
		// Handle result/decision.
		fmt.Printf("%+v",results)
	}

}
