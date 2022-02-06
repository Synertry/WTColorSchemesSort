/*
 *           Copyright Synertry 2022.
 *  Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE_1_0.txt or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

//WTColorSchemeSort is a json array sorter for the color schemes of Windows Terminal settings.json
package WTColorSchemesSort

import (
	"fmt"
	"strings"
)

// Sort needs the array from the schemes key including the brackets [] as string
// returns sorted json string
func Sort(str string) (string, error) {
	sortedSchemes, err := parseSchemes(str)
	if err != nil {
		return fmt.Sprintf("%+v", sortedSchemes), err
	}

	sortedJSON, err := PrettyStruct(sortedSchemes)
	sortedJSON = fmt.Sprintf("\t%s", strings.Trim(sortedJSON, "{}\r\n\t\"schemes: "))
	if err != nil {
		return sortedJSON, err
	}

	return sortedJSON, err
}
