package pg

// Framework code is generated by the generator.

import (
	"github.com/bytebase/bytebase/plugin/advisor"
	"github.com/bytebase/bytebase/plugin/advisor/db"
	"github.com/bytebase/bytebase/plugin/parser/ast"
)

var (
	_ advisor.Advisor = (*StatementDisallowAddColumnWithDefaultAdvisor)(nil)
	_ ast.Visitor     = (*statementDisallowAddColumnWithDefaultChecker)(nil)
)

func init() {
	advisor.Register(db.Postgres, advisor.PostgreSQLDisallowAddColumnWithDefault, &StatementDisallowAddColumnWithDefaultAdvisor{})
}

// StatementDisallowAddColumnWithDefaultAdvisor is the advisor checking for to disallow add column with default.
type StatementDisallowAddColumnWithDefaultAdvisor struct {
}

// Check checks for to disallow add column with default.
func (*StatementDisallowAddColumnWithDefaultAdvisor) Check(ctx advisor.Context, statement string) ([]advisor.Advice, error) {
	stmtList, errAdvice := parseStatement(statement)
	if errAdvice != nil {
		return errAdvice, nil
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementDisallowAddColumnWithDefaultChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.line = stmt.LastLine()
		ast.Walk(checker, stmt)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type statementDisallowAddColumnWithDefaultChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	line       int
}

// Visit implements ast.Visitor interface.
func (checker *statementDisallowAddColumnWithDefaultChecker) Visit(in ast.Node) ast.Visitor {
	if node, ok := in.(*ast.AddColumnListStmt); ok {
		for _, column := range node.ColumnList {
			if setDefault(column) {
				checker.adviceList = append(checker.adviceList, advisor.Advice{
					Status:  checker.level,
					Code:    advisor.StatementAddColumnWithDefault,
					Title:   checker.title,
					Content: "Adding column with DEFAULT will locked the whole table and rewriting each rows",
					Line:    checker.line,
				})
			}
		}
	}
	return checker
}

func setDefault(column *ast.ColumnDef) bool {
	for _, constraint := range column.ConstraintList {
		if constraint.Type == ast.ConstraintTypeDefault {
			return true
		}
	}

	return false
}