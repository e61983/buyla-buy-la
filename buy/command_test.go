package buy

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_ParseCommand_SpaceTest(t *testing.T) {
	var testCases = []struct {
		UserID  string
		Message string
		Want    Command
	}{
		{
			UserID:  "TESTER",
			Message: "[開團] TEST",
			Want:    &OpenNewBuyLaCommand{UserID: "TESTER", ShopName: "TEST"},
		},
		{
			UserID:  "TESTER",
			Message: "[我要] FOO BAR BAR",
			Want:    &WantCommand{UserID: "TESTER", Goods: "FOO BAR BAR"},
		},
		{
			UserID:  "TESTER",
			Message: "[結單]",
			Want:    &CloseBuyLaCommand{UserID: "TESTER"},
		},
		{
			UserID:  "TESTER",
			Message: "@FOO [咪兔]",
			Want:    &MeTooCommand{UserID: "TESTER", TargetName: "FOO"},
		},
		{
			UserID:  "TESTER",
			Message: "[說明]",
			Want:    &HelpCommand{UserID: "TESTER"},
		},
		{
			UserID:  "TESTER",
			Message: "[加訂] FOO",
			Want:    &AttchCommand{UserID: "TESTER", Goods: "FOO"},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			command, err := ParseCommand(tc.UserID, tc.Message)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(command, tc.Want) {
				t.Errorf("GOT: %v, WANT: %#v", command, tc.Want)
			} else {
				t.Logf("MESSAGE: '%v', GOT: '%#v'", tc.Message, tc.Want)
			}
		})
	}
}
