package buy

import (
	"strconv"
	"testing"
)

type RegexpExpect struct {
	ItemsNumber int
	TokenIndex  int
	Token       string
}

func Test_Regexp_Get_Command(t *testing.T) {
	var testCases = []struct {
		Message string
		Want    RegexpExpect
	}{
		{
			Message: "[開團]",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[ 開團]",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[開團 ]",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[ 開團 ]",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[ 開團 ]",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[開團 ]     ",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[   開團 ]     ",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[開團    ]     ",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "[   開團    ]     ",
			Want:    RegexpExpect{ItemsNumber: 1, TokenIndex: 3, Token: "開團"},
		},
		{
			Message: "   [開團]",
			Want:    RegexpExpect{ItemsNumber: 0},
		},
		{
			Message: "     ",
			Want:    RegexpExpect{ItemsNumber: 0},
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			token := parse.FindAllStringSubmatch(tc.Message, 1)
			if len(token) != tc.Want.ItemsNumber {
				t.Fatal("len = ", len(token))
			}
			if len(token) > 0 {
				if token[0][tc.Want.TokenIndex] != tc.Want.Token {
					t.Errorf("GOT: '%v', WANT: '%v'", token[0][tc.Want.TokenIndex], tc.Want.Token)
				} else {
					t.Logf("MESSAGE: '%v', GOT: '%v'", tc.Message, tc.Want.Token)
				}
			} else {
				t.Logf("MESSAGE: '%v' is not a command", tc.Message)
			}

		})
	}
}
