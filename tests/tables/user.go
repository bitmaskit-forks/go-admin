package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

// GetUserTable return the model of table user.
func GetUserTable(ctx *context.Context) (userTable table.Table) {

	userTable = table.NewDefaultTable(table.Config{
		Driver:     config.Get().Databases.GetDefault().Driver,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := userTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).Where("gender", "=", 1)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "men"
		}
		if model.Value == "1" {
			return "women"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "👦"},
		{Value: "1", Text: "👧"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "men"},
		{Value: "1", Text: "women"},
	})
	info.AddField("Experience", "experience", db.Tinyint).
		FieldFilterable(types.FilterType{FormType: form.Radio}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "0", Text: "one"},
			{Value: "1", Text: "two"},
			{Value: "3", Text: "three"},
		}).FieldHide()
	info.AddField("Drink", "drink", db.Tinyint).
		FieldFilterable(types.FilterType{FormType: form.Select}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "water", Text: "water"},
			{Value: "juice", Text: "juice"},
			{Value: "red bull", Text: "red bull"},
		}).FieldHide()
	info.AddField("City", "city", db.Varchar).FieldFilterable()
	info.AddField("Avatar", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().Image().
			SetSrc(`//quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png`).
			SetHeight("120").SetWidth("120").WithModal().GetContent()
	})
	info.AddField("CreatedAt", "created_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("UpdatedAt", "updated_at", db.Timestamp).FieldEditAble(editType.Datetime)

	// ===========================
	// Buttons
	// ===========================

	info.AddActionButton("google", action.Jump("https://google.com"))
	info.AddActionButton("Audit", action.Ajax("/admin/audit",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			fmt.Println("PostForm", ctx.PostForm())
			return true, "success", ""
		}))
	info.AddActionButton("Preview", action.PopUp("/admin/preview", "Preview",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>preview content</h2>"
		}))
	info.AddButton("jump", icon.User, action.JumpInNewTab("/admin/info/authors", "authors"))
	info.AddButton("popup", icon.Terminal, action.PopUp("/admin/popup", "Popup Example",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>hello world</h2>"
		}))
	info.AddButton("ajax", icon.Android, action.Ajax("/admin/ajax",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "Oh li get", ""
		}))
	info.AddSelectBox("gender", types.FieldOptions{
		{Value: "0", Text: "men"},
		{Value: "1", Text: "women"},
	}, action.FieldFilter("gender"))

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := userTable.GetForm()

	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Age", "age", db.Int, form.Number)
	formList.AddField("Homepage", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")
	formList.AddField("Email", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")
	formList.AddField("Birthday", "birthday", db.Varchar, form.Datetime).FieldDefault("2010-09-05")
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("IP", "ip", db.Varchar, form.Ip)
	formList.AddField("Cert", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})
	formList.AddField("Amout", "currency", db.Int, form.Currency)
	formList.AddField("Content", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`)

	formList.AddField("Switch", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("Will not be able to access when the site was off").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	formList.AddField("Fruit", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "Apple", Value: "apple"},
			{Text: "Banana", Value: "banana"},
			{Text: "Watermelon", Value: "watermelon"},
			{Text: "Pear", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"Pear"}
		})
	formList.AddField("Gender", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Boy", Value: "0"},
			{Text: "Girl", Value: "1"},
		})
	formList.AddField("Drink", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Beer", Value: "beer"},
			{Text: "Juice", Value: "juice"},
			{Text: "Water", Value: "water"},
			{Text: "Red bull", Value: "red bull"},
		}).FieldDefault("beer")
	formList.AddField("Work Experience", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "two years", Value: "0"},
			{Text: "three years", Value: "1"},
			{Text: "four years", Value: "2"},
			{Text: "five years", Value: "3"},
		}).FieldDefault("beer")
	formList.SetTabGroups(types.TabGroups{
		{"name", "age", "homepage", "email", "birthday", "password", "ip", "certificate", "currency", "content"},
		{"website", "fruit", "gender", "drink", "experience"},
	})
	formList.SetTabHeaders("input", "select")

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList.SetPostHook(func(values form2.Values) error {
		fmt.Println("userTable.GetForm().PostHook", values)
		return nil
	})

	return
}
