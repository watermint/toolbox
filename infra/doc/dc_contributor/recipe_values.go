package dc_contributor

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"reflect"
	"strconv"
)

type RecipeValues struct {
}

func (z RecipeValues) DocId() dc_index.DocId {
	return dc_index.DocContributorRecipeValues
}

func (z RecipeValues) DocDesc() app_msg.Message {
	return MDeveloper.DeveloperDesc
}

func (z RecipeValues) Sections() []dc_section.Section {
	return []dc_section.Section{
		&RecipeValuesSection{},
	}
}

type RecipeValuesSection struct {
}

func (z RecipeValuesSection) Title() app_msg.Message {
	return MDeveloper.RecipeValueTitle
}

func (z RecipeValuesSection) Body(ui app_ui.UI) {
	connTypes := make([]string, 0)
	connTypeValues := make(map[string]rc_recipe.Value)

	ui.WithTable("ValueTypes", func(t app_ui.Table) {
		t.Header(
			MDeveloper.RecipeValueTypeImpl,
			MDeveloper.RecipeValueTypeConn,
			MDeveloper.RecipeValueTypeConns,
			MDeveloper.RecipeValueTypeCustomValueText,
			MDeveloper.RecipeValueTypeErrorHandler,
			MDeveloper.RecipeValueTypeFeed,
			MDeveloper.RecipeValueTypeGridDataInput,
			MDeveloper.RecipeValueTypeGridDataOutput,
			MDeveloper.RecipeValueTypeJsonInput,
			MDeveloper.RecipeValueTypeMessage,
			MDeveloper.RecipeValueTypeMessages,
			MDeveloper.RecipeValueTypeReport,
			MDeveloper.RecipeValueTypeReports,
			MDeveloper.RecipeValueTypeTextInput,
		)

		for _, v := range rc_value.ValueTypes {
			vv := v.Init()
			vt := reflect.TypeOf(vv)
			if vt.Kind() == reflect.Pointer {
				vt = reflect.ValueOf(vv).Elem().Type()
			}
			var implName string
			if vt.PkgPath() != "" {
				implName = vt.PkgPath() + "." + vt.Name()
			} else {
				implName = vt.Name()
			}

			_, isValueTypeConn := v.(rc_recipe.ValueConn)
			_, isValueTypeConns := v.(rc_recipe.ValueConns)
			_, isValueTypeCustomValueText := v.(rc_recipe.ValueCustomValueText)
			_, isValueTypeErrorHandler := v.(rc_recipe.ValueErrorHandler)
			_, isValueTypeFeed := v.(rc_recipe.ValueFeed)
			_, isValueTypeGridDataInput := v.(rc_recipe.ValueGridDataInput)
			_, isValueTypeGridDataOutput := v.(rc_recipe.ValueGridDataOutput)
			_, isValueTypeJsonInput := v.(rc_recipe.ValueJsonInput)
			_, isValueTypeMessage := v.(rc_recipe.ValueMessage)
			_, isValueTypeMessages := v.(rc_recipe.ValueMessages)
			_, isValueTypeReport := v.(rc_recipe.ValueReport)
			_, isValueTypeReports := v.(rc_recipe.ValueReports)
			_, isValueTypeTextInput := v.(rc_recipe.ValueTextInput)

			if isValueTypeConn {
				connTypes = append(connTypes, implName)
				connTypeValues[implName] = v
			}

			t.RowRaw(
				implName,
				strconv.FormatBool(isValueTypeConn),
				strconv.FormatBool(isValueTypeConns),
				strconv.FormatBool(isValueTypeCustomValueText),
				strconv.FormatBool(isValueTypeErrorHandler),
				strconv.FormatBool(isValueTypeFeed),
				strconv.FormatBool(isValueTypeGridDataInput),
				strconv.FormatBool(isValueTypeGridDataOutput),
				strconv.FormatBool(isValueTypeJsonInput),
				strconv.FormatBool(isValueTypeMessage),
				strconv.FormatBool(isValueTypeMessages),
				strconv.FormatBool(isValueTypeReport),
				strconv.FormatBool(isValueTypeReports),
				strconv.FormatBool(isValueTypeTextInput),
			)
		}
	})

	ui.SubHeader(MDeveloper.RecipeValueConnValueTypes)

	ui.WithTable("ValueTypes", func(t app_ui.Table) {
		t.Header(
			MDeveloper.RecipeValueTypeImpl,
			MDeveloper.RecipeValueTypeCustomValueText,
			MDeveloper.RecipeValueConnServiceName,
			MDeveloper.RecipeValueConnScopeLabel,
		)
		for _, ct := range connTypes {
			v := connTypeValues[ct]
			_, isValueTypeCustomValueText := v.(rc_recipe.ValueCustomValueText)
			if !isValueTypeCustomValueText {
				l := esl.Default()
				l.Error("Found Connection Value Type without Custom Value Text", esl.String("Impl", ct))
				panic("found connection value type without custom value text")
			}

			vc := v.(rc_recipe.ValueConn)
			conn, ok := vc.Conn()
			if !ok {
				l := esl.Default()
				l.Error("Found Connection Value Type without valid connection")
				panic("found connection value type without valid connection")
			}

			t.RowRaw(
				ct,
				strconv.FormatBool(isValueTypeCustomValueText),
				conn.AppKeyName(),
				conn.ScopeLabel(),
			)
		}
	})
}
