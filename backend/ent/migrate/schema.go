// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AdminrepairsColumns holds the columns for the "adminrepairs" table.
	AdminrepairsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "numberrepair", Type: field.TypeString, Unique: true, Size: 6},
		{Name: "equipmentdamate", Type: field.TypeString, Size: 75},
		{Name: "repairinformation", Type: field.TypeString, Size: 500},
		{Name: "fix_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "personal_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// AdminrepairsTable holds the schema information for the "adminrepairs" table.
	AdminrepairsTable = &schema.Table{
		Name:       "adminrepairs",
		Columns:    AdminrepairsColumns,
		PrimaryKey: []*schema.Column{AdminrepairsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "adminrepairs_fixes_fix",
				Columns: []*schema.Column{AdminrepairsColumns[4]},

				RefColumns: []*schema.Column{FixesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "adminrepairs_personals_personal",
				Columns: []*schema.Column{AdminrepairsColumns[5]},

				RefColumns: []*schema.Column{PersonalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "adminrepairs_products_product",
				Columns: []*schema.Column{AdminrepairsColumns[6]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BrandsColumns holds the columns for the "brands" table.
	BrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brandname", Type: field.TypeString, Unique: true},
	}
	// BrandsTable holds the schema information for the "brands" table.
	BrandsTable = &schema.Table{
		Name:        "brands",
		Columns:     BrandsColumns,
		PrimaryKey:  []*schema.Column{BrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "customername", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "phonenumber", Type: field.TypeString, Size: 10},
		{Name: "identificationnumber", Type: field.TypeString, Size: 13},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "personal_id", Type: field.TypeInt, Nullable: true},
		{Name: "title_id", Type: field.TypeInt, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "customers_genders_customer",
				Columns: []*schema.Column{CustomersColumns[5]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "customers_personals_customer",
				Columns: []*schema.Column{CustomersColumns[6]},

				RefColumns: []*schema.Column{PersonalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "customers_titles_customer",
				Columns: []*schema.Column{CustomersColumns[7]},

				RefColumns: []*schema.Column{TitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "departmentname", Type: field.TypeString, Unique: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FixesColumns holds the columns for the "fixes" table.
	FixesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "productnumber", Type: field.TypeString},
		{Name: "problemtype", Type: field.TypeString, Size: 100},
		{Name: "queue", Type: field.TypeString, Unique: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "fixbrand_id", Type: field.TypeInt, Nullable: true},
		{Name: "fixcomtype_id", Type: field.TypeInt, Nullable: true},
		{Name: "personal_id", Type: field.TypeInt, Nullable: true},
	}
	// FixesTable holds the schema information for the "fixes" table.
	FixesTable = &schema.Table{
		Name:       "fixes",
		Columns:    FixesColumns,
		PrimaryKey: []*schema.Column{FixesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "fixes_customers_fix",
				Columns: []*schema.Column{FixesColumns[5]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "fixes_fixbrands_fix",
				Columns: []*schema.Column{FixesColumns[6]},

				RefColumns: []*schema.Column{FixbrandsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "fixes_fixcomtypes_fix",
				Columns: []*schema.Column{FixesColumns[7]},

				RefColumns: []*schema.Column{FixcomtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "fixes_personals_fix",
				Columns: []*schema.Column{FixesColumns[8]},

				RefColumns: []*schema.Column{PersonalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FixbrandsColumns holds the columns for the "fixbrands" table.
	FixbrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fixbrandname", Type: field.TypeString, Unique: true},
	}
	// FixbrandsTable holds the schema information for the "fixbrands" table.
	FixbrandsTable = &schema.Table{
		Name:        "fixbrands",
		Columns:     FixbrandsColumns,
		PrimaryKey:  []*schema.Column{FixbrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FixcomtypesColumns holds the columns for the "fixcomtypes" table.
	FixcomtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fixcomtypename", Type: field.TypeString, Unique: true},
	}
	// FixcomtypesTable holds the schema information for the "fixcomtypes" table.
	FixcomtypesTable = &schema.Table{
		Name:        "fixcomtypes",
		Columns:     FixcomtypesColumns,
		PrimaryKey:  []*schema.Column{FixcomtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gendername", Type: field.TypeString, Unique: true},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PaymentTypesColumns holds the columns for the "payment_types" table.
	PaymentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typename", Type: field.TypeString, Unique: true},
	}
	// PaymentTypesTable holds the schema information for the "payment_types" table.
	PaymentTypesTable = &schema.Table{
		Name:        "payment_types",
		Columns:     PaymentTypesColumns,
		PrimaryKey:  []*schema.Column{PaymentTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PersonalsColumns holds the columns for the "personals" table.
	PersonalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "personalname", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "title_id", Type: field.TypeInt, Nullable: true},
	}
	// PersonalsTable holds the schema information for the "personals" table.
	PersonalsTable = &schema.Table{
		Name:       "personals",
		Columns:    PersonalsColumns,
		PrimaryKey: []*schema.Column{PersonalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "personals_departments_personal",
				Columns: []*schema.Column{PersonalsColumns[4]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "personals_genders_personal",
				Columns: []*schema.Column{PersonalsColumns[5]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "personals_titles_personal",
				Columns: []*schema.Column{PersonalsColumns[6]},

				RefColumns: []*schema.Column{TitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "productname", Type: field.TypeString},
		{Name: "numberofproduct", Type: field.TypeString},
		{Name: "price", Type: field.TypeString},
		{Name: "Brand", Type: field.TypeInt, Nullable: true},
		{Name: "Personal", Type: field.TypeInt, Nullable: true},
		{Name: "Typeproduct", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "products_brands_product",
				Columns: []*schema.Column{ProductsColumns[4]},

				RefColumns: []*schema.Column{BrandsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "products_personals_product",
				Columns: []*schema.Column{ProductsColumns[5]},

				RefColumns: []*schema.Column{PersonalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "products_typeproducts_product",
				Columns: []*schema.Column{ProductsColumns[6]},

				RefColumns: []*schema.Column{TypeproductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReceiptsColumns holds the columns for the "receipts" table.
	ReceiptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "serviceprovider", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "productname", Type: field.TypeString},
		{Name: "adminrepair_id", Type: field.TypeInt, Nullable: true},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "paymenttype_id", Type: field.TypeInt, Nullable: true},
		{Name: "personal_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// ReceiptsTable holds the schema information for the "receipts" table.
	ReceiptsTable = &schema.Table{
		Name:       "receipts",
		Columns:    ReceiptsColumns,
		PrimaryKey: []*schema.Column{ReceiptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "receipts_adminrepairs_receipt",
				Columns: []*schema.Column{ReceiptsColumns[5]},

				RefColumns: []*schema.Column{AdminrepairsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "receipts_customers_receipt",
				Columns: []*schema.Column{ReceiptsColumns[6]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "receipts_payment_types_receipt",
				Columns: []*schema.Column{ReceiptsColumns[7]},

				RefColumns: []*schema.Column{PaymentTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "receipts_personals_receipt",
				Columns: []*schema.Column{ReceiptsColumns[8]},

				RefColumns: []*schema.Column{PersonalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "receipts_products_receipt",
				Columns: []*schema.Column{ReceiptsColumns[9]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TitlesColumns holds the columns for the "titles" table.
	TitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "titlename", Type: field.TypeString, Unique: true},
	}
	// TitlesTable holds the schema information for the "titles" table.
	TitlesTable = &schema.Table{
		Name:        "titles",
		Columns:     TitlesColumns,
		PrimaryKey:  []*schema.Column{TitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TypeproductsColumns holds the columns for the "typeproducts" table.
	TypeproductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typeproductname", Type: field.TypeString, Unique: true},
	}
	// TypeproductsTable holds the schema information for the "typeproducts" table.
	TypeproductsTable = &schema.Table{
		Name:        "typeproducts",
		Columns:     TypeproductsColumns,
		PrimaryKey:  []*schema.Column{TypeproductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminrepairsTable,
		BrandsTable,
		CustomersTable,
		DepartmentsTable,
		FixesTable,
		FixbrandsTable,
		FixcomtypesTable,
		GendersTable,
		PaymentTypesTable,
		PersonalsTable,
		ProductsTable,
		ReceiptsTable,
		TitlesTable,
		TypeproductsTable,
	}
)

func init() {
	AdminrepairsTable.ForeignKeys[0].RefTable = FixesTable
	AdminrepairsTable.ForeignKeys[1].RefTable = PersonalsTable
	AdminrepairsTable.ForeignKeys[2].RefTable = ProductsTable
	CustomersTable.ForeignKeys[0].RefTable = GendersTable
	CustomersTable.ForeignKeys[1].RefTable = PersonalsTable
	CustomersTable.ForeignKeys[2].RefTable = TitlesTable
	FixesTable.ForeignKeys[0].RefTable = CustomersTable
	FixesTable.ForeignKeys[1].RefTable = FixbrandsTable
	FixesTable.ForeignKeys[2].RefTable = FixcomtypesTable
	FixesTable.ForeignKeys[3].RefTable = PersonalsTable
	PersonalsTable.ForeignKeys[0].RefTable = DepartmentsTable
	PersonalsTable.ForeignKeys[1].RefTable = GendersTable
	PersonalsTable.ForeignKeys[2].RefTable = TitlesTable
	ProductsTable.ForeignKeys[0].RefTable = BrandsTable
	ProductsTable.ForeignKeys[1].RefTable = PersonalsTable
	ProductsTable.ForeignKeys[2].RefTable = TypeproductsTable
	ReceiptsTable.ForeignKeys[0].RefTable = AdminrepairsTable
	ReceiptsTable.ForeignKeys[1].RefTable = CustomersTable
	ReceiptsTable.ForeignKeys[2].RefTable = PaymentTypesTable
	ReceiptsTable.ForeignKeys[3].RefTable = PersonalsTable
	ReceiptsTable.ForeignKeys[4].RefTable = ProductsTable
}
