// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsEventMetadataUnion()                                {}
func (UnionString) ImplementsEventInputMetadataUnionParam()                      {}
func (UnionString) ImplementsFilterTypeMeterFilterConditionListValueUnionParam() {}
func (UnionString) ImplementsFilterTypeMeterFilterConditionListValueUnion()      {}

type UnionBool bool

func (UnionBool) ImplementsEventMetadataUnion()                                {}
func (UnionBool) ImplementsEventInputMetadataUnionParam()                      {}
func (UnionBool) ImplementsFilterTypeMeterFilterConditionListValueUnionParam() {}
func (UnionBool) ImplementsFilterTypeMeterFilterConditionListValueUnion()      {}

type UnionFloat float64

func (UnionFloat) ImplementsEventMetadataUnion()                                {}
func (UnionFloat) ImplementsEventInputMetadataUnionParam()                      {}
func (UnionFloat) ImplementsFilterTypeMeterFilterConditionListValueUnionParam() {}
func (UnionFloat) ImplementsFilterTypeMeterFilterConditionListValueUnion()      {}
