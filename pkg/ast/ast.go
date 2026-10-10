package ast

import "rpp-parser/pkg/token"

type Node interface {
	TokenLiteral() string
}

type Property interface {
	Node
}

type Project struct {
	Properties []Property
}

func (p *Project) TokenLiteral() string {
	if len(p.Properties) > 0 {
		return p.Properties[0].TokenLiteral()
	}
	return "Project"
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Value
}

type Option struct {
	Token token.Token
	Name  *Identifier
	Value Property
}

func (o *Option) TokenLiteral() string {
	return o.Token.Value
}
