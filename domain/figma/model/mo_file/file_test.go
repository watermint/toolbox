package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"testing"
)

func TestDocument(t *testing.T) {
	d := `{"document":{"id":"0:0","name":"Document","type":"DOCUMENT","scrollBehavior":"SCROLLS","children":[{"id":"0:1","name":"Page 1","type":"CANVAS","scrollBehavior":"SCROLLS","children":[{"id":"108:9","name":"Frame 1","type":"FRAME","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","children":[{"id":"108:11","name":"Star 1","type":"STAR","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-248,"y":102,"width":142,"height":103},"absoluteRenderBounds":{"x":-244.52500915527344,"y":102,"width":135.05001831054688,"height":93.16436767578125},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]}],"absoluteBoundingBox":{"x":-333,"y":56,"width":335,"height":190},"absoluteRenderBounds":{"x":-333,"y":56,"width":335,"height":190},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"clipsContent":true,"background":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","backgroundColor":{"r":1,"g":1,"b":1,"a":1},"effects":[]},{"id":"108:10","name":"Frame 2","type":"FRAME","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","children":[{"id":"101:3","name":"Rectangle 1","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-300,"y":-279,"width":282,"height":263},"absoluteRenderBounds":{"x":-300,"y":-279,"width":282,"height":263},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]}],"absoluteBoundingBox":{"x":-361,"y":-320,"width":400,"height":326},"absoluteRenderBounds":{"x":-361,"y":-320,"width":400,"height":326},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"clipsContent":true,"background":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","backgroundColor":{"r":1,"g":1,"b":1,"a":1},"effects":[]}],"backgroundColor":{"r":0.9607843160629272,"g":0.9607843160629272,"b":0.9607843160629272,"a":1},"prototypeStartNodeID":null,"flowStartingPoints":[],"prototypeDevice":{"type":"NONE","rotation":"NONE"}},{"id":"101:2","name":"Page 2","type":"CANVAS","scrollBehavior":"SCROLLS","children":[{"id":"101:4","name":"Rectangle 1","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-422,"y":-304,"width":139,"height":134},"absoluteRenderBounds":{"x":-422,"y":-304,"width":139,"height":134},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]},{"id":"101:5","name":"Rectangle 2","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-169,"y":-272,"width":102,"height":166},"absoluteRenderBounds":{"x":-169,"y":-272,"width":102,"height":166},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]},{"id":"108:7","name":"Frame 1","type":"FRAME","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","children":[{"id":"101:6","name":"Rectangle 3","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-370,"y":-40,"width":269,"height":173},"absoluteRenderBounds":{"x":-370,"y":-40,"width":269,"height":173},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]},{"id":"108:8","name":"Ellipse 1","type":"ELLIPSE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-169,"y":47,"width":127,"height":110},"absoluteRenderBounds":{"x":-169,"y":47,"width":127,"height":110},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[],"arcData":{"startingAngle":0,"endingAngle":6.2831854820251465,"innerRadius":0}}],"absoluteBoundingBox":{"x":-434,"y":-78,"width":562,"height":276},"absoluteRenderBounds":{"x":-434,"y":-78,"width":562,"height":276},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"clipsContent":true,"background":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":1,"g":1,"b":1,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","backgroundColor":{"r":1,"g":1,"b":1,"a":1},"effects":[]}],"backgroundColor":{"r":0.9607843160629272,"g":0.9607843160629272,"b":0.9607843160629272,"a":1},"prototypeStartNodeID":null,"flowStartingPoints":[],"prototypeDevice":{"type":"NONE","rotation":"NONE"}}]},"components":{},"componentSets":{},"schemaVersion":0,"styles":{},"name":"Untitled","lastModified":"2023-04-01T01:05:55Z","thumbnailUrl":"https://s3-alpha.figma.com/thumbnails/30a4bfd5-10df-4863-94b3-f7fe8c83ac0a?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAQ4GOSFWCX7GOXQGG%2F20230330%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230330T000000Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=8a556835cf18b253e0817dda62235e92a4352faa456566606aff054ce45ffc01","version":"3269678037","role":"owner","editorType":"figma","linkAccess":"view"}`
	doc := &Document{}
	j, err := es_json.ParseString(d)
	if err = j.Model(doc); err != nil {
		t.Error(err)
		return
	}
	if err := json.Unmarshal([]byte(d), doc); err != nil {
		t.Error(err)
		return
	}
	nodes := doc.Nodes()
	for _, n := range nodes {
		t.Log(n)
	}
	nodesWithPath := doc.NodesWithPath()
	for _, n := range nodesWithPath {
		t.Log(n)
	}
}

