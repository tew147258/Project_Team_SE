// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/receipt"
	"github.com/tanapon395/playlist-video/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	adminrepairFields := schema.Adminrepair{}.Fields()
	_ = adminrepairFields
	// adminrepairDescNumberrepair is the schema descriptor for numberrepair field.
	adminrepairDescNumberrepair := adminrepairFields[0].Descriptor()
	// adminrepair.NumberrepairValidator is a validator for the "numberrepair" field. It is called by the builders before save.
	adminrepair.NumberrepairValidator = func() func(string) error {
		validators := adminrepairDescNumberrepair.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(numberrepair string) error {
			for _, fn := range fns {
				if err := fn(numberrepair); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminrepairDescEquipmentdamate is the schema descriptor for equipmentdamate field.
	adminrepairDescEquipmentdamate := adminrepairFields[1].Descriptor()
	// adminrepair.EquipmentdamateValidator is a validator for the "equipmentdamate" field. It is called by the builders before save.
	adminrepair.EquipmentdamateValidator = func() func(string) error {
		validators := adminrepairDescEquipmentdamate.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(equipmentdamate string) error {
			for _, fn := range fns {
				if err := fn(equipmentdamate); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminrepairDescRepairinformation is the schema descriptor for repairinformation field.
	adminrepairDescRepairinformation := adminrepairFields[2].Descriptor()
	// adminrepair.RepairinformationValidator is a validator for the "repairinformation" field. It is called by the builders before save.
	adminrepair.RepairinformationValidator = func() func(string) error {
		validators := adminrepairDescRepairinformation.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(repairinformation string) error {
			for _, fn := range fns {
				if err := fn(repairinformation); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCustomername is the schema descriptor for Customername field.
	customerDescCustomername := customerFields[0].Descriptor()
	// customer.CustomernameValidator is a validator for the "Customername" field. It is called by the builders before save.
	customer.CustomernameValidator = customerDescCustomername.Validators[0].(func(string) error)
	// customerDescAddress is the schema descriptor for Address field.
	customerDescAddress := customerFields[1].Descriptor()
	// customer.AddressValidator is a validator for the "Address" field. It is called by the builders before save.
	customer.AddressValidator = customerDescAddress.Validators[0].(func(string) error)
	// customerDescPhonenumber is the schema descriptor for Phonenumber field.
	customerDescPhonenumber := customerFields[2].Descriptor()
	// customer.PhonenumberValidator is a validator for the "Phonenumber" field. It is called by the builders before save.
	customer.PhonenumberValidator = func() func(string) error {
		validators := customerDescPhonenumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Phonenumber string) error {
			for _, fn := range fns {
				if err := fn(_Phonenumber); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// customerDescIdentificationnumber is the schema descriptor for Identificationnumber field.
	customerDescIdentificationnumber := customerFields[3].Descriptor()
	// customer.IdentificationnumberValidator is a validator for the "Identificationnumber" field. It is called by the builders before save.
	customer.IdentificationnumberValidator = func() func(string) error {
		validators := customerDescIdentificationnumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Identificationnumber string) error {
			for _, fn := range fns {
				if err := fn(_Identificationnumber); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	fixFields := schema.Fix{}.Fields()
	_ = fixFields
	// fixDescProductnumber is the schema descriptor for Productnumber field.
	fixDescProductnumber := fixFields[0].Descriptor()
	// fix.ProductnumberValidator is a validator for the "Productnumber" field. It is called by the builders before save.
	fix.ProductnumberValidator = func() func(string) error {
		validators := fixDescProductnumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Productnumber string) error {
			for _, fn := range fns {
				if err := fn(_Productnumber); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fixDescProblemtype is the schema descriptor for Problemtype field.
	fixDescProblemtype := fixFields[1].Descriptor()
	// fix.ProblemtypeValidator is a validator for the "Problemtype" field. It is called by the builders before save.
	fix.ProblemtypeValidator = func() func(string) error {
		validators := fixDescProblemtype.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Problemtype string) error {
			for _, fn := range fns {
				if err := fn(_Problemtype); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fixDescQueue is the schema descriptor for Queue field.
	fixDescQueue := fixFields[2].Descriptor()
	// fix.QueueValidator is a validator for the "Queue" field. It is called by the builders before save.
	fix.QueueValidator = func() func(string) error {
		validators := fixDescQueue.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Queue string) error {
			for _, fn := range fns {
				if err := fn(_Queue); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	personalFields := schema.Personal{}.Fields()
	_ = personalFields
	// personalDescPersonalname is the schema descriptor for Personalname field.
	personalDescPersonalname := personalFields[0].Descriptor()
	// personal.PersonalnameValidator is a validator for the "Personalname" field. It is called by the builders before save.
	personal.PersonalnameValidator = personalDescPersonalname.Validators[0].(func(string) error)
	// personalDescEmail is the schema descriptor for Email field.
	personalDescEmail := personalFields[1].Descriptor()
	// personal.EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	personal.EmailValidator = personalDescEmail.Validators[0].(func(string) error)
	// personalDescPassword is the schema descriptor for Password field.
	personalDescPassword := personalFields[2].Descriptor()
	// personal.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	personal.PasswordValidator = func() func(string) error {
		validators := personalDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Password string) error {
			for _, fn := range fns {
				if err := fn(_Password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescProductname is the schema descriptor for Productname field.
	productDescProductname := productFields[0].Descriptor()
	// product.ProductnameValidator is a validator for the "Productname" field. It is called by the builders before save.
	product.ProductnameValidator = productDescProductname.Validators[0].(func(string) error)
	// productDescAmountofproduct is the schema descriptor for Amountofproduct field.
	productDescAmountofproduct := productFields[1].Descriptor()
	// product.AmountofproductValidator is a validator for the "Amountofproduct" field. It is called by the builders before save.
	product.AmountofproductValidator = productDescAmountofproduct.Validators[0].(func(int) error)
	// productDescPrice is the schema descriptor for Price field.
	productDescPrice := productFields[2].Descriptor()
	// product.PriceValidator is a validator for the "Price" field. It is called by the builders before save.
	product.PriceValidator = productDescPrice.Validators[0].(func(int) error)
	receiptFields := schema.Receipt{}.Fields()
	_ = receiptFields
	// receiptDescAddedTime is the schema descriptor for added_time field.
	receiptDescAddedTime := receiptFields[0].Descriptor()
	// receipt.DefaultAddedTime holds the default value on creation for the added_time field.
	receipt.DefaultAddedTime = receiptDescAddedTime.Default.(func() time.Time)
	// receiptDescAddress is the schema descriptor for Address field.
	receiptDescAddress := receiptFields[1].Descriptor()
	// receipt.AddressValidator is a validator for the "Address" field. It is called by the builders before save.
	receipt.AddressValidator = receiptDescAddress.Validators[0].(func(string) error)
	// receiptDescProductname is the schema descriptor for Productname field.
	receiptDescProductname := receiptFields[2].Descriptor()
	// receipt.ProductnameValidator is a validator for the "Productname" field. It is called by the builders before save.
	receipt.ProductnameValidator = receiptDescProductname.Validators[0].(func(string) error)
	// receiptDescReceiptcode is the schema descriptor for Receiptcode field.
	receiptDescReceiptcode := receiptFields[3].Descriptor()
	// receipt.ReceiptcodeValidator is a validator for the "Receiptcode" field. It is called by the builders before save.
	receipt.ReceiptcodeValidator = func() func(string) error {
		validators := receiptDescReceiptcode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Receiptcode string) error {
			for _, fn := range fns {
				if err := fn(_Receiptcode); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
