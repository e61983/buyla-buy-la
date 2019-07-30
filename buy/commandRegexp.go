package buy

import (
	"regexp"
)

// This regexp has three groups: mentionName, [Command], and Others
// [0]: match string
// [1]: @mentionName: has tail space, you have to remove it by your self.
// [2]: mentionName: has tail space, you have to remove it by your self.
// [3]: Command: it remove [] arealy.
// [4]: Others: maybe is Goods, or Shop Name

var parse = regexp.MustCompile(`^(@([^[]*))?\[[\s\n\t ]*([\p{Han}]*)[\s\n\t ]*\][\s\n\t ]*(.*)`)