func TestPages(t *testing.T) {
	d := `{"document":{"id":"0:0","name":"Document","type":"DOCUMENT","scrollBehavior":"SCROLLS","children":[{"id":"0:1","name":"Page 1","type":"CANVAS","scrollBehavior":"SCROLLS","children":[{"id":"101:3","name":"Rectangle 1","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-300,"y":-279,"width":282,"height":263},"absoluteRenderBounds":{"x":-300,"y":-279,"width":282,"height":263},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]}],"backgroundColor":{"r":0.9607843160629272,"g":0.9607843160629272,"b":0.9607843160629272,"a":1},"prototypeStartNodeID":null,"flowStartingPoints":[],"prototypeDevice":{"type":"NONE","rotation":"NONE"}},{"id":"101:2","name":"Page 2","type":"CANVAS","scrollBehavior":"SCROLLS","children":[{"id":"101:4","name":"Rectangle 1","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-422,"y":-304,"width":139,"height":134},"absoluteRenderBounds":{"x":-422,"y":-304,"width":139,"height":134},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]},{"id":"101:5","name":"Rectangle 2","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-169,"y":-272,"width":102,"height":166},"absoluteRenderBounds":{"x":-169,"y":-272,"width":102,"height":166},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]},{"id":"101:6","name":"Rectangle 3","type":"RECTANGLE","scrollBehavior":"SCROLLS","blendMode":"PASS_THROUGH","absoluteBoundingBox":{"x":-370,"y":-40,"width":269,"height":173},"absoluteRenderBounds":{"x":-370,"y":-40,"width":269,"height":173},"constraints":{"vertical":"TOP","horizontal":"LEFT"},"fills":[{"blendMode":"NORMAL","type":"SOLID","color":{"r":0.8509804010391235,"g":0.8509804010391235,"b":0.8509804010391235,"a":1}}],"strokes":[],"strokeWeight":1,"strokeAlign":"INSIDE","effects":[]}],"backgroundColor":{"r":0.9607843160629272,"g":0.9607843160629272,"b":0.9607843160629272,"a":1},"prototypeStartNodeID":null,"flowStartingPoints":[],"prototypeDevice":{"type":"NONE","rotation":"NONE"}}]},"components":{},"componentSets":{},"schemaVersion":0,"styles":{},"name":"Untitled","lastModified":"2023-03-29T13:29:07Z","thumbnailUrl":"https://s3-alpha.figma.com/thumbnails/c59092b7-de93-4bcc-b893-5499b1108449?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAQ4GOSFWCX7GOXQGG%2F20230326%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230326T120000Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=0f6e6d41010458cf250056753c16889d94acf3a8a4f9d7cf54dc9382cc2f653e","version":"3253141199","role":"owner","editorType":"figma","linkAccess":"view"}`
	dj, err := es_json.ParseString(d)
	if err != nil {
		t.Error(err)
		return
	}
	doc := Document{}
	if err = dj.Model(&doc); err != nil {
		t.Error(err)
		return
	}
}

func TestNodeWithPath_Path(t *testing.T) {
	n := NodeWithPath{
		Name: []string{"Document", "Page 1"},
		Id:   []string{"0:0", "0:1"},
		Node: Node{},
	}

	if p := n.Path(" ", "_"); p != "Document_0_0 Page 1_0_1" {
		t.Error(p)
	}
}
