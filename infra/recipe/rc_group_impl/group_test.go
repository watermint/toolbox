package rc_group_impl

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"strings"
	"testing"
)

func TestGroupImpl_AddNoRecipeGroup(t *testing.T) {
	rg := NewGroup()
	leafLeft := &MockSpec{path: []string{"content", "member"}, name: "list"}
	leafRight := &MockSpec{path: []string{"content", "policy"}, name: "list"}

	rg.Add(leafLeft)
	rg.Add(leafRight)

	if len(rg.SubGroups()) != 1 {
		t.Error(rg.SubGroups())
	}

	gContent := rg.SubGroups()["content"]
	if len(gContent.SubGroups()) != 2 {
		t.Error(gContent.SubGroups())
	}
	gMember := gContent.SubGroups()["member"]
	if len(gMember.SubGroups()) != 0 {
		t.Error(gMember.SubGroups())
	}
	gPolicy := gContent.SubGroups()["policy"]
	if len(gPolicy.SubGroups()) != 0 {
		t.Error(gContent.SubGroups())
	}

	msgs := app_msg_container_impl.NewSingleWithMessagesForTest(map[string]string{})
	usage := app_ui.MakeMarkdown(msgs, func(ui app_ui.UI) {
		gContent.PrintUsage(ui, app_definitions.ExecutableName, "x.y.z")
	})
	if strings.Contains(usage, "recipe.content.member.member.title") {
		t.Error(usage)
	}
	if strings.Contains(usage, "recipe.content.member.policy.title") {
		t.Error(usage)
	}
}
