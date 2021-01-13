// Code generated by entc, DO NOT EDIT.

package adminrepair

const (
	// Label holds the string label denoting the adminrepair type in the database.
	Label = "adminrepair"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEquipmentdamate holds the string denoting the equipmentdamate field in the database.
	FieldEquipmentdamate = "equipmentdamate"

	// EdgeAdminrepairPersonal holds the string denoting the adminrepairpersonal edge name in mutations.
	EdgeAdminrepairPersonal = "AdminrepairPersonal"
	// EdgeAdminrepairFix holds the string denoting the adminrepairfix edge name in mutations.
	EdgeAdminrepairFix = "AdminrepairFix"
	// EdgeAdminrepairProduct holds the string denoting the adminrepairproduct edge name in mutations.
	EdgeAdminrepairProduct = "AdminrepairProduct"

	// Table holds the table name of the adminrepair in the database.
	Table = "adminrepairs"
	// AdminrepairPersonalTable is the table the holds the AdminrepairPersonal relation/edge.
	AdminrepairPersonalTable = "adminrepairs"
	// AdminrepairPersonalInverseTable is the table name for the Personal entity.
	// It exists in this package in order to avoid circular dependency with the "personal" package.
	AdminrepairPersonalInverseTable = "personals"
	// AdminrepairPersonalColumn is the table column denoting the AdminrepairPersonal relation/edge.
	AdminrepairPersonalColumn = "personal_id"
	// AdminrepairFixTable is the table the holds the AdminrepairFix relation/edge.
	AdminrepairFixTable = "adminrepairs"
	// AdminrepairFixInverseTable is the table name for the Fix entity.
	// It exists in this package in order to avoid circular dependency with the "fix" package.
	AdminrepairFixInverseTable = "fixes"
	// AdminrepairFixColumn is the table column denoting the AdminrepairFix relation/edge.
	AdminrepairFixColumn = "fix_id"
	// AdminrepairProductTable is the table the holds the AdminrepairProduct relation/edge.
	AdminrepairProductTable = "adminrepairs"
	// AdminrepairProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	AdminrepairProductInverseTable = "products"
	// AdminrepairProductColumn is the table column denoting the AdminrepairProduct relation/edge.
	AdminrepairProductColumn = "product_id"
)

// Columns holds all SQL columns for adminrepair fields.
var Columns = []string{
	FieldID,
	FieldEquipmentdamate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Adminrepair type.
var ForeignKeys = []string{
	"fix_id",
	"personal_id",
	"product_id",
}

var (
	// EquipmentdamateValidator is a validator for the "equipmentdamate" field. It is called by the builders before save.
	EquipmentdamateValidator func(string) error
)
