/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : propertyloader.go
 *creationDate: AUG 04 2014
 *updatedDate : 
 *updatedBy   : 
 *purpose     : 
 *****************************************************************************/
package loderutil

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var re *regexp.Regexp
var pat = "[#].*\\n|\\s+\\n|\\S+[=]|.*\n"

func init() {
	re, _ = regexp.Compile(pat)
}
func Load(filename string, dest map[string]string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buff := make([]byte, fi.Size())
	f.Read(buff)
	f.Close()
	str := string(buff)
	if !strings.HasSuffix(str, "\n") {
		return errors.New("Loder property does not end with a newline character.")
	}
	s2 := re.FindAllString(str, -1)
	for i := 0; i < len(s2); {
		if strings.HasPrefix(s2[i], "#") {
			i++
		} else if strings.HasSuffix(s2[i], "=") {
			key := strings.ToLower(s2[i])[0 : len(s2[i])-1]
			i++
			if strings.HasSuffix(s2[i], "\n") {
				val := s2[i][0 : len(s2[i])-1]
				if strings.HasSuffix(val, "\r") {
					val = val[0 : len(val)-1]
				}
				i++
				dest[key] = val
			}
		} else if strings.Index(" \t\r\n", s2[i][0:1]) > -1 {
			i++
		} else {
			return errors.New("Unable to process line in Loder property file containing " + s2[i])
		}
	}
	return nil
}
