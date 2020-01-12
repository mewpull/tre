package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/zegl/tre/compiler/lexer"
)

func TestAllocType(t *testing.T) {
	lexed := lexer.Lex(`var a int`)

	expected := FileNode{
		Instructions: []Node{
			&AllocNode{
				Name: []string{"a"},
				Type: &SingleTypeNode{TypeName: "int"},
			},
		},
	}

	assert.Equal(t, expected, Parse(lexed, false))
}

func TestAllocTypeWithValue(t *testing.T) {
	lexed := lexer.Lex(`var a int = 10`)

	expected := FileNode{
		Instructions: []Node{
			&AllocNode{
				Name: []string{"a"},
				Type: &SingleTypeNode{TypeName: "int"},
				Val:  []Node{&ConstantNode{Type: NUMBER, Value: 10}},
			},
		},
	}

	assert.Equal(t, expected, Parse(lexed, false))
}

func TestAllocImplicitTypeValue(t *testing.T) {
	lexed := lexer.Lex(`var a = 10`)

	expected := FileNode{
		Instructions: []Node{
			&AllocNode{
				Name: []string{"a"},
				Val:  []Node{&ConstantNode{Type: NUMBER, Value: 10}},
			},
		},
	}

	assert.Equal(t, expected, Parse(lexed, false))
}

func TestAllocMultiWithType(t *testing.T) {
	lexed := lexer.Lex(`var a, b int`)

	expected := FileNode{
		Instructions: []Node{
			&AllocNode{
				Name: []string{"a", "b"},
				Type: &SingleTypeNode{TypeName: "int"},
			},
		},
	}

	assert.Equal(t, expected, Parse(lexed, false))
}

func TestAllocGroup(t *testing.T) {
	lexed := lexer.Lex(`var (
	a int
	b, c uint8
	d = 10
)`)

	expected := FileNode{
		Instructions: []Node{
			&AllocGroup{
				Allocs: []*AllocNode{
					{
						Name: []string{"a"},
						Type: &SingleTypeNode{TypeName: "int"},
					},
					{
						Name: []string{"b", "c"},
						Type: &SingleTypeNode{TypeName: "uint8"},
					},
					{
						Name: []string{"d"},
						Val:  []Node{&ConstantNode{Type: NUMBER, Value: 10}},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(lexed, false))
}
