package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Switch node
type Switch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Cond         node.Node
	CaseList     *CaseList
}

// NewSwitch node constructor
func NewSwitch(Cond node.Node, CaseList *CaseList) *Switch {
	return &Switch{
		FreeFloating: nil,
		Cond:         Cond,
		CaseList:     CaseList,
	}
}

// SetPosition sets node position
func (n *Switch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Switch) GetPosition() *position.Position {
	return n.Position
}

func (n *Switch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Switch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Switch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.CaseList != nil {
		v.EnterChildNode("CaseList", n)
		n.CaseList.Walk(v)
		v.LeaveChildNode("CaseList", n)
	}

	v.LeaveNode(n)
}
