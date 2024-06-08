/**
ASSIGNMENT
Let's clean up Textio's authentication logic. We store our user's authentication data inside
an authenticationInfo struct.
We need a method that can take that data and return a basic authorization string.

The format of the string should be:

Authorization: Basic USERNAME:PASSWORD

Create a method on the authenticationInfo struct called getBasicAuth
that returns the formatted string.
*/

package structs

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (authenticationInfo authenticationInfo) getBasicAuth() string {
	authString := fmt.Sprintf("Authorization: Basic %s:%s", authenticationInfo.username, authenticationInfo.password)
	return authString
}
