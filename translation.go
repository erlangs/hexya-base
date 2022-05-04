// Copyright 2018 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package base

import (
	"fmt"

	"github.com/erlangs/okoo/src/actions"
	"github.com/erlangs/okoo/src/models"
	"github.com/erlangs/okoo/src/models/types"
	"github.com/erlangs/pool/h"
	"github.com/erlangs/pool/m"
)

// TranslateFields opens the translation window for the given field
func translation_TranslateFields(rs m.TranslationSet, modelName string, id int64, fieldName models.FieldName) *actions.Action {
	fi := models.Registry.MustGet(modelName).FieldsGet(fieldName)[fieldName.JSON()]
	model := fmt.Sprintf("%sHexya%s", modelName, fi.Name)
	return &actions.Action{
		Name:     rs.T("Translate"),
		Type:     actions.ActionActWindow,
		Model:    model,
		ViewMode: "list",
		Domain:   fmt.Sprintf("[('record_id', '=', %d)]", id),
		Context:  types.NewContext().WithKey("default_record_id", id),
	}
}

func init() {
	models.NewModel("Translation")
	h.Translation().NewMethod("TranslateFields", translation_TranslateFields)
}
