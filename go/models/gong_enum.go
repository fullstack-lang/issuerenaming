// generated by ModelGongEnumFileTemplate
package models

// insertion point of enum utility functions
// Utility function for GONG__ExpressionType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (gong__expressiontype GONG__ExpressionType) ToString() (res string) {

	// migration of former implementation of enum
	switch gong__expressiontype {
	// insertion code per enum code
	case GONG__STRUCT_INSTANCE:
		res = "STRUCT_INSTANCE"
	case GONG__FIELD_OR_CONST_VALUE:
		res = "FIELD_OR_CONST_VALUE"
	case GONG__FIELD_VALUE:
		res = "FIELD_VALUE"
	case GONG__ENUM_CAST_INT:
		res = "ENUM_CAST_INT"
	case GONG__ENUM_CAST_STRING:
		res = "ENUM_CAST_STRING"
	case GONG__IDENTIFIER_CONST:
		res = "IDENTIFIER_CONST"
	}
	return
}

func (gong__expressiontype *GONG__ExpressionType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STRUCT_INSTANCE":
		*gong__expressiontype = GONG__STRUCT_INSTANCE
	case "FIELD_OR_CONST_VALUE":
		*gong__expressiontype = GONG__FIELD_OR_CONST_VALUE
	case "FIELD_VALUE":
		*gong__expressiontype = GONG__FIELD_VALUE
	case "ENUM_CAST_INT":
		*gong__expressiontype = GONG__ENUM_CAST_INT
	case "ENUM_CAST_STRING":
		*gong__expressiontype = GONG__ENUM_CAST_STRING
	case "IDENTIFIER_CONST":
		*gong__expressiontype = GONG__IDENTIFIER_CONST
	default:
		return errUnkownEnum
	}
	return
}

func (gong__expressiontype *GONG__ExpressionType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GONG__STRUCT_INSTANCE":
		*gong__expressiontype = GONG__STRUCT_INSTANCE
	case "GONG__FIELD_OR_CONST_VALUE":
		*gong__expressiontype = GONG__FIELD_OR_CONST_VALUE
	case "GONG__FIELD_VALUE":
		*gong__expressiontype = GONG__FIELD_VALUE
	case "GONG__ENUM_CAST_INT":
		*gong__expressiontype = GONG__ENUM_CAST_INT
	case "GONG__ENUM_CAST_STRING":
		*gong__expressiontype = GONG__ENUM_CAST_STRING
	case "GONG__IDENTIFIER_CONST":
		*gong__expressiontype = GONG__IDENTIFIER_CONST
	default:
		return errUnkownEnum
	}
	return
}

func (gong__expressiontype *GONG__ExpressionType) ToCodeString() (res string) {

	switch *gong__expressiontype {
	// insertion code per enum code
	case GONG__STRUCT_INSTANCE:
		res = "GONG__STRUCT_INSTANCE"
	case GONG__FIELD_OR_CONST_VALUE:
		res = "GONG__FIELD_OR_CONST_VALUE"
	case GONG__FIELD_VALUE:
		res = "GONG__FIELD_VALUE"
	case GONG__ENUM_CAST_INT:
		res = "GONG__ENUM_CAST_INT"
	case GONG__ENUM_CAST_STRING:
		res = "GONG__ENUM_CAST_STRING"
	case GONG__IDENTIFIER_CONST:
		res = "GONG__IDENTIFIER_CONST"
	}
	return
}

// Last line of the template
