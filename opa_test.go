package main

import (
	"fmt"
)

func allUsersShouldBeAdmins() error {
	if authz() {
		return nil
	}
	return fmt.Errorf("the user is not an admin")
}

func iValidateTheFiles() error {
	if len(inputFile) == 0 || len(regoFile) == 0 {
		return fmt.Errorf("either inputFile size is %d or regoFile size is %d", len(inputFile), len(regoFile))
	}
	return nil
}

func thereIsAJsonFile() error {
	if inputErr != nil {
		return fmt.Errorf("%d when loading file", inputErr)
	}
	return nil
}

func thereIsARegoFile() error {
	if regoErr != nil {
		return fmt.Errorf("%d when loading file", regoErr)
	}
	return nil
}

