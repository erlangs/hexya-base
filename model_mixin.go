// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package base

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ToggleActive toggles the Active field of this object if it exists.
func modelMixin_ToggleActive(rs m.BaseMixinSet) {
	activeField, exists := rs.Collection().Model().Fields().Get("active")
	if !exists {
		return
	}
	if rs.Get(activeField).(bool) {
		rs.Set(activeField, false)
	} else {
		rs.Set(activeField, true)
	}
}

func modelMixin_Search(rs m.ModelMixinSet, cond q.ModelMixinCondition) m.ModelMixinSet {
	activeField, exists := rs.Collection().Model().Fields().Get("active")
	activeTest := !rs.Env().Context().HasKey("active_test") || rs.Env().Context().GetBool("active_test")
	if !exists || !activeTest || cond.HasField(activeField) {
		return rs.Super().Search(cond)
	}
	activeCond := q.ModelMixinCondition{
		Condition: models.Registry.MustGet(rs.ModelName()).Field(activeField).Equals(true),
	}
	cond = cond.AndCond(activeCond)
	return rs.Super().Search(cond)
}

func modelMixin_SearchAll(rs m.ModelMixinSet) m.ModelMixinSet {
	activeField, exists := rs.Collection().Model().Fields().Get("active")
	activeTest := !rs.Env().Context().HasKey("active_test") || rs.Env().Context().GetBool("active_test")
	if !exists || !activeTest {
		return rs.Super().SearchAll()
	}
	activeCond := q.ModelMixinCondition{
		Condition: models.Registry.MustGet(rs.ModelName()).Field(activeField).Equals(true),
	}
	return rs.Search(activeCond)
}

func init() {

	h.ModelMixin().NewMethod("ToggleActive", modelMixin_ToggleActive)
	h.ModelMixin().Methods().Search().Extend(modelMixin_Search)
	h.ModelMixin().Methods().SearchAll().Extend(modelMixin_SearchAll)
}
