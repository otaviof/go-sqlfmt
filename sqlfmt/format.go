package sqlfmt

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kanmu/go-sqlfmt/sqlfmt/lexer"
	"github.com/kanmu/go-sqlfmt/sqlfmt/parser"
	"github.com/kanmu/go-sqlfmt/sqlfmt/parser/group"
	"github.com/pkg/errors"
)

// format formats given sql statement and returns the formatted statement
func format(src string) (string, error) {
	t := lexer.NewTokenizer(src)
	tokens, err := t.GetTokens()
	if err != nil {
		return src, errors.Wrap(err, "Tokenize failed")
	}

	rs, err := parser.ParseTokens(tokens)
	if err != nil {
		return src, errors.Wrap(err, "ParseTokens failed")
	}

	/*
		sql, err := parse(tokens)
		if err != nil {
			return err
		}


		buf := &bytes.Buffer{}
		if err := write(buf, sql); err != nil {
			return err
		}


		or

		printer := &printer{
			buf: &bytes.Buffer{},
			options: options,
		}

		if err := printer.print(sql); err != nil {
			return err
		}
		if compare(printer.buf.String(), sr

		if !compare(src, buf.String()) {
			return err
		}

		return buf.String()

	*/

	res, err := getFormattedStmt(rs)
	if err != nil {
		return src, errors.Wrap(err, "getFormattedStmt failed")
	}

	if !compare(src, res) {
		return src, fmt.Errorf("the formatted statement has diffed from the source")
	}
	return res, nil
}

func getFormattedStmt(rs []group.Reindenter) (string, error) {
	var buf bytes.Buffer

	for _, r := range rs {
		if err := r.Reindent(&buf); err != nil {
			return "", errors.Wrap(err, "Reindent failed")
		}
	}
	return buf.String(), nil
}

// returns false if the value of formatted statement  (without any space) differs from source statement
func compare(src string, res string) bool {
	before := removeSpace(src)
	after := removeSpace(res)

	if v := strings.Compare(before, after); v != 0 {
		return false
	}
	return true
}

// removes whitespaces and new lines from src
func removeSpace(src string) string {
	var result []rune
	for _, r := range src {
		if string(r) == "\n" || string(r) == " " || string(r) == "\t" || string(r) == "　" {
			continue
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
